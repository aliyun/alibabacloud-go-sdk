// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTraceAppResponseBody
	GetRequestId() *string
	SetTraceApp(v *GetTraceAppResponseBodyTraceApp) *GetTraceAppResponseBody
	GetTraceApp() *GetTraceAppResponseBodyTraceApp
}

type GetTraceAppResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceApp  *GetTraceAppResponseBodyTraceApp `json:"TraceApp,omitempty" xml:"TraceApp,omitempty" type:"Struct"`
}

func (s GetTraceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTraceAppResponseBody) GetTraceApp() *GetTraceAppResponseBodyTraceApp {
	return s.TraceApp
}

func (s *GetTraceAppResponseBody) SetRequestId(v string) *GetTraceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceAppResponseBody) SetTraceApp(v *GetTraceAppResponseBodyTraceApp) *GetTraceAppResponseBody {
	s.TraceApp = v
	return s
}

func (s *GetTraceAppResponseBody) Validate() error {
	if s.TraceApp != nil {
		if err := s.TraceApp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTraceAppResponseBodyTraceApp struct {
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

func (s GetTraceAppResponseBodyTraceApp) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppResponseBodyTraceApp) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBodyTraceApp) GetAppId() *int64 {
	return s.AppId
}

func (s *GetTraceAppResponseBodyTraceApp) GetAppName() *string {
	return s.AppName
}

func (s *GetTraceAppResponseBodyTraceApp) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTraceAppResponseBodyTraceApp) GetLabels() []*string {
	return s.Labels
}

func (s *GetTraceAppResponseBodyTraceApp) GetPid() *string {
	return s.Pid
}

func (s *GetTraceAppResponseBodyTraceApp) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTraceAppResponseBodyTraceApp) GetShow() *bool {
	return s.Show
}

func (s *GetTraceAppResponseBodyTraceApp) GetType() *string {
	return s.Type
}

func (s *GetTraceAppResponseBodyTraceApp) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetTraceAppResponseBodyTraceApp) GetUserId() *string {
	return s.UserId
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppId(v int64) *GetTraceAppResponseBodyTraceApp {
	s.AppId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppName(v string) *GetTraceAppResponseBodyTraceApp {
	s.AppName = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetCreateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.CreateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetLabels(v []*string) *GetTraceAppResponseBodyTraceApp {
	s.Labels = v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetPid(v string) *GetTraceAppResponseBodyTraceApp {
	s.Pid = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetRegionId(v string) *GetTraceAppResponseBodyTraceApp {
	s.RegionId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetShow(v bool) *GetTraceAppResponseBodyTraceApp {
	s.Show = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetType(v string) *GetTraceAppResponseBodyTraceApp {
	s.Type = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUpdateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.UpdateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUserId(v string) *GetTraceAppResponseBodyTraceApp {
	s.UserId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) Validate() error {
	return dara.Validate(s)
}
