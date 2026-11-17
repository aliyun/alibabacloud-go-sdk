// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody
	GetApplications() []*ListApplicationsResponseBodyApplications
	SetPageNo(v int64) *ListApplicationsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListApplicationsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationsResponseBody
	GetTotalCount() *int64
}

type ListApplicationsResponseBody struct {
	// The list of applications. In the Chinese environment, the applications are sorted in descending order of group name and ascending order of application level.
	Applications []*ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
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
	// The total number of applications that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetApplications() []*ListApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListApplicationsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationsResponseBody) SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetPageNo(v int64) *ListApplicationsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListApplicationsResponseBody) SetPageSize(v int64) *ListApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalCount(v int64) *ListApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
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

type ListApplicationsResponseBodyApplications struct {
	// The application ID in the built-in application signature library.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application name. The name is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// DingTalk
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application category ID.
	//
	// example:
	//
	// 8
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The application category name. The name is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// Office collaboration
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The default action for the application in the built-in application library. The value is fixed as Alert (Monitor).
	//
	// example:
	//
	// Alert
	DefaultAction *string `json:"DefaultAction,omitempty" xml:"DefaultAction,omitempty"`
	// The application description. The description is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// An enterprise-grade intelligent mobile office and collaboration platform that provides features such as instant messaging, audio and video conferencing, organizational management, online documents, approval and attendance, and AI office assistants
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application group name.
	//
	// example:
	//
	// Dingtalk
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The level value of the application. The value is obtained from the built-in application signature library and used for list sorting.
	//
	// example:
	//
	// 700
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The list of protocols used by the application.
	Protocol []*string `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Repeated"`
	// The list of detection prerequisites for the application.
	Requirements []*string `json:"Requirements,omitempty" xml:"Requirements,omitempty" type:"Repeated"`
	// The risk level. Valid values:
	//
	// - **Information**: Informational.
	//
	// - **Low**: Low.
	//
	// - **Medium**: Medium.
	//
	// - **High**: High.
	//
	// - **Critical**: Critical.
	//
	// example:
	//
	// Low
	Risk *string `json:"Risk,omitempty" xml:"Risk,omitempty"`
	// The technology type name. The name is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// Browser-based
	Technology *string `json:"Technology,omitempty" xml:"Technology,omitempty"`
	// The last update time of the application information. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731465600
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The application vendor name. The name is returned in Chinese or English based on the Lang parameter.
	//
	// example:
	//
	// DingTalk Technology Co., Ltd
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationsResponseBodyApplications) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ListApplicationsResponseBodyApplications) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListApplicationsResponseBodyApplications) GetDefaultAction() *string {
	return s.DefaultAction
}

func (s *ListApplicationsResponseBodyApplications) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsResponseBodyApplications) GetGroupName() *string {
	return s.GroupName
}

func (s *ListApplicationsResponseBodyApplications) GetLevel() *int64 {
	return s.Level
}

func (s *ListApplicationsResponseBodyApplications) GetProtocol() []*string {
	return s.Protocol
}

func (s *ListApplicationsResponseBodyApplications) GetRequirements() []*string {
	return s.Requirements
}

func (s *ListApplicationsResponseBodyApplications) GetRisk() *string {
	return s.Risk
}

func (s *ListApplicationsResponseBodyApplications) GetTechnology() *string {
	return s.Technology
}

func (s *ListApplicationsResponseBodyApplications) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListApplicationsResponseBodyApplications) GetVendor() *string {
	return s.Vendor
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationId(v int64) *ListApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationName(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCategoryId(v int64) *ListApplicationsResponseBodyApplications {
	s.CategoryId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCategoryName(v string) *ListApplicationsResponseBodyApplications {
	s.CategoryName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetDefaultAction(v string) *ListApplicationsResponseBodyApplications {
	s.DefaultAction = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetDescription(v string) *ListApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetGroupName(v string) *ListApplicationsResponseBodyApplications {
	s.GroupName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetLevel(v int64) *ListApplicationsResponseBodyApplications {
	s.Level = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetProtocol(v []*string) *ListApplicationsResponseBodyApplications {
	s.Protocol = v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetRequirements(v []*string) *ListApplicationsResponseBodyApplications {
	s.Requirements = v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetRisk(v string) *ListApplicationsResponseBodyApplications {
	s.Risk = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetTechnology(v string) *ListApplicationsResponseBodyApplications {
	s.Technology = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetUpdatedTime(v int64) *ListApplicationsResponseBodyApplications {
	s.UpdatedTime = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetVendor(v string) *ListApplicationsResponseBodyApplications {
	s.Vendor = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
