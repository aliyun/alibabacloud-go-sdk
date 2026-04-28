// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportBizAlertInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertDescription(v string) *ReportBizAlertInfoRequest
	GetAlertDescription() *string
	SetAlertDetail(v string) *ReportBizAlertInfoRequest
	GetAlertDetail() *string
	SetAlertGrade(v string) *ReportBizAlertInfoRequest
	GetAlertGrade() *string
	SetAlertLabels(v string) *ReportBizAlertInfoRequest
	GetAlertLabels() *string
	SetAlertScene(v string) *ReportBizAlertInfoRequest
	GetAlertScene() *string
	SetAlertToken(v string) *ReportBizAlertInfoRequest
	GetAlertToken() *string
	SetAlertUid(v []*int64) *ReportBizAlertInfoRequest
	GetAlertUid() []*int64
	SetLanguage(v string) *ReportBizAlertInfoRequest
	GetLanguage() *string
}

type ReportBizAlertInfoRequest struct {
	AlertDescription *string `json:"AlertDescription,omitempty" xml:"AlertDescription,omitempty"`
	// This parameter is required.
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	AlertGrade  *string `json:"AlertGrade,omitempty" xml:"AlertGrade,omitempty"`
	AlertLabels *string `json:"AlertLabels,omitempty" xml:"AlertLabels,omitempty"`
	// This parameter is required.
	AlertScene *string `json:"AlertScene,omitempty" xml:"AlertScene,omitempty"`
	// This parameter is required.
	AlertToken *string  `json:"AlertToken,omitempty" xml:"AlertToken,omitempty"`
	AlertUid   []*int64 `json:"AlertUid,omitempty" xml:"AlertUid,omitempty" type:"Repeated"`
	Language   *string  `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ReportBizAlertInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportBizAlertInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoRequest) GetAlertDescription() *string {
	return s.AlertDescription
}

func (s *ReportBizAlertInfoRequest) GetAlertDetail() *string {
	return s.AlertDetail
}

func (s *ReportBizAlertInfoRequest) GetAlertGrade() *string {
	return s.AlertGrade
}

func (s *ReportBizAlertInfoRequest) GetAlertLabels() *string {
	return s.AlertLabels
}

func (s *ReportBizAlertInfoRequest) GetAlertScene() *string {
	return s.AlertScene
}

func (s *ReportBizAlertInfoRequest) GetAlertToken() *string {
	return s.AlertToken
}

func (s *ReportBizAlertInfoRequest) GetAlertUid() []*int64 {
	return s.AlertUid
}

func (s *ReportBizAlertInfoRequest) GetLanguage() *string {
	return s.Language
}

func (s *ReportBizAlertInfoRequest) SetAlertDescription(v string) *ReportBizAlertInfoRequest {
	s.AlertDescription = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertDetail(v string) *ReportBizAlertInfoRequest {
	s.AlertDetail = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertGrade(v string) *ReportBizAlertInfoRequest {
	s.AlertGrade = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertLabels(v string) *ReportBizAlertInfoRequest {
	s.AlertLabels = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertScene(v string) *ReportBizAlertInfoRequest {
	s.AlertScene = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertToken(v string) *ReportBizAlertInfoRequest {
	s.AlertToken = &v
	return s
}

func (s *ReportBizAlertInfoRequest) SetAlertUid(v []*int64) *ReportBizAlertInfoRequest {
	s.AlertUid = v
	return s
}

func (s *ReportBizAlertInfoRequest) SetLanguage(v string) *ReportBizAlertInfoRequest {
	s.Language = &v
	return s
}

func (s *ReportBizAlertInfoRequest) Validate() error {
	return dara.Validate(s)
}
