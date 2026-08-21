// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLineInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeCustomLineInfoResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeCustomLineInfoResponseBody
	GetCreateTimestamp() *int64
	SetCreator(v string) *DescribeCustomLineInfoResponseBody
	GetCreator() *string
	SetCreatorSubType(v string) *DescribeCustomLineInfoResponseBody
	GetCreatorSubType() *string
	SetCreatorType(v string) *DescribeCustomLineInfoResponseBody
	GetCreatorType() *string
	SetDnscategory(v string) *DescribeCustomLineInfoResponseBody
	GetDnscategory() *string
	SetIpv4s(v []*string) *DescribeCustomLineInfoResponseBody
	GetIpv4s() []*string
	SetLineId(v string) *DescribeCustomLineInfoResponseBody
	GetLineId() *string
	SetName(v string) *DescribeCustomLineInfoResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeCustomLineInfoResponseBody
	GetRequestId() *string
	SetUpdateTime(v string) *DescribeCustomLineInfoResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeCustomLineInfoResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeCustomLineInfoResponseBody struct {
	// The creation time.
	//
	// example:
	//
	// 2018-01-23T03:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creation timestamp in milliseconds.
	//
	// example:
	//
	// 1516775741000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The creator.
	//
	// example:
	//
	// 260282302749096109
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The subtype of the creator. Valid values:
	//
	// - `CUSTOM`: Alibaba Cloud account
	//
	// - `SUB`: RAM user
	//
	// - `STS`: assumed role
	//
	// - `OTHER`: other roles
	//
	// example:
	//
	// CUSTOM
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// The creator type. Valid values:
	//
	// - `USER`: user
	//
	// - `SYSTEM`: system
	//
	// example:
	//
	// USER
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	Dnscategory *string `json:"Dnscategory,omitempty" xml:"Dnscategory,omitempty"`
	// A list of IPv4 address segments.
	Ipv4s []*string `json:"Ipv4s,omitempty" xml:"Ipv4s,omitempty" type:"Repeated"`
	// The unique ID of the custom line.
	//
	// example:
	//
	// 100003
	LineId *string `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The custom line name.
	//
	// example:
	//
	// Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The last update time.
	//
	// example:
	//
	// 2018-01-24T06:35Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The last update timestamp in milliseconds.
	//
	// example:
	//
	// 1516775741000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeCustomLineInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLineInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineInfoResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCustomLineInfoResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCustomLineInfoResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *DescribeCustomLineInfoResponseBody) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *DescribeCustomLineInfoResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *DescribeCustomLineInfoResponseBody) GetDnscategory() *string {
	return s.Dnscategory
}

func (s *DescribeCustomLineInfoResponseBody) GetIpv4s() []*string {
	return s.Ipv4s
}

func (s *DescribeCustomLineInfoResponseBody) GetLineId() *string {
	return s.LineId
}

func (s *DescribeCustomLineInfoResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeCustomLineInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomLineInfoResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCustomLineInfoResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCustomLineInfoResponseBody) SetCreateTime(v string) *DescribeCustomLineInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetCreateTimestamp(v int64) *DescribeCustomLineInfoResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetCreator(v string) *DescribeCustomLineInfoResponseBody {
	s.Creator = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetCreatorSubType(v string) *DescribeCustomLineInfoResponseBody {
	s.CreatorSubType = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetCreatorType(v string) *DescribeCustomLineInfoResponseBody {
	s.CreatorType = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetDnscategory(v string) *DescribeCustomLineInfoResponseBody {
	s.Dnscategory = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetIpv4s(v []*string) *DescribeCustomLineInfoResponseBody {
	s.Ipv4s = v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetLineId(v string) *DescribeCustomLineInfoResponseBody {
	s.LineId = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetName(v string) *DescribeCustomLineInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetRequestId(v string) *DescribeCustomLineInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetUpdateTime(v string) *DescribeCustomLineInfoResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) SetUpdateTimestamp(v int64) *DescribeCustomLineInfoResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCustomLineInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
