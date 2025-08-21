// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeTemplateRequest
	GetId() *string
	SetOwnerId(v int64) *DescribeTemplateRequest
	GetOwnerId() *int64
}

type DescribeTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323*****998-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateRequest) GetId() *string {
	return s.Id
}

func (s *DescribeTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTemplateRequest) SetId(v string) *DescribeTemplateRequest {
	s.Id = &v
	return s
}

func (s *DescribeTemplateRequest) SetOwnerId(v int64) *DescribeTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
