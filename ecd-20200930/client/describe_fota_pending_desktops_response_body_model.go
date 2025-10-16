// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFotaPendingDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFotaPendingDesktopsResponseBody
	GetCode() *string
	SetFotaPendingDesktops(v []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) *DescribeFotaPendingDesktopsResponseBody
	GetFotaPendingDesktops() []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops
	SetMessage(v string) *DescribeFotaPendingDesktopsResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeFotaPendingDesktopsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeFotaPendingDesktopsResponseBody
	GetRequestId() *string
}

type DescribeFotaPendingDesktopsResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cloud computers whose images can be and are pending to be updated to the version specified in `TaskUid`.
	FotaPendingDesktops []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops `json:"FotaPendingDesktops,omitempty" xml:"FotaPendingDesktops,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFotaPendingDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaPendingDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFotaPendingDesktopsResponseBody) GetFotaPendingDesktops() []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	return s.FotaPendingDesktops
}

func (s *DescribeFotaPendingDesktopsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFotaPendingDesktopsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFotaPendingDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFotaPendingDesktopsResponseBody) SetCode(v string) *DescribeFotaPendingDesktopsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBody) SetFotaPendingDesktops(v []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) *DescribeFotaPendingDesktopsResponseBody {
	s.FotaPendingDesktops = v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBody) SetMessage(v string) *DescribeFotaPendingDesktopsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBody) SetNextToken(v string) *DescribeFotaPendingDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBody) SetRequestId(v string) *DescribeFotaPendingDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBody) Validate() error {
	if s.FotaPendingDesktops != nil {
		for _, item := range s.FotaPendingDesktops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops struct {
	// The current version of the image used by the cloud computer.
	//
	// example:
	//
	// 0.0.1-D-20220513.143129
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-bvdtu3jn97o1r****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// TestDesktop
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// To be hidden
	FotaProject *string `json:"FotaProject,omitempty" xml:"FotaProject,omitempty"`
	// The ID of the office network.
	//
	// example:
	//
	// cn-hangzhou+dir-815419****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The connected sessions.
	Sessions []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The status of the cloud computer.
	//
	// Valid values:
	//
	// 	- 0: The cloud computer is being created.
	//
	// 	- 1: The cloud computer is being started.
	//
	// 	- 2: The cloud computer is running.
	//
	// 	- 3: The cloud computer is being stopped.
	//
	// 	- 5: The cloud computer is stopped.
	//
	// 	- 6: The cloud computer expires.
	//
	// 	- 7: The cloud computer is deleted.
	//
	// 	- 9: Failed to create the cloud computer.
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GetCurrentAppVersion() *string {
	return s.CurrentAppVersion
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GetFotaProject() *string {
	return s.FotaProject
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GetSessions() []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions {
	return s.Sessions
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetCurrentAppVersion(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetDesktopId(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetDesktopName(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetFotaProject(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.FotaProject = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetOfficeSiteId(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetSessions(v []*DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) SetStatus(v int64) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops {
	s.Status = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktops) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions struct {
	// The ID of the end user that connects to the cloud computer.
	//
	// example:
	//
	// end user id
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
}

func (s DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions) SetEndUserId(v string) *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeFotaPendingDesktopsResponseBodyFotaPendingDesktopsSessions) Validate() error {
	return dara.Validate(s)
}
