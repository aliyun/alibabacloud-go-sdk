// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAndSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *DescribeServiceAndSceneRequest
	GetAuthType() *string
	SetCustomerModuleId(v int32) *DescribeServiceAndSceneRequest
	GetCustomerModuleId() *int32
}

type DescribeServiceAndSceneRequest struct {
	// example:
	//
	// READ
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s DescribeServiceAndSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAndSceneRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceAndSceneRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeServiceAndSceneRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *DescribeServiceAndSceneRequest) SetAuthType(v string) *DescribeServiceAndSceneRequest {
	s.AuthType = &v
	return s
}

func (s *DescribeServiceAndSceneRequest) SetCustomerModuleId(v int32) *DescribeServiceAndSceneRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *DescribeServiceAndSceneRequest) Validate() error {
	return dara.Validate(s)
}
