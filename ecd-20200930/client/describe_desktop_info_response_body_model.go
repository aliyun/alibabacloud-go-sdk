// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktops(v []*DescribeDesktopInfoResponseBodyDesktops) *DescribeDesktopInfoResponseBody
	GetDesktops() []*DescribeDesktopInfoResponseBodyDesktops
	SetRequestId(v string) *DescribeDesktopInfoResponseBody
	GetRequestId() *string
}

type DescribeDesktopInfoResponseBody struct {
	// The basic information about cloud computers.
	Desktops []*DescribeDesktopInfoResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 062B1439-709A-580E-85DF-CE97A1560565
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopInfoResponseBody) GetDesktops() []*DescribeDesktopInfoResponseBodyDesktops {
	return s.Desktops
}

func (s *DescribeDesktopInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopInfoResponseBody) SetDesktops(v []*DescribeDesktopInfoResponseBodyDesktops) *DescribeDesktopInfoResponseBody {
	s.Desktops = v
	return s
}

func (s *DescribeDesktopInfoResponseBody) SetRequestId(v string) *DescribeDesktopInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopInfoResponseBody) Validate() error {
	if s.Desktops != nil {
		for _, item := range s.Desktops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopInfoResponseBodyDesktops struct {
	// The connection status of the user.
	//
	// Valid values:
	//
	// 	- Connected
	//
	// 	- Disconnected
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The version of the cloud computer image.
	//
	// example:
	//
	// 1.4.0-R-***
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	// The ID of the cloud computer pool.
	//
	// example:
	//
	// dg-3uiojcc0j4kh7****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The status of the cloud computer.
	//
	// Valid values:
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Expired
	//
	// 	- Deleted
	//
	// 	- Pending
	//
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// The information about flags that are used to manage cloud computers.
	ManagementFlag []*string `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty" type:"Repeated"`
	// The size of the update package. Unit: KB.
	//
	// example:
	//
	// 568533470
	NewAppSize *int64 `json:"NewAppSize,omitempty" xml:"NewAppSize,omitempty"`
	// The version number of the image that can be updated on the cloud computer.
	//
	// example:
	//
	// 1.6.0-R-***
	NewAppVersion *string `json:"NewAppVersion,omitempty" xml:"NewAppVersion,omitempty"`
	// The description of the image version that can be updated.
	//
	// example:
	//
	// Test package 03-07
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The time when the cloud computer was first started.
	//
	// example:
	//
	// 2020-11-06T08:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDesktopInfoResponseBodyDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopInfoResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetCurrentAppVersion() *string {
	return s.CurrentAppVersion
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetManagementFlag() []*string {
	return s.ManagementFlag
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetNewAppSize() *int64 {
	return s.NewAppSize
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetNewAppVersion() *string {
	return s.NewAppVersion
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDesktopInfoResponseBodyDesktops) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetConnectionStatus(v string) *DescribeDesktopInfoResponseBodyDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetCurrentAppVersion(v string) *DescribeDesktopInfoResponseBodyDesktops {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetDesktopGroupId(v string) *DescribeDesktopInfoResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetDesktopId(v string) *DescribeDesktopInfoResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetDesktopStatus(v string) *DescribeDesktopInfoResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetManagementFlag(v []*string) *DescribeDesktopInfoResponseBodyDesktops {
	s.ManagementFlag = v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetNewAppSize(v int64) *DescribeDesktopInfoResponseBodyDesktops {
	s.NewAppSize = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetNewAppVersion(v string) *DescribeDesktopInfoResponseBodyDesktops {
	s.NewAppVersion = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetReleaseNote(v string) *DescribeDesktopInfoResponseBodyDesktops {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) SetStartTime(v string) *DescribeDesktopInfoResponseBodyDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopInfoResponseBodyDesktops) Validate() error {
	return dara.Validate(s)
}
