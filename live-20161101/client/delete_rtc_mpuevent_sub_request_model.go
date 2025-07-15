// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtcMPUEventSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteRtcMPUEventSubRequest
	GetAppId() *string
}

type DeleteRtcMPUEventSubRequest struct {
	// The ID of the application.
	//
	// >  The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteRtcMPUEventSubRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtcMPUEventSubRequest) GoString() string {
	return s.String()
}

func (s *DeleteRtcMPUEventSubRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteRtcMPUEventSubRequest) SetAppId(v string) *DeleteRtcMPUEventSubRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRtcMPUEventSubRequest) Validate() error {
	return dara.Validate(s)
}
