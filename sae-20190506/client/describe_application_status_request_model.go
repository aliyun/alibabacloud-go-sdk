// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationStatusRequest
	GetAppId() *string
}

type DescribeApplicationStatusRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeApplicationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationStatusRequest) SetAppId(v string) *DescribeApplicationStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationStatusRequest) Validate() error {
	return dara.Validate(s)
}
