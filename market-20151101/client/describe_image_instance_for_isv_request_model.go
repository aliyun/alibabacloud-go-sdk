// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInstanceForIsvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerPk(v int64) *DescribeImageInstanceForIsvRequest
	GetCustomerPk() *int64
	SetEcsInstanceId(v string) *DescribeImageInstanceForIsvRequest
	GetEcsInstanceId() *string
}

type DescribeImageInstanceForIsvRequest struct {
	CustomerPk    *int64  `json:"CustomerPk,omitempty" xml:"CustomerPk,omitempty"`
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
}

func (s DescribeImageInstanceForIsvRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstanceForIsvRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageInstanceForIsvRequest) GetCustomerPk() *int64 {
	return s.CustomerPk
}

func (s *DescribeImageInstanceForIsvRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeImageInstanceForIsvRequest) SetCustomerPk(v int64) *DescribeImageInstanceForIsvRequest {
	s.CustomerPk = &v
	return s
}

func (s *DescribeImageInstanceForIsvRequest) SetEcsInstanceId(v string) *DescribeImageInstanceForIsvRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeImageInstanceForIsvRequest) Validate() error {
	return dara.Validate(s)
}
