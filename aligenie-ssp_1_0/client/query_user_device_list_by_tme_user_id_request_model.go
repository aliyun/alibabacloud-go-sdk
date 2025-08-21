// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserDeviceListByTmeUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSp(v string) *QueryUserDeviceListByTmeUserIdRequest
	GetSp() *string
	SetTmeUserId(v string) *QueryUserDeviceListByTmeUserIdRequest
	GetTmeUserId() *string
}

type QueryUserDeviceListByTmeUserIdRequest struct {
	// This parameter is required.
	Sp *string `json:"Sp,omitempty" xml:"Sp,omitempty"`
	// This parameter is required.
	TmeUserId *string `json:"TmeUserId,omitempty" xml:"TmeUserId,omitempty"`
}

func (s QueryUserDeviceListByTmeUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserDeviceListByTmeUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryUserDeviceListByTmeUserIdRequest) GetSp() *string {
	return s.Sp
}

func (s *QueryUserDeviceListByTmeUserIdRequest) GetTmeUserId() *string {
	return s.TmeUserId
}

func (s *QueryUserDeviceListByTmeUserIdRequest) SetSp(v string) *QueryUserDeviceListByTmeUserIdRequest {
	s.Sp = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdRequest) SetTmeUserId(v string) *QueryUserDeviceListByTmeUserIdRequest {
	s.TmeUserId = &v
	return s
}

func (s *QueryUserDeviceListByTmeUserIdRequest) Validate() error {
	return dara.Validate(s)
}
