// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportBizAlertInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertDescription(v string) *ReportBizAlertInfoShrinkRequest
	GetAlertDescription() *string
	SetAlertDetail(v string) *ReportBizAlertInfoShrinkRequest
	GetAlertDetail() *string
	SetAlertGrade(v string) *ReportBizAlertInfoShrinkRequest
	GetAlertGrade() *string
	SetAlertLabels(v string) *ReportBizAlertInfoShrinkRequest
	GetAlertLabels() *string
	SetAlertScene(v string) *ReportBizAlertInfoShrinkRequest
	GetAlertScene() *string
	SetAlertToken(v string) *ReportBizAlertInfoShrinkRequest
	GetAlertToken() *string
	SetAlertUidShrink(v string) *ReportBizAlertInfoShrinkRequest
	GetAlertUidShrink() *string
	SetLanguage(v string) *ReportBizAlertInfoShrinkRequest
	GetLanguage() *string
}

type ReportBizAlertInfoShrinkRequest struct {
	AlertDescription *string `json:"AlertDescription,omitempty" xml:"AlertDescription,omitempty"`
	// This parameter is required.
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	AlertGrade  *string `json:"AlertGrade,omitempty" xml:"AlertGrade,omitempty"`
	AlertLabels *string `json:"AlertLabels,omitempty" xml:"AlertLabels,omitempty"`
	// This parameter is required.
	AlertScene *string `json:"AlertScene,omitempty" xml:"AlertScene,omitempty"`
	// This parameter is required.
	AlertToken     *string `json:"AlertToken,omitempty" xml:"AlertToken,omitempty"`
	AlertUidShrink *string `json:"AlertUid,omitempty" xml:"AlertUid,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ReportBizAlertInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportBizAlertInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoShrinkRequest) GetAlertDescription() *string {
	return s.AlertDescription
}

func (s *ReportBizAlertInfoShrinkRequest) GetAlertDetail() *string {
	return s.AlertDetail
}

func (s *ReportBizAlertInfoShrinkRequest) GetAlertGrade() *string {
	return s.AlertGrade
}

func (s *ReportBizAlertInfoShrinkRequest) GetAlertLabels() *string {
	return s.AlertLabels
}

func (s *ReportBizAlertInfoShrinkRequest) GetAlertScene() *string {
	return s.AlertScene
}

func (s *ReportBizAlertInfoShrinkRequest) GetAlertToken() *string {
	return s.AlertToken
}

func (s *ReportBizAlertInfoShrinkRequest) GetAlertUidShrink() *string {
	return s.AlertUidShrink
}

func (s *ReportBizAlertInfoShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertDescription(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertDescription = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertDetail(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertDetail = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertGrade(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertGrade = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertLabels(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertLabels = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertScene(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertScene = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertToken(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertToken = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetAlertUidShrink(v string) *ReportBizAlertInfoShrinkRequest {
	s.AlertUidShrink = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) SetLanguage(v string) *ReportBizAlertInfoShrinkRequest {
	s.Language = &v
	return s
}

func (s *ReportBizAlertInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
