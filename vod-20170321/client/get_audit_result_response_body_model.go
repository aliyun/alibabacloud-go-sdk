// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIAuditResult(v *GetAuditResultResponseBodyAIAuditResult) *GetAuditResultResponseBody
	GetAIAuditResult() *GetAuditResultResponseBodyAIAuditResult
	SetRequestId(v string) *GetAuditResultResponseBody
	GetRequestId() *string
}

type GetAuditResultResponseBody struct {
	AIAuditResult *GetAuditResultResponseBodyAIAuditResult `json:"AIAuditResult,omitempty" xml:"AIAuditResult,omitempty" type:"Struct"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuditResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditResultResponseBody) GetAIAuditResult() *GetAuditResultResponseBodyAIAuditResult {
	return s.AIAuditResult
}

func (s *GetAuditResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuditResultResponseBody) SetAIAuditResult(v *GetAuditResultResponseBodyAIAuditResult) *GetAuditResultResponseBody {
	s.AIAuditResult = v
	return s
}

func (s *GetAuditResultResponseBody) SetRequestId(v string) *GetAuditResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAuditResultResponseBodyAIAuditResult struct {
	AbnormalModules *string `json:"AbnormalModules,omitempty" xml:"AbnormalModules,omitempty"`
	CoverResult     *string `json:"CoverResult,omitempty" xml:"CoverResult,omitempty"`
	ImageResult     *string `json:"ImageResult,omitempty" xml:"ImageResult,omitempty"`
	Label           *string `json:"Label,omitempty" xml:"Label,omitempty"`
	PornResult      *string `json:"PornResult,omitempty" xml:"PornResult,omitempty"`
	TerrorismResult *string `json:"TerrorismResult,omitempty" xml:"TerrorismResult,omitempty"`
	TitleResult     *string `json:"TitleResult,omitempty" xml:"TitleResult,omitempty"`
}

func (s GetAuditResultResponseBodyAIAuditResult) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultResponseBodyAIAuditResult) GoString() string {
	return s.String()
}

func (s *GetAuditResultResponseBodyAIAuditResult) GetAbnormalModules() *string {
	return s.AbnormalModules
}

func (s *GetAuditResultResponseBodyAIAuditResult) GetCoverResult() *string {
	return s.CoverResult
}

func (s *GetAuditResultResponseBodyAIAuditResult) GetImageResult() *string {
	return s.ImageResult
}

func (s *GetAuditResultResponseBodyAIAuditResult) GetLabel() *string {
	return s.Label
}

func (s *GetAuditResultResponseBodyAIAuditResult) GetPornResult() *string {
	return s.PornResult
}

func (s *GetAuditResultResponseBodyAIAuditResult) GetTerrorismResult() *string {
	return s.TerrorismResult
}

func (s *GetAuditResultResponseBodyAIAuditResult) GetTitleResult() *string {
	return s.TitleResult
}

func (s *GetAuditResultResponseBodyAIAuditResult) SetAbnormalModules(v string) *GetAuditResultResponseBodyAIAuditResult {
	s.AbnormalModules = &v
	return s
}

func (s *GetAuditResultResponseBodyAIAuditResult) SetCoverResult(v string) *GetAuditResultResponseBodyAIAuditResult {
	s.CoverResult = &v
	return s
}

func (s *GetAuditResultResponseBodyAIAuditResult) SetImageResult(v string) *GetAuditResultResponseBodyAIAuditResult {
	s.ImageResult = &v
	return s
}

func (s *GetAuditResultResponseBodyAIAuditResult) SetLabel(v string) *GetAuditResultResponseBodyAIAuditResult {
	s.Label = &v
	return s
}

func (s *GetAuditResultResponseBodyAIAuditResult) SetPornResult(v string) *GetAuditResultResponseBodyAIAuditResult {
	s.PornResult = &v
	return s
}

func (s *GetAuditResultResponseBodyAIAuditResult) SetTerrorismResult(v string) *GetAuditResultResponseBodyAIAuditResult {
	s.TerrorismResult = &v
	return s
}

func (s *GetAuditResultResponseBodyAIAuditResult) SetTitleResult(v string) *GetAuditResultResponseBodyAIAuditResult {
	s.TitleResult = &v
	return s
}

func (s *GetAuditResultResponseBodyAIAuditResult) Validate() error {
	return dara.Validate(s)
}
