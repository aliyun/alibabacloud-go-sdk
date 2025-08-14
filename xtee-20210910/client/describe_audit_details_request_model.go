// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAuditDetailsRequest
	GetLang() *string
	SetId(v int64) *DescribeAuditDetailsRequest
	GetId() *int64
	SetRegId(v string) *DescribeAuditDetailsRequest
	GetRegId() *string
}

type DescribeAuditDetailsRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2557
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAuditDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditDetailsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuditDetailsRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeAuditDetailsRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAuditDetailsRequest) SetLang(v string) *DescribeAuditDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuditDetailsRequest) SetId(v int64) *DescribeAuditDetailsRequest {
	s.Id = &v
	return s
}

func (s *DescribeAuditDetailsRequest) SetRegId(v string) *DescribeAuditDetailsRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAuditDetailsRequest) Validate() error {
	return dara.Validate(s)
}
