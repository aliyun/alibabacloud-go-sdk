// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *DescribeProductRequest
	GetAliUid() *string
	SetCode(v string) *DescribeProductRequest
	GetCode() *string
	SetQueryDraft(v bool) *DescribeProductRequest
	GetQueryDraft() *bool
}

type DescribeProductRequest struct {
	// AliUid
	//
	// example:
	//
	// 190********569
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cmjj01**45
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	QueryDraft *bool `json:"QueryDraft,omitempty" xml:"QueryDraft,omitempty"`
}

func (s DescribeProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductRequest) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeProductRequest) GetCode() *string {
	return s.Code
}

func (s *DescribeProductRequest) GetQueryDraft() *bool {
	return s.QueryDraft
}

func (s *DescribeProductRequest) SetAliUid(v string) *DescribeProductRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeProductRequest) SetCode(v string) *DescribeProductRequest {
	s.Code = &v
	return s
}

func (s *DescribeProductRequest) SetQueryDraft(v bool) *DescribeProductRequest {
	s.QueryDraft = &v
	return s
}

func (s *DescribeProductRequest) Validate() error {
	return dara.Validate(s)
}
