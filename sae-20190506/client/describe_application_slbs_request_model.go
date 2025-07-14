// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSlbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationSlbsRequest
	GetAppId() *string
}

type DescribeApplicationSlbsRequest struct {
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeApplicationSlbsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSlbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationSlbsRequest) SetAppId(v string) *DescribeApplicationSlbsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationSlbsRequest) Validate() error {
	return dara.Validate(s)
}
