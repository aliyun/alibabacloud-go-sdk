// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTraceAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListTraceAppsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListTraceAppsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTraceAppsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTraceAppsResponseBody
	GetSuccess() *bool
	SetTraceApps(v []*ListTraceAppsResponseBodyTraceApps) *ListTraceAppsResponseBody
	GetTraceApps() []*ListTraceAppsResponseBodyTraceApps
}

type ListTraceAppsResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceApps []*ListTraceAppsResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s ListTraceAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTraceAppsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTraceAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTraceAppsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTraceAppsResponseBody) GetTraceApps() []*ListTraceAppsResponseBodyTraceApps {
	return s.TraceApps
}

func (s *ListTraceAppsResponseBody) SetCode(v int32) *ListTraceAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetMessage(v string) *ListTraceAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetRequestId(v string) *ListTraceAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetSuccess(v bool) *ListTraceAppsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetTraceApps(v []*ListTraceAppsResponseBodyTraceApps) *ListTraceAppsResponseBody {
	s.TraceApps = v
	return s
}

func (s *ListTraceAppsResponseBody) Validate() error {
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

type ListTraceAppsResponseBodyTraceApps struct {
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

func (s ListTraceAppsResponseBodyTraceApps) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBodyTraceApps) GetAppId() *int64 {
	return s.AppId
}

func (s *ListTraceAppsResponseBodyTraceApps) GetAppName() *string {
	return s.AppName
}

func (s *ListTraceAppsResponseBodyTraceApps) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTraceAppsResponseBodyTraceApps) GetLabels() []*string {
	return s.Labels
}

func (s *ListTraceAppsResponseBodyTraceApps) GetPid() *string {
	return s.Pid
}

func (s *ListTraceAppsResponseBodyTraceApps) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTraceAppsResponseBodyTraceApps) GetShow() *bool {
	return s.Show
}

func (s *ListTraceAppsResponseBodyTraceApps) GetType() *string {
	return s.Type
}

func (s *ListTraceAppsResponseBodyTraceApps) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListTraceAppsResponseBodyTraceApps) GetUserId() *string {
	return s.UserId
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppId(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppName(v string) *ListTraceAppsResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetCreateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetLabels(v []*string) *ListTraceAppsResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetPid(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetRegionId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetShow(v bool) *ListTraceAppsResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetType(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUpdateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUserId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.UserId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) Validate() error {
	return dara.Validate(s)
}
