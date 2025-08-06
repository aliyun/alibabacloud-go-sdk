// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *GetPlayConfigResponseBody
	GetConfigType() *string
	SetPlayDeviceAbilityList(v []*string) *GetPlayConfigResponseBody
	GetPlayDeviceAbilityList() []*string
	SetRequestId(v string) *GetPlayConfigResponseBody
	GetRequestId() *string
}

type GetPlayConfigResponseBody struct {
	ConfigType            *string   `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	PlayDeviceAbilityList []*string `json:"PlayDeviceAbilityList,omitempty" xml:"PlayDeviceAbilityList,omitempty" type:"Repeated"`
	RequestId             *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPlayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPlayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlayConfigResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetPlayConfigResponseBody) GetPlayDeviceAbilityList() []*string {
	return s.PlayDeviceAbilityList
}

func (s *GetPlayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPlayConfigResponseBody) SetConfigType(v string) *GetPlayConfigResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetPlayConfigResponseBody) SetPlayDeviceAbilityList(v []*string) *GetPlayConfigResponseBody {
	s.PlayDeviceAbilityList = v
	return s
}

func (s *GetPlayConfigResponseBody) SetRequestId(v string) *GetPlayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
