// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcMPUEventSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRtcMPUEventSubRequest
	GetAppId() *string
}

type DescribeRtcMPUEventSubRequest struct {
	// The ID of your application.
	//
	// > The ID can be up to 64 characters in length and can contain letters, digits, underscores, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeRtcMPUEventSubRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcMPUEventSubRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcMPUEventSubRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcMPUEventSubRequest) SetAppId(v string) *DescribeRtcMPUEventSubRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcMPUEventSubRequest) Validate() error {
	return dara.Validate(s)
}
