// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicesByAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceIds(v *QueryDevicesByAccountResponseBodyDeviceIds) *QueryDevicesByAccountResponseBody
	GetDeviceIds() *QueryDevicesByAccountResponseBodyDeviceIds
	SetRequestId(v string) *QueryDevicesByAccountResponseBody
	GetRequestId() *string
}

type QueryDevicesByAccountResponseBody struct {
	DeviceIds *QueryDevicesByAccountResponseBodyDeviceIds `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Struct"`
	// example:
	//
	// A8A24108-2AD0-4F6E-81C7-A8A24C2C2AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDevicesByAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicesByAccountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAccountResponseBody) GetDeviceIds() *QueryDevicesByAccountResponseBodyDeviceIds {
	return s.DeviceIds
}

func (s *QueryDevicesByAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDevicesByAccountResponseBody) SetDeviceIds(v *QueryDevicesByAccountResponseBodyDeviceIds) *QueryDevicesByAccountResponseBody {
	s.DeviceIds = v
	return s
}

func (s *QueryDevicesByAccountResponseBody) SetRequestId(v string) *QueryDevicesByAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicesByAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDevicesByAccountResponseBodyDeviceIds struct {
	DeviceId []*string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" type:"Repeated"`
}

func (s QueryDevicesByAccountResponseBodyDeviceIds) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicesByAccountResponseBodyDeviceIds) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAccountResponseBodyDeviceIds) GetDeviceId() []*string {
	return s.DeviceId
}

func (s *QueryDevicesByAccountResponseBodyDeviceIds) SetDeviceId(v []*string) *QueryDevicesByAccountResponseBodyDeviceIds {
	s.DeviceId = v
	return s
}

func (s *QueryDevicesByAccountResponseBodyDeviceIds) Validate() error {
	return dara.Validate(s)
}
