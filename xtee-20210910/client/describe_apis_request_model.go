// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeApisRequest
	GetLang() *string
	SetApiGroupId(v string) *DescribeApisRequest
	GetApiGroupId() *string
	SetApiRegionId(v string) *DescribeApisRequest
	GetApiRegionId() *string
	SetApiType(v string) *DescribeApisRequest
	GetApiType() *string
	SetRegId(v string) *DescribeApisRequest
	GetRegId() *string
}

type DescribeApisRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// API group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3309b0f4b1e243cd8bd9dd029f9c5f0a
	ApiGroupId *string `json:"apiGroupId,omitempty" xml:"apiGroupId,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou、cn-shanghai
	ApiRegionId *string `json:"apiRegionId,omitempty" xml:"apiRegionId,omitempty"`
	// API type.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELF
	ApiType *string `json:"apiType,omitempty" xml:"apiType,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeApisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApisRequest) GetApiGroupId() *string {
	return s.ApiGroupId
}

func (s *DescribeApisRequest) GetApiRegionId() *string {
	return s.ApiRegionId
}

func (s *DescribeApisRequest) GetApiType() *string {
	return s.ApiType
}

func (s *DescribeApisRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeApisRequest) SetLang(v string) *DescribeApisRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApisRequest) SetApiGroupId(v string) *DescribeApisRequest {
	s.ApiGroupId = &v
	return s
}

func (s *DescribeApisRequest) SetApiRegionId(v string) *DescribeApisRequest {
	s.ApiRegionId = &v
	return s
}

func (s *DescribeApisRequest) SetApiType(v string) *DescribeApisRequest {
	s.ApiType = &v
	return s
}

func (s *DescribeApisRequest) SetRegId(v string) *DescribeApisRequest {
	s.RegId = &v
	return s
}

func (s *DescribeApisRequest) Validate() error {
	return dara.Validate(s)
}
