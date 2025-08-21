// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDeviceChannelsResponseBody
	GetRequestId() *string
}

type ModifyDeviceChannelsResponseBody struct {
	// example:
	//
	// 8F4D95B6-EB19-5245-AD77-95EDA83E53B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDeviceChannelsResponseBody) SetRequestId(v string) *ModifyDeviceChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceChannelsResponseBody) Validate() error {
	return dara.Validate(s)
}
