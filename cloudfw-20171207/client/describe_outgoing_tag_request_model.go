// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstType(v string) *DescribeOutgoingTagRequest
	GetDstType() *string
	SetEndTime(v string) *DescribeOutgoingTagRequest
	GetEndTime() *string
	SetLang(v string) *DescribeOutgoingTagRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeOutgoingTagRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeOutgoingTagRequest
	GetStartTime() *string
	SetTagId(v string) *DescribeOutgoingTagRequest
	GetTagId() *string
}

type DescribeOutgoingTagRequest struct {
	// example:
	//
	// Domain
	DstType *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1749003483
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 60.179.226.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1743646544
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 103208
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DescribeOutgoingTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingTagRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingTagRequest) GetDstType() *string {
	return s.DstType
}

func (s *DescribeOutgoingTagRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingTagRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingTagRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOutgoingTagRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingTagRequest) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingTagRequest) SetDstType(v string) *DescribeOutgoingTagRequest {
	s.DstType = &v
	return s
}

func (s *DescribeOutgoingTagRequest) SetEndTime(v string) *DescribeOutgoingTagRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingTagRequest) SetLang(v string) *DescribeOutgoingTagRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingTagRequest) SetSourceIp(v string) *DescribeOutgoingTagRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingTagRequest) SetStartTime(v string) *DescribeOutgoingTagRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingTagRequest) SetTagId(v string) *DescribeOutgoingTagRequest {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingTagRequest) Validate() error {
	return dara.Validate(s)
}
