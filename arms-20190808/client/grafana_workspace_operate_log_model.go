// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceOperateLog interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v float32) *GrafanaWorkspaceOperateLog
	GetDate() *float32
	SetDetail(v string) *GrafanaWorkspaceOperateLog
	GetDetail() *string
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceOperateLog
	GetGrafanaWorkspaceId() *string
	SetId(v int64) *GrafanaWorkspaceOperateLog
	GetId() *int64
	SetOperatorId(v string) *GrafanaWorkspaceOperateLog
	GetOperatorId() *string
}

type GrafanaWorkspaceOperateLog struct {
	Date               *float32 `json:"date,omitempty" xml:"date,omitempty"`
	Detail             *string  `json:"detail,omitempty" xml:"detail,omitempty"`
	GrafanaWorkspaceId *string  `json:"grafanaWorkspaceId,omitempty" xml:"grafanaWorkspaceId,omitempty"`
	Id                 *int64   `json:"id,omitempty" xml:"id,omitempty"`
	OperatorId         *string  `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s GrafanaWorkspaceOperateLog) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceOperateLog) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceOperateLog) GetDate() *float32 {
	return s.Date
}

func (s *GrafanaWorkspaceOperateLog) GetDetail() *string {
	return s.Detail
}

func (s *GrafanaWorkspaceOperateLog) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspaceOperateLog) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceOperateLog) GetOperatorId() *string {
	return s.OperatorId
}

func (s *GrafanaWorkspaceOperateLog) SetDate(v float32) *GrafanaWorkspaceOperateLog {
	s.Date = &v
	return s
}

func (s *GrafanaWorkspaceOperateLog) SetDetail(v string) *GrafanaWorkspaceOperateLog {
	s.Detail = &v
	return s
}

func (s *GrafanaWorkspaceOperateLog) SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceOperateLog {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspaceOperateLog) SetId(v int64) *GrafanaWorkspaceOperateLog {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceOperateLog) SetOperatorId(v string) *GrafanaWorkspaceOperateLog {
	s.OperatorId = &v
	return s
}

func (s *GrafanaWorkspaceOperateLog) Validate() error {
	return dara.Validate(s)
}
