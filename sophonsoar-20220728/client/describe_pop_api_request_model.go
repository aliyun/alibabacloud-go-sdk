// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePopApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DescribePopApiRequest
	GetApiName() *string
	SetApiVersion(v string) *DescribePopApiRequest
	GetApiVersion() *string
	SetPopCode(v string) *DescribePopApiRequest
	GetPopCode() *string
}

type DescribePopApiRequest struct {
	// The operation name of the Alibaba Cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// DescribeInstanceInfo
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The version number of the API.
	//
	// >  You can call the [DescribePopApiVersionList](~~DescribePopApiVersionList~~) operation to query the version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-10-01
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The POP code of the Alibaba Cloud service.
	//
	// >  You can call the [DescribeApiList](~~DescribeApiList~~) operation to query the POP code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
}

func (s DescribePopApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePopApiRequest) GoString() string {
	return s.String()
}

func (s *DescribePopApiRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribePopApiRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *DescribePopApiRequest) GetPopCode() *string {
	return s.PopCode
}

func (s *DescribePopApiRequest) SetApiName(v string) *DescribePopApiRequest {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiRequest) SetApiVersion(v string) *DescribePopApiRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribePopApiRequest) SetPopCode(v string) *DescribePopApiRequest {
	s.PopCode = &v
	return s
}

func (s *DescribePopApiRequest) Validate() error {
	return dara.Validate(s)
}
