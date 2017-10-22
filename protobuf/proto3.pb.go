// Code generated by protoc-gen-go.
// source: proto3.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LargerP3Define struct {
	Provider        uint32  `protobuf:"varint,1,opt,name=provider" json:"provider,omitempty"`
	CompetitionId   uint32  `protobuf:"varint,2,opt,name=competition_id,json=competitionId" json:"competition_id,omitempty"`
	EventId         uint32  `protobuf:"varint,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	MarketId        uint32  `protobuf:"varint,4,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
	Outcome         string  `protobuf:"bytes,5,opt,name=outcome" json:"outcome,omitempty"`
	SpecialBetValue string  `protobuf:"bytes,6,opt,name=special_bet_value,json=specialBetValue" json:"special_bet_value,omitempty"`
	Odds            float64 `protobuf:"fixed64,7,opt,name=odds" json:"odds,omitempty"`
	IsLive          bool    `protobuf:"varint,8,opt,name=is_live,json=isLive" json:"is_live,omitempty"`
	CurrencyPair    string  `protobuf:"bytes,9,opt,name=currency_pair,json=currencyPair" json:"currency_pair,omitempty"`
	StakeFactor     float64 `protobuf:"fixed64,10,opt,name=stake_factor,json=stakeFactor" json:"stake_factor,omitempty"`
}

func (m *LargerP3Define) Reset()                    { *m = LargerP3Define{} }
func (m *LargerP3Define) String() string            { return proto.CompactTextString(m) }
func (*LargerP3Define) ProtoMessage()               {}
func (*LargerP3Define) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LargerP3Define) GetProvider() uint32 {
	if m != nil {
		return m.Provider
	}
	return 0
}

func (m *LargerP3Define) GetCompetitionId() uint32 {
	if m != nil {
		return m.CompetitionId
	}
	return 0
}

func (m *LargerP3Define) GetEventId() uint32 {
	if m != nil {
		return m.EventId
	}
	return 0
}

func (m *LargerP3Define) GetMarketId() uint32 {
	if m != nil {
		return m.MarketId
	}
	return 0
}

func (m *LargerP3Define) GetOutcome() string {
	if m != nil {
		return m.Outcome
	}
	return ""
}

func (m *LargerP3Define) GetSpecialBetValue() string {
	if m != nil {
		return m.SpecialBetValue
	}
	return ""
}

func (m *LargerP3Define) GetOdds() float64 {
	if m != nil {
		return m.Odds
	}
	return 0
}

func (m *LargerP3Define) GetIsLive() bool {
	if m != nil {
		return m.IsLive
	}
	return false
}

func (m *LargerP3Define) GetCurrencyPair() string {
	if m != nil {
		return m.CurrencyPair
	}
	return ""
}

func (m *LargerP3Define) GetStakeFactor() float64 {
	if m != nil {
		return m.StakeFactor
	}
	return 0
}

type SmallerP3Define struct {
	Provider        uint32  `protobuf:"varint,1,opt,name=provider" json:"provider,omitempty"`
	CompetitionId   uint32  `protobuf:"varint,2,opt,name=competition_id,json=competitionId" json:"competition_id,omitempty"`
	SpecialBetValue string  `protobuf:"bytes,6,opt,name=special_bet_value,json=specialBetValue" json:"special_bet_value,omitempty"`
	Odds            float64 `protobuf:"fixed64,7,opt,name=odds" json:"odds,omitempty"`
	IsLive          bool    `protobuf:"varint,8,opt,name=is_live,json=isLive" json:"is_live,omitempty"`
	StakeFactor     float64 `protobuf:"fixed64,10,opt,name=stake_factor,json=stakeFactor" json:"stake_factor,omitempty"`
}

func (m *SmallerP3Define) Reset()                    { *m = SmallerP3Define{} }
func (m *SmallerP3Define) String() string            { return proto.CompactTextString(m) }
func (*SmallerP3Define) ProtoMessage()               {}
func (*SmallerP3Define) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SmallerP3Define) GetProvider() uint32 {
	if m != nil {
		return m.Provider
	}
	return 0
}

func (m *SmallerP3Define) GetCompetitionId() uint32 {
	if m != nil {
		return m.CompetitionId
	}
	return 0
}

func (m *SmallerP3Define) GetSpecialBetValue() string {
	if m != nil {
		return m.SpecialBetValue
	}
	return ""
}

func (m *SmallerP3Define) GetOdds() float64 {
	if m != nil {
		return m.Odds
	}
	return 0
}

func (m *SmallerP3Define) GetIsLive() bool {
	if m != nil {
		return m.IsLive
	}
	return false
}

func (m *SmallerP3Define) GetStakeFactor() float64 {
	if m != nil {
		return m.StakeFactor
	}
	return 0
}

func init() {
	proto.RegisterType((*LargerP3Define)(nil), "protobuf.LargerP3Define")
	proto.RegisterType((*SmallerP3Define)(nil), "protobuf.SmallerP3Define")
}

func init() { proto.RegisterFile("proto3.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x92, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xe9, 0xba, 0xf6, 0xcf, 0xd8, 0xee, 0x62, 0x2e, 0x46, 0xbd, 0xe8, 0x8a, 0x20, 0x1e,
	0xbc, 0xec, 0x37, 0x10, 0x11, 0x84, 0x3d, 0x2c, 0x15, 0xbc, 0x96, 0xb4, 0x9d, 0x4a, 0xd8, 0xb6,
	0x29, 0x49, 0x5a, 0xf0, 0x0b, 0x7a, 0xf6, 0x23, 0x99, 0x4e, 0x5d, 0xf1, 0xe6, 0x45, 0x4f, 0x93,
	0xf7, 0x7b, 0x2f, 0x0f, 0x26, 0x04, 0xe2, 0x4e, 0x2b, 0xab, 0xd6, 0x77, 0x34, 0x58, 0x48, 0x23,
	0xef, 0xab, 0xd5, 0xfb, 0x0c, 0x16, 0x1b, 0xa1, 0x5f, 0x51, 0x6f, 0xd7, 0x0f, 0x58, 0xc9, 0x16,
	0xd9, 0x19, 0x8c, 0xf6, 0x20, 0x4b, 0xd4, 0xdc, 0xbb, 0xf0, 0x6e, 0x92, 0xf4, 0x5b, 0xb3, 0x6b,
	0x58, 0x14, 0xaa, 0xe9, 0xd0, 0x4a, 0x2b, 0x55, 0x9b, 0xc9, 0x92, 0xcf, 0x28, 0x91, 0xfc, 0xa0,
	0x4f, 0x25, 0x3b, 0x85, 0x10, 0x07, 0x6c, 0xed, 0x18, 0x38, 0xa0, 0x40, 0x40, 0xda, 0x59, 0xe7,
	0x10, 0x35, 0x42, 0xef, 0x90, 0xbc, 0xf9, 0x54, 0x3f, 0x01, 0x67, 0x72, 0x08, 0x54, 0x6f, 0x5d,
	0x17, 0xf2, 0x43, 0x67, 0x45, 0xe9, 0x5e, 0xb2, 0x5b, 0x38, 0x36, 0x1d, 0x16, 0x52, 0xd4, 0x59,
	0xee, 0xee, 0x0e, 0xa2, 0xee, 0x91, 0xfb, 0x94, 0x59, 0x7e, 0x19, 0xf7, 0x68, 0x5f, 0x46, 0xcc,
	0x18, 0xcc, 0x55, 0x59, 0x1a, 0x1e, 0x38, 0xdb, 0x4b, 0xe9, 0xcc, 0x4e, 0x20, 0x90, 0x26, 0xab,
	0xe5, 0x80, 0x3c, 0x74, 0x38, 0x4c, 0x7d, 0x69, 0x36, 0x4e, 0xb1, 0x2b, 0x48, 0x8a, 0x5e, 0x6b,
	0x6c, 0x8b, 0xb7, 0xac, 0x13, 0x52, 0xf3, 0x88, 0x4a, 0xe3, 0x3d, 0xdc, 0x3a, 0xc6, 0x2e, 0x21,
	0x36, 0x56, 0xec, 0x30, 0xab, 0x44, 0x61, 0x95, 0xe6, 0x40, 0xcd, 0x47, 0xc4, 0x1e, 0x09, 0xad,
	0x3e, 0x3c, 0x58, 0x3e, 0x37, 0xa2, 0xae, 0xff, 0xf6, 0x25, 0xff, 0x6d, 0xef, 0xdf, 0x57, 0xca,
	0xfd, 0xe9, 0xcf, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x62, 0xf5, 0x7a, 0x3c, 0x02, 0x00,
	0x00,
}
