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
	// The ID of the subscribed application. You can view your application IDs by navigating to **ApsaraVideo Live > Live+ > ApsaraVideo Real-time Communication > Application Management**.
	//
	// >
	//
	// > - The application ID consists of uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters.
	//
	// > - You must first call CreateRtcMPUEventSub to create a stream mixing and forwarding event subscription for this application ID.
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
