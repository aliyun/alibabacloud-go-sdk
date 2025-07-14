// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeComponentsRequest
	GetAppId() *string
	SetType(v string) *DescribeComponentsRequest
	GetType() *string
}

type DescribeComponentsRequest struct {
	// The application ID.
	//
	// example:
	//
	// d700e680-aa4d-4ec1-afc2-6566b5ff****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The type of the supported components. Valid values:
	//
	// 	- **TOMCAT**
	//
	// 	- **JDK**
	//
	// This parameter is required.
	//
	// example:
	//
	// TOMCAT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeComponentsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeComponentsRequest) SetAppId(v string) *DescribeComponentsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeComponentsRequest) SetType(v string) *DescribeComponentsRequest {
	s.Type = &v
	return s
}

func (s *DescribeComponentsRequest) Validate() error {
	return dara.Validate(s)
}
