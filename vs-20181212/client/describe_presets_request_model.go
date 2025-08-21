// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePresetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribePresetsRequest
	GetId() *string
	SetOwnerId(v int64) *DescribePresetsRequest
	GetOwnerId() *int64
}

type DescribePresetsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribePresetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePresetsRequest) GoString() string {
	return s.String()
}

func (s *DescribePresetsRequest) GetId() *string {
	return s.Id
}

func (s *DescribePresetsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePresetsRequest) SetId(v string) *DescribePresetsRequest {
	s.Id = &v
	return s
}

func (s *DescribePresetsRequest) SetOwnerId(v int64) *DescribePresetsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePresetsRequest) Validate() error {
	return dara.Validate(s)
}
