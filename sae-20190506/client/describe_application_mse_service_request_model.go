// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationMseServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationMseServiceRequest
	GetAppId() *string
	SetEnableAhas(v bool) *DescribeApplicationMseServiceRequest
	GetEnableAhas() *bool
}

type DescribeApplicationMseServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f7730764-d88f-4b9a-8d8e-cd8efbfe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableAhas *bool `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
}

func (s DescribeApplicationMseServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationMseServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMseServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationMseServiceRequest) GetEnableAhas() *bool {
	return s.EnableAhas
}

func (s *DescribeApplicationMseServiceRequest) SetAppId(v string) *DescribeApplicationMseServiceRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationMseServiceRequest) SetEnableAhas(v bool) *DescribeApplicationMseServiceRequest {
	s.EnableAhas = &v
	return s
}

func (s *DescribeApplicationMseServiceRequest) Validate() error {
	return dara.Validate(s)
}
