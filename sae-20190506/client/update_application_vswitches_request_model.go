// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationVswitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateApplicationVswitchesRequest
	GetAppId() *string
	SetVSwitchId(v string) *UpdateApplicationVswitchesRequest
	GetVSwitchId() *string
}

type UpdateApplicationVswitchesRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2ze559r1z1bpwqxwp****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateApplicationVswitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVswitchesRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVswitchesRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationVswitchesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateApplicationVswitchesRequest) SetAppId(v string) *UpdateApplicationVswitchesRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationVswitchesRequest) SetVSwitchId(v string) *UpdateApplicationVswitchesRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateApplicationVswitchesRequest) Validate() error {
	return dara.Validate(s)
}
