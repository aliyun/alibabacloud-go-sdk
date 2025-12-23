// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpcDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfos(v []*GetIpcDeviceInfoResponseBodyDeviceInfos) *GetIpcDeviceInfoResponseBody
	GetDeviceInfos() []*GetIpcDeviceInfoResponseBodyDeviceInfos
	SetRequestId(v string) *GetIpcDeviceInfoResponseBody
	GetRequestId() *string
	SetTotal(v int64) *GetIpcDeviceInfoResponseBody
	GetTotal() *int64
}

type GetIpcDeviceInfoResponseBody struct {
	DeviceInfos []*GetIpcDeviceInfoResponseBodyDeviceInfos `json:"DeviceInfos,omitempty" xml:"DeviceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 78
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetIpcDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIpcDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpcDeviceInfoResponseBody) GetDeviceInfos() []*GetIpcDeviceInfoResponseBodyDeviceInfos {
	return s.DeviceInfos
}

func (s *GetIpcDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIpcDeviceInfoResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetIpcDeviceInfoResponseBody) SetDeviceInfos(v []*GetIpcDeviceInfoResponseBodyDeviceInfos) *GetIpcDeviceInfoResponseBody {
	s.DeviceInfos = v
	return s
}

func (s *GetIpcDeviceInfoResponseBody) SetRequestId(v string) *GetIpcDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpcDeviceInfoResponseBody) SetTotal(v int64) *GetIpcDeviceInfoResponseBody {
	s.Total = &v
	return s
}

func (s *GetIpcDeviceInfoResponseBody) Validate() error {
	if s.DeviceInfos != nil {
		for _, item := range s.DeviceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetIpcDeviceInfoResponseBodyDeviceInfos struct {
	// example:
	//
	// understand
	Capability *string `json:"Capability,omitempty" xml:"Capability,omitempty"`
	// example:
	//
	// d123
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 2017-02-11T12:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s GetIpcDeviceInfoResponseBodyDeviceInfos) String() string {
	return dara.Prettify(s)
}

func (s GetIpcDeviceInfoResponseBodyDeviceInfos) GoString() string {
	return s.String()
}

func (s *GetIpcDeviceInfoResponseBodyDeviceInfos) GetCapability() *string {
	return s.Capability
}

func (s *GetIpcDeviceInfoResponseBodyDeviceInfos) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetIpcDeviceInfoResponseBodyDeviceInfos) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetIpcDeviceInfoResponseBodyDeviceInfos) SetCapability(v string) *GetIpcDeviceInfoResponseBodyDeviceInfos {
	s.Capability = &v
	return s
}

func (s *GetIpcDeviceInfoResponseBodyDeviceInfos) SetDeviceId(v string) *GetIpcDeviceInfoResponseBodyDeviceInfos {
	s.DeviceId = &v
	return s
}

func (s *GetIpcDeviceInfoResponseBodyDeviceInfos) SetExpireTime(v string) *GetIpcDeviceInfoResponseBodyDeviceInfos {
	s.ExpireTime = &v
	return s
}

func (s *GetIpcDeviceInfoResponseBodyDeviceInfos) Validate() error {
	return dara.Validate(s)
}
