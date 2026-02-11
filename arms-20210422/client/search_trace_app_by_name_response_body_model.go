// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchTraceAppByNameResponseBody
	GetRequestId() *string
	SetTraceApps(v []*SearchTraceAppByNameResponseBodyTraceApps) *SearchTraceAppByNameResponseBody
	GetTraceApps() []*SearchTraceAppByNameResponseBodyTraceApps
}

type SearchTraceAppByNameResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceApps []*SearchTraceAppByNameResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s SearchTraceAppByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTraceAppByNameResponseBody) GetTraceApps() []*SearchTraceAppByNameResponseBodyTraceApps {
	return s.TraceApps
}

func (s *SearchTraceAppByNameResponseBody) SetRequestId(v string) *SearchTraceAppByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBody) SetTraceApps(v []*SearchTraceAppByNameResponseBodyTraceApps) *SearchTraceAppByNameResponseBody {
	s.TraceApps = v
	return s
}

func (s *SearchTraceAppByNameResponseBody) Validate() error {
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

type SearchTraceAppByNameResponseBodyTraceApps struct {
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

func (s SearchTraceAppByNameResponseBodyTraceApps) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetAppId() *int64 {
	return s.AppId
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetAppName() *string {
	return s.AppName
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetLabels() []*string {
	return s.Labels
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetPid() *string {
	return s.Pid
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetShow() *bool {
	return s.Show
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetType() *string {
	return s.Type
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) GetUserId() *string {
	return s.UserId
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppId(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppName(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetCreateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetLabels(v []*string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetPid(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetRegionId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetShow(v bool) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetType(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUpdateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUserId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UserId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) Validate() error {
	return dara.Validate(s)
}
