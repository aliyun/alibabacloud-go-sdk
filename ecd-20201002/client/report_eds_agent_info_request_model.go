// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportEdsAgentInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *ReportEdsAgentInfoRequest
	GetAliUid() *int64
	SetDesktopId(v string) *ReportEdsAgentInfoRequest
	GetDesktopId() *string
	SetEcsInstanceId(v string) *ReportEdsAgentInfoRequest
	GetEcsInstanceId() *string
	SetEdsAgentInfo(v string) *ReportEdsAgentInfoRequest
	GetEdsAgentInfo() *string
}

type ReportEdsAgentInfoRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EdsAgentInfo  *string `json:"EdsAgentInfo,omitempty" xml:"EdsAgentInfo,omitempty"`
}

func (s ReportEdsAgentInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportEdsAgentInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoRequest) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ReportEdsAgentInfoRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ReportEdsAgentInfoRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *ReportEdsAgentInfoRequest) GetEdsAgentInfo() *string {
	return s.EdsAgentInfo
}

func (s *ReportEdsAgentInfoRequest) SetAliUid(v int64) *ReportEdsAgentInfoRequest {
	s.AliUid = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetDesktopId(v string) *ReportEdsAgentInfoRequest {
	s.DesktopId = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetEcsInstanceId(v string) *ReportEdsAgentInfoRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetEdsAgentInfo(v string) *ReportEdsAgentInfoRequest {
	s.EdsAgentInfo = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) Validate() error {
	return dara.Validate(s)
}
