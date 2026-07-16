// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVersionsResponseBody
	GetRequestId() *string
	SetSiteVersionList(v []*ListVersionsResponseBodySiteVersionList) *ListVersionsResponseBody
	GetSiteVersionList() []*ListVersionsResponseBodySiteVersionList
}

type ListVersionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7EBEC214-805D-5FE9-AEED-258FE0F8850F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version list of the site.
	SiteVersionList []*ListVersionsResponseBodySiteVersionList `json:"SiteVersionList,omitempty" xml:"SiteVersionList,omitempty" type:"Repeated"`
}

func (s ListVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVersionsResponseBody) GetSiteVersionList() []*ListVersionsResponseBodySiteVersionList {
	return s.SiteVersionList
}

func (s *ListVersionsResponseBody) SetRequestId(v string) *ListVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVersionsResponseBody) SetSiteVersionList(v []*ListVersionsResponseBodySiteVersionList) *ListVersionsResponseBody {
	s.SiteVersionList = v
	return s
}

func (s *ListVersionsResponseBody) Validate() error {
	if s.SiteVersionList != nil {
		for _, item := range s.SiteVersionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVersionsResponseBodySiteVersionList struct {
	// The creation time.
	//
	// example:
	//
	// 2024-06-15T17:41:27+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// 站点A使用多个环境。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The environment list of the site version. The version may have no environment or one or more environments configured, such as the default environment or environment 2.
	EnvironmentNameList []*string `json:"EnvironmentNameList,omitempty" xml:"EnvironmentNameList,omitempty" type:"Repeated"`
	// The parent version of the site version.
	//
	// example:
	//
	// 1
	ParentSiteVersion *int32 `json:"ParentSiteVersion,omitempty" xml:"ParentSiteVersion,omitempty"`
	// Indicates whether the version is read-only. Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The site version.
	//
	// example:
	//
	// 8
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The status. Valid values:
	//
	// - **online**: active.
	//
	// - **configuring**: being configured.
	//
	// - **configure_faild**: configuration failed.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-06-15T17:41:27+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListVersionsResponseBodySiteVersionList) String() string {
	return dara.Prettify(s)
}

func (s ListVersionsResponseBodySiteVersionList) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBodySiteVersionList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVersionsResponseBodySiteVersionList) GetDescription() *string {
	return s.Description
}

func (s *ListVersionsResponseBodySiteVersionList) GetEnvironmentNameList() []*string {
	return s.EnvironmentNameList
}

func (s *ListVersionsResponseBodySiteVersionList) GetParentSiteVersion() *int32 {
	return s.ParentSiteVersion
}

func (s *ListVersionsResponseBodySiteVersionList) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListVersionsResponseBodySiteVersionList) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListVersionsResponseBodySiteVersionList) GetStatus() *string {
	return s.Status
}

func (s *ListVersionsResponseBodySiteVersionList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListVersionsResponseBodySiteVersionList) SetCreateTime(v string) *ListVersionsResponseBodySiteVersionList {
	s.CreateTime = &v
	return s
}

func (s *ListVersionsResponseBodySiteVersionList) SetDescription(v string) *ListVersionsResponseBodySiteVersionList {
	s.Description = &v
	return s
}

func (s *ListVersionsResponseBodySiteVersionList) SetEnvironmentNameList(v []*string) *ListVersionsResponseBodySiteVersionList {
	s.EnvironmentNameList = v
	return s
}

func (s *ListVersionsResponseBodySiteVersionList) SetParentSiteVersion(v int32) *ListVersionsResponseBodySiteVersionList {
	s.ParentSiteVersion = &v
	return s
}

func (s *ListVersionsResponseBodySiteVersionList) SetReadOnly(v bool) *ListVersionsResponseBodySiteVersionList {
	s.ReadOnly = &v
	return s
}

func (s *ListVersionsResponseBodySiteVersionList) SetSiteVersion(v int32) *ListVersionsResponseBodySiteVersionList {
	s.SiteVersion = &v
	return s
}

func (s *ListVersionsResponseBodySiteVersionList) SetStatus(v string) *ListVersionsResponseBodySiteVersionList {
	s.Status = &v
	return s
}

func (s *ListVersionsResponseBodySiteVersionList) SetUpdateTime(v string) *ListVersionsResponseBodySiteVersionList {
	s.UpdateTime = &v
	return s
}

func (s *ListVersionsResponseBodySiteVersionList) Validate() error {
	return dara.Validate(s)
}
