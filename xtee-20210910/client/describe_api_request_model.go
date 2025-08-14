// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeApiRequest
	GetLang() *string
	SetApiId(v string) *DescribeApiRequest
	GetApiId() *string
	SetApiRegionId(v string) *DescribeApiRequest
	GetApiRegionId() *string
	SetApiType(v string) *DescribeApiRequest
	GetApiType() *string
	SetRegId(v string) *DescribeApiRequest
	GetRegId() *string
}

type DescribeApiRequest struct {
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
	// API unique ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 33
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou、cn-shanghai
	ApiRegionId *string `json:"apiRegionId,omitempty" xml:"apiRegionId,omitempty"`
	// API type
	//
	// This parameter is required.
	//
	// example:
	//
	// SELF
	ApiType *string `json:"apiType,omitempty" xml:"apiType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiRequest) GetApiRegionId() *string {
	return s.ApiRegionId
}

func (s *DescribeApiRequest) GetApiType() *string {
	return s.ApiType
}

func (s *DescribeApiRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeApiRequest) SetLang(v string) *DescribeApiRequest {
	s.Lang = &v
	return s
}

func (s *DescribeApiRequest) SetApiId(v string) *DescribeApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiRequest) SetApiRegionId(v string) *DescribeApiRequest {
	s.ApiRegionId = &v
	return s
}

func (s *DescribeApiRequest) SetApiType(v string) *DescribeApiRequest {
	s.ApiType = &v
	return s
}

func (s *DescribeApiRequest) SetRegId(v string) *DescribeApiRequest {
	s.RegId = &v
	return s
}

func (s *DescribeApiRequest) Validate() error {
	return dara.Validate(s)
}
