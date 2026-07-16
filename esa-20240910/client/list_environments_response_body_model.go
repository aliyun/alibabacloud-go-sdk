// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentList(v []*ListEnvironmentsResponseBodyEnvironmentList) *ListEnvironmentsResponseBody
	GetEnvironmentList() []*ListEnvironmentsResponseBodyEnvironmentList
	SetRequestId(v string) *ListEnvironmentsResponseBody
	GetRequestId() *string
}

type ListEnvironmentsResponseBody struct {
	// The list of environments.
	EnvironmentList []*ListEnvironmentsResponseBodyEnvironmentList `json:"EnvironmentList,omitempty" xml:"EnvironmentList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) GetEnvironmentList() []*ListEnvironmentsResponseBodyEnvironmentList {
	return s.EnvironmentList
}

func (s *ListEnvironmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentsResponseBody) SetEnvironmentList(v []*ListEnvironmentsResponseBodyEnvironmentList) *ListEnvironmentsResponseBody {
	s.EnvironmentList = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetRequestId(v string) *ListEnvironmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentsResponseBody) Validate() error {
	if s.EnvironmentList != nil {
		for _, item := range s.EnvironmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnvironmentsResponseBodyEnvironmentList struct {
	// The creation time.
	//
	// example:
	//
	// 2024-08-15T08:15:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The environment name.
	//
	// example:
	//
	// yichao-test-yctest
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// Indicates whether this is the default environment.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The previous version number.
	//
	// example:
	//
	// -1
	PreSiteVersion *int32 `json:"PreSiteVersion,omitempty" xml:"PreSiteVersion,omitempty"`
	// The priority.
	//
	// example:
	//
	// 56
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Indicates whether the environment is read-only.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The environment rule.
	//
	// example:
	//
	// ("ip" eq "1.1.1.1")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The site version number.
	//
	// example:
	//
	// 8
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2023-09-15T08:16:04Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEnvironmentsResponseBodyEnvironmentList) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBodyEnvironmentList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetPreSiteVersion() *int32 {
	return s.PreSiteVersion
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetRule() *string {
	return s.Rule
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetCreateTime(v string) *ListEnvironmentsResponseBodyEnvironmentList {
	s.CreateTime = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetEnvironmentName(v string) *ListEnvironmentsResponseBodyEnvironmentList {
	s.EnvironmentName = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetIsDefault(v bool) *ListEnvironmentsResponseBodyEnvironmentList {
	s.IsDefault = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetPreSiteVersion(v int32) *ListEnvironmentsResponseBodyEnvironmentList {
	s.PreSiteVersion = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetPriority(v int32) *ListEnvironmentsResponseBodyEnvironmentList {
	s.Priority = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetReadOnly(v bool) *ListEnvironmentsResponseBodyEnvironmentList {
	s.ReadOnly = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetRule(v string) *ListEnvironmentsResponseBodyEnvironmentList {
	s.Rule = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetSiteVersion(v int32) *ListEnvironmentsResponseBodyEnvironmentList {
	s.SiteVersion = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) SetUpdateTime(v string) *ListEnvironmentsResponseBodyEnvironmentList {
	s.UpdateTime = &v
	return s
}

func (s *ListEnvironmentsResponseBodyEnvironmentList) Validate() error {
	return dara.Validate(s)
}
