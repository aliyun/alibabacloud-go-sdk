// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRetcodeAppByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchRetcodeAppByPageResponseBodyPageBean) *SearchRetcodeAppByPageResponseBody
	GetPageBean() *SearchRetcodeAppByPageResponseBodyPageBean
	SetRequestId(v string) *SearchRetcodeAppByPageResponseBody
	GetRequestId() *string
}

type SearchRetcodeAppByPageResponseBody struct {
	PageBean  *SearchRetcodeAppByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBody) GetPageBean() *SearchRetcodeAppByPageResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchRetcodeAppByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchRetcodeAppByPageResponseBody) SetPageBean(v *SearchRetcodeAppByPageResponseBodyPageBean) *SearchRetcodeAppByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBody) SetRequestId(v string) *SearchRetcodeAppByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchRetcodeAppByPageResponseBodyPageBean struct {
	PageNumber  *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RetcodeApps []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
	TotalCount  *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) GetRetcodeApps() []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	return s.RetcodeApps
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetPageSize(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetRetcodeApps(v []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.RetcodeApps = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) Validate() error {
	if s.RetcodeApps != nil {
		for _, item := range s.RetcodeApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps struct {
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Pid        *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetAppId() *int64 {
	return s.AppId
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetAppName() *string {
	return s.AppName
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetPid() *string {
	return s.Pid
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetType() *string {
	return s.Type
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GetUserId() *string {
	return s.UserId
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppId(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppName(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppName = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetCreateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.CreateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetPid(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Pid = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetRegionId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.RegionId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetType(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Type = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUpdateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUserId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UserId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) Validate() error {
	return dara.Validate(s)
}
