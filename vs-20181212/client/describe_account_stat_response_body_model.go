// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupLimit(v int64) *DescribeAccountStatResponseBody
	GetGroupLimit() *int64
	SetGroupNum(v int64) *DescribeAccountStatResponseBody
	GetGroupNum() *int64
	SetId(v string) *DescribeAccountStatResponseBody
	GetId() *string
	SetRequestId(v string) *DescribeAccountStatResponseBody
	GetRequestId() *string
	SetTemplateLimit(v int64) *DescribeAccountStatResponseBody
	GetTemplateLimit() *int64
	SetTemplateNum(v int64) *DescribeAccountStatResponseBody
	GetTemplateNum() *int64
}

type DescribeAccountStatResponseBody struct {
	// example:
	//
	// 100
	GroupLimit *int64 `json:"GroupLimit,omitempty" xml:"GroupLimit,omitempty"`
	// example:
	//
	// 6
	GroupNum *int64 `json:"GroupNum,omitempty" xml:"GroupNum,omitempty"`
	// ID
	//
	// example:
	//
	// 3238848****092996
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TemplateLimit *int64 `json:"TemplateLimit,omitempty" xml:"TemplateLimit,omitempty"`
	// example:
	//
	// 10
	TemplateNum *int64 `json:"TemplateNum,omitempty" xml:"TemplateNum,omitempty"`
}

func (s DescribeAccountStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountStatResponseBody) GetGroupLimit() *int64 {
	return s.GroupLimit
}

func (s *DescribeAccountStatResponseBody) GetGroupNum() *int64 {
	return s.GroupNum
}

func (s *DescribeAccountStatResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeAccountStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountStatResponseBody) GetTemplateLimit() *int64 {
	return s.TemplateLimit
}

func (s *DescribeAccountStatResponseBody) GetTemplateNum() *int64 {
	return s.TemplateNum
}

func (s *DescribeAccountStatResponseBody) SetGroupLimit(v int64) *DescribeAccountStatResponseBody {
	s.GroupLimit = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetGroupNum(v int64) *DescribeAccountStatResponseBody {
	s.GroupNum = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetId(v string) *DescribeAccountStatResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetRequestId(v string) *DescribeAccountStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetTemplateLimit(v int64) *DescribeAccountStatResponseBody {
	s.TemplateLimit = &v
	return s
}

func (s *DescribeAccountStatResponseBody) SetTemplateNum(v int64) *DescribeAccountStatResponseBody {
	s.TemplateNum = &v
	return s
}

func (s *DescribeAccountStatResponseBody) Validate() error {
	return dara.Validate(s)
}
