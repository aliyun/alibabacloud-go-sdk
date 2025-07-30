// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEdsAgentReportConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *QueryEdsAgentReportConfigRequest
	GetAliUid() *int64
	SetDesktopId(v string) *QueryEdsAgentReportConfigRequest
	GetDesktopId() *string
	SetEcsInstanceId(v string) *QueryEdsAgentReportConfigRequest
	GetEcsInstanceId() *string
}

type QueryEdsAgentReportConfigRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
}

func (s QueryEdsAgentReportConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEdsAgentReportConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigRequest) GetAliUid() *int64 {
	return s.AliUid
}

func (s *QueryEdsAgentReportConfigRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *QueryEdsAgentReportConfigRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *QueryEdsAgentReportConfigRequest) SetAliUid(v int64) *QueryEdsAgentReportConfigRequest {
	s.AliUid = &v
	return s
}

func (s *QueryEdsAgentReportConfigRequest) SetDesktopId(v string) *QueryEdsAgentReportConfigRequest {
	s.DesktopId = &v
	return s
}

func (s *QueryEdsAgentReportConfigRequest) SetEcsInstanceId(v string) *QueryEdsAgentReportConfigRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *QueryEdsAgentReportConfigRequest) Validate() error {
	return dara.Validate(s)
}
