// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *DeleteNetDeviceInfoRequest
	GetAppCode() *string
	SetAppName(v string) *DeleteNetDeviceInfoRequest
	GetAppName() *string
	SetIds(v string) *DeleteNetDeviceInfoRequest
	GetIds() *string
	SetRequestId(v string) *DeleteNetDeviceInfoRequest
	GetRequestId() *string
}

type DeleteNetDeviceInfoRequest struct {
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetDeviceInfoRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DeleteNetDeviceInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteNetDeviceInfoRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteNetDeviceInfoRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetDeviceInfoRequest) SetAppCode(v string) *DeleteNetDeviceInfoRequest {
	s.AppCode = &v
	return s
}

func (s *DeleteNetDeviceInfoRequest) SetAppName(v string) *DeleteNetDeviceInfoRequest {
	s.AppName = &v
	return s
}

func (s *DeleteNetDeviceInfoRequest) SetIds(v string) *DeleteNetDeviceInfoRequest {
	s.Ids = &v
	return s
}

func (s *DeleteNetDeviceInfoRequest) SetRequestId(v string) *DeleteNetDeviceInfoRequest {
	s.RequestId = &v
	return s
}

func (s *DeleteNetDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
