// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDevicesByAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceIds(v *QueryDevicesByAliasResponseBodyDeviceIds) *QueryDevicesByAliasResponseBody
	GetDeviceIds() *QueryDevicesByAliasResponseBodyDeviceIds
	SetRequestId(v string) *QueryDevicesByAliasResponseBody
	GetRequestId() *string
}

type QueryDevicesByAliasResponseBody struct {
	DeviceIds *QueryDevicesByAliasResponseBodyDeviceIds `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Struct"`
	// example:
	//
	// 6A9FD644-35A5-40E4-89B0-2021CAEDC1B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDevicesByAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicesByAliasResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAliasResponseBody) GetDeviceIds() *QueryDevicesByAliasResponseBodyDeviceIds {
	return s.DeviceIds
}

func (s *QueryDevicesByAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDevicesByAliasResponseBody) SetDeviceIds(v *QueryDevicesByAliasResponseBodyDeviceIds) *QueryDevicesByAliasResponseBody {
	s.DeviceIds = v
	return s
}

func (s *QueryDevicesByAliasResponseBody) SetRequestId(v string) *QueryDevicesByAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicesByAliasResponseBody) Validate() error {
	if s.DeviceIds != nil {
		if err := s.DeviceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDevicesByAliasResponseBodyDeviceIds struct {
	DeviceId []*string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty" type:"Repeated"`
}

func (s QueryDevicesByAliasResponseBodyDeviceIds) String() string {
	return dara.Prettify(s)
}

func (s QueryDevicesByAliasResponseBodyDeviceIds) GoString() string {
	return s.String()
}

func (s *QueryDevicesByAliasResponseBodyDeviceIds) GetDeviceId() []*string {
	return s.DeviceId
}

func (s *QueryDevicesByAliasResponseBodyDeviceIds) SetDeviceId(v []*string) *QueryDevicesByAliasResponseBodyDeviceIds {
	s.DeviceId = v
	return s
}

func (s *QueryDevicesByAliasResponseBodyDeviceIds) Validate() error {
	return dara.Validate(s)
}
