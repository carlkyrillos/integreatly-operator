//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/alts/v3/alts.proto

package altsv3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *Alts) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Alts) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Alts) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.PeerServiceAccounts) > 0 {
		for iNdEx := len(m.PeerServiceAccounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PeerServiceAccounts[iNdEx])
			copy(dAtA[i:], m.PeerServiceAccounts[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.PeerServiceAccounts[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.HandshakerService) > 0 {
		i -= len(m.HandshakerService)
		copy(dAtA[i:], m.HandshakerService)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.HandshakerService)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Alts) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HandshakerService)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.PeerServiceAccounts) > 0 {
		for _, s := range m.PeerServiceAccounts {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}