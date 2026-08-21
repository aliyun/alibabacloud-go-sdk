// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwareForUserDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSoftwareForUserDeviceResponseBody
	GetRequestId() *string
	SetSoftware(v []*ListSoftwareForUserDeviceResponseBodySoftware) *ListSoftwareForUserDeviceResponseBody
	GetSoftware() []*ListSoftwareForUserDeviceResponseBodySoftware
	SetTotalNum(v int64) *ListSoftwareForUserDeviceResponseBody
	GetTotalNum() *int64
}

type ListSoftwareForUserDeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of software installed on the endpoint device.
	Software []*ListSoftwareForUserDeviceResponseBodySoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Repeated"`
	// The total number of software installed on the endpoint device.
	//
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListSoftwareForUserDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwareForUserDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwareForUserDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSoftwareForUserDeviceResponseBody) GetSoftware() []*ListSoftwareForUserDeviceResponseBodySoftware {
	return s.Software
}

func (s *ListSoftwareForUserDeviceResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListSoftwareForUserDeviceResponseBody) SetRequestId(v string) *ListSoftwareForUserDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBody) SetSoftware(v []*ListSoftwareForUserDeviceResponseBodySoftware) *ListSoftwareForUserDeviceResponseBody {
	s.Software = v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBody) SetTotalNum(v int64) *ListSoftwareForUserDeviceResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBody) Validate() error {
	if s.Software != nil {
		for _, item := range s.Software {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSoftwareForUserDeviceResponseBodySoftware struct {
	// The software publisher.
	//
	// example:
	//
	// Alibaba (China) Network Technology Co.,Ltd.
	Inc *string `json:"Inc,omitempty" xml:"Inc,omitempty"`
	// The software installation time.
	//
	// example:
	//
	// 2023-08-18 02:43:02
	InstallTime *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	// The software name.
	//
	// example:
	//
	// DingTalk
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The collection of software versions.
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListSoftwareForUserDeviceResponseBodySoftware) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwareForUserDeviceResponseBodySoftware) GoString() string {
	return s.String()
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) GetInc() *string {
	return s.Inc
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) GetInstallTime() *string {
	return s.InstallTime
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) GetName() *string {
	return s.Name
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) GetVersions() []*string {
	return s.Versions
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) SetInc(v string) *ListSoftwareForUserDeviceResponseBodySoftware {
	s.Inc = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) SetInstallTime(v string) *ListSoftwareForUserDeviceResponseBodySoftware {
	s.InstallTime = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) SetName(v string) *ListSoftwareForUserDeviceResponseBodySoftware {
	s.Name = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) SetVersions(v []*string) *ListSoftwareForUserDeviceResponseBodySoftware {
	s.Versions = v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) Validate() error {
	return dara.Validate(s)
}
