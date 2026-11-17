// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationUpdateDynamicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamics(v []*ListApplicationUpdateDynamicsResponseBodyDynamics) *ListApplicationUpdateDynamicsResponseBody
	GetDynamics() []*ListApplicationUpdateDynamicsResponseBodyDynamics
	SetPageNo(v int64) *ListApplicationUpdateDynamicsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListApplicationUpdateDynamicsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListApplicationUpdateDynamicsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationUpdateDynamicsResponseBody
	GetTotalCount() *int64
}

type ListApplicationUpdateDynamicsResponseBody struct {
	// The list of update records for the application signature library, sorted in descending order of update time.
	Dynamics []*ListApplicationUpdateDynamicsResponseBodyDynamics `json:"Dynamics,omitempty" xml:"Dynamics,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of update records that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationUpdateDynamicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationUpdateDynamicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationUpdateDynamicsResponseBody) GetDynamics() []*ListApplicationUpdateDynamicsResponseBodyDynamics {
	return s.Dynamics
}

func (s *ListApplicationUpdateDynamicsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListApplicationUpdateDynamicsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationUpdateDynamicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationUpdateDynamicsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationUpdateDynamicsResponseBody) SetDynamics(v []*ListApplicationUpdateDynamicsResponseBodyDynamics) *ListApplicationUpdateDynamicsResponseBody {
	s.Dynamics = v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBody) SetPageNo(v int64) *ListApplicationUpdateDynamicsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBody) SetPageSize(v int64) *ListApplicationUpdateDynamicsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBody) SetRequestId(v string) *ListApplicationUpdateDynamicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBody) SetTotalCount(v int64) *ListApplicationUpdateDynamicsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBody) Validate() error {
	if s.Dynamics != nil {
		for _, item := range s.Dynamics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationUpdateDynamicsResponseBodyDynamics struct {
	// The list of applications added or updated in this update.
	Applications []*ListApplicationUpdateDynamicsResponseBodyDynamicsApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The update description. The description is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// Added 1 application signature library, such as DingTalk
	DynamicInfo *string `json:"DynamicInfo,omitempty" xml:"DynamicInfo,omitempty"`
	// The time when the signature library was updated. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731465600
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListApplicationUpdateDynamicsResponseBodyDynamics) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationUpdateDynamicsResponseBodyDynamics) GoString() string {
	return s.String()
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamics) GetApplications() []*ListApplicationUpdateDynamicsResponseBodyDynamicsApplications {
	return s.Applications
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamics) GetDynamicInfo() *string {
	return s.DynamicInfo
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamics) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamics) SetApplications(v []*ListApplicationUpdateDynamicsResponseBodyDynamicsApplications) *ListApplicationUpdateDynamicsResponseBodyDynamics {
	s.Applications = v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamics) SetDynamicInfo(v string) *ListApplicationUpdateDynamicsResponseBodyDynamics {
	s.DynamicInfo = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamics) SetUpdateTime(v int64) *ListApplicationUpdateDynamicsResponseBodyDynamics {
	s.UpdateTime = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamics) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationUpdateDynamicsResponseBodyDynamicsApplications struct {
	// The ID of the application in the built-in application signature library.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the application. The name is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// DingTalk
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
}

func (s ListApplicationUpdateDynamicsResponseBodyDynamicsApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationUpdateDynamicsResponseBodyDynamicsApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamicsApplications) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamicsApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamicsApplications) SetApplicationId(v int64) *ListApplicationUpdateDynamicsResponseBodyDynamicsApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamicsApplications) SetApplicationName(v string) *ListApplicationUpdateDynamicsResponseBodyDynamicsApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponseBodyDynamicsApplications) Validate() error {
	return dara.Validate(s)
}
