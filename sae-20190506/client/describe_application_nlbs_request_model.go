// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationNlbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationNlbsRequest
	GetAppId() *string
}

type DescribeApplicationNlbsRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeApplicationNlbsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationNlbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationNlbsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationNlbsRequest) SetAppId(v string) *DescribeApplicationNlbsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationNlbsRequest) Validate() error {
	return dara.Validate(s)
}
