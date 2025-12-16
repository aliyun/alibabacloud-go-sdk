// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPassThroughAuthInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetPassThroughAuthInfoRequest
	GetAppId() *string
	SetDeviceName(v string) *GetPassThroughAuthInfoRequest
	GetDeviceName() *string
}

type GetPassThroughAuthInfoRequest struct {
	// example:
	//
	// mm_2b7d37b695fb44faa983e5fbb032
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 35f9ed99b27a1e8374f6593bf969f7d6
	DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
}

func (s GetPassThroughAuthInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPassThroughAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPassThroughAuthInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetPassThroughAuthInfoRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *GetPassThroughAuthInfoRequest) SetAppId(v string) *GetPassThroughAuthInfoRequest {
	s.AppId = &v
	return s
}

func (s *GetPassThroughAuthInfoRequest) SetDeviceName(v string) *GetPassThroughAuthInfoRequest {
	s.DeviceName = &v
	return s
}

func (s *GetPassThroughAuthInfoRequest) Validate() error {
	return dara.Validate(s)
}
