// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchTraceAppByPageResponseBodyPageBean) *SearchTraceAppByPageResponseBody
	GetPageBean() *SearchTraceAppByPageResponseBodyPageBean
	SetRequestId(v string) *SearchTraceAppByPageResponseBody
	GetRequestId() *string
}

type SearchTraceAppByPageResponseBody struct {
	PageBean  *SearchTraceAppByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTraceAppByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBody) GetPageBean() *SearchTraceAppByPageResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchTraceAppByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTraceAppByPageResponseBody) SetPageBean(v *SearchTraceAppByPageResponseBodyPageBean) *SearchTraceAppByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTraceAppByPageResponseBody) SetRequestId(v string) *SearchTraceAppByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTraceAppByPageResponseBodyPageBean struct {
	PageNumber *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TraceApps  []*SearchTraceAppByPageResponseBodyPageBeanTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s SearchTraceAppByPageResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTraceAppByPageResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTraceAppByPageResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchTraceAppByPageResponseBodyPageBean) GetTraceApps() []*SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	return s.TraceApps
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetPageSize(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetTraceApps(v []*SearchTraceAppByPageResponseBodyPageBeanTraceApps) *SearchTraceAppByPageResponseBodyPageBean {
	s.TraceApps = v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) Validate() error {
	if s.TraceApps != nil {
		for _, item := range s.TraceApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTraceAppByPageResponseBodyPageBeanTraceApps struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetAppId() *int64 {
	return s.AppId
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetAppName() *string {
	return s.AppName
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetLabels() []*string {
	return s.Labels
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetPid() *string {
	return s.Pid
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetShow() *bool {
	return s.Show
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetType() *string {
	return s.Type
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) GetUserId() *string {
	return s.UserId
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppId(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppName(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetCreateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetLabels(v []*string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetPid(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetRegionId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetShow(v bool) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetType(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUpdateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUserId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UserId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) Validate() error {
	return dara.Validate(s)
}
