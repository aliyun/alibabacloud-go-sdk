// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecurityCheckReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySecurityCheckReportResponseBody
	GetCode() *string
	SetData(v *QuerySecurityCheckReportResponseBodyData) *QuerySecurityCheckReportResponseBody
	GetData() *QuerySecurityCheckReportResponseBodyData
	SetMessage(v string) *QuerySecurityCheckReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySecurityCheckReportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySecurityCheckReportResponseBody
	GetSuccess() *bool
}

type QuerySecurityCheckReportResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QuerySecurityCheckReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySecurityCheckReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySecurityCheckReportResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecurityCheckReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySecurityCheckReportResponseBody) GetData() *QuerySecurityCheckReportResponseBodyData {
	return s.Data
}

func (s *QuerySecurityCheckReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySecurityCheckReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySecurityCheckReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySecurityCheckReportResponseBody) SetCode(v string) *QuerySecurityCheckReportResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBody) SetData(v *QuerySecurityCheckReportResponseBodyData) *QuerySecurityCheckReportResponseBody {
	s.Data = v
	return s
}

func (s *QuerySecurityCheckReportResponseBody) SetMessage(v string) *QuerySecurityCheckReportResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBody) SetRequestId(v string) *QuerySecurityCheckReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBody) SetSuccess(v bool) *QuerySecurityCheckReportResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySecurityCheckReportResponseBodyData struct {
	CloudSecurityGuide *int32  `json:"CloudSecurityGuide,omitempty" xml:"CloudSecurityGuide,omitempty"`
	ConfigCheckNumber  *int32  `json:"ConfigCheckNumber,omitempty" xml:"ConfigCheckNumber,omitempty"`
	ContactCheckNumber *int32  `json:"ContactCheckNumber,omitempty" xml:"ContactCheckNumber,omitempty"`
	RiskEventNumber    *int32  `json:"RiskEventNumber,omitempty" xml:"RiskEventNumber,omitempty"`
	SasCheckNumber     *int32  `json:"SasCheckNumber,omitempty" xml:"SasCheckNumber,omitempty"`
	SecurityStatus     *int32  `json:"SecurityStatus,omitempty" xml:"SecurityStatus,omitempty"`
	SuggestionText     *string `json:"SuggestionText,omitempty" xml:"SuggestionText,omitempty"`
}

func (s QuerySecurityCheckReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySecurityCheckReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySecurityCheckReportResponseBodyData) GetCloudSecurityGuide() *int32 {
	return s.CloudSecurityGuide
}

func (s *QuerySecurityCheckReportResponseBodyData) GetConfigCheckNumber() *int32 {
	return s.ConfigCheckNumber
}

func (s *QuerySecurityCheckReportResponseBodyData) GetContactCheckNumber() *int32 {
	return s.ContactCheckNumber
}

func (s *QuerySecurityCheckReportResponseBodyData) GetRiskEventNumber() *int32 {
	return s.RiskEventNumber
}

func (s *QuerySecurityCheckReportResponseBodyData) GetSasCheckNumber() *int32 {
	return s.SasCheckNumber
}

func (s *QuerySecurityCheckReportResponseBodyData) GetSecurityStatus() *int32 {
	return s.SecurityStatus
}

func (s *QuerySecurityCheckReportResponseBodyData) GetSuggestionText() *string {
	return s.SuggestionText
}

func (s *QuerySecurityCheckReportResponseBodyData) SetCloudSecurityGuide(v int32) *QuerySecurityCheckReportResponseBodyData {
	s.CloudSecurityGuide = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBodyData) SetConfigCheckNumber(v int32) *QuerySecurityCheckReportResponseBodyData {
	s.ConfigCheckNumber = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBodyData) SetContactCheckNumber(v int32) *QuerySecurityCheckReportResponseBodyData {
	s.ContactCheckNumber = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBodyData) SetRiskEventNumber(v int32) *QuerySecurityCheckReportResponseBodyData {
	s.RiskEventNumber = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBodyData) SetSasCheckNumber(v int32) *QuerySecurityCheckReportResponseBodyData {
	s.SasCheckNumber = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBodyData) SetSecurityStatus(v int32) *QuerySecurityCheckReportResponseBodyData {
	s.SecurityStatus = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBodyData) SetSuggestionText(v string) *QuerySecurityCheckReportResponseBodyData {
	s.SuggestionText = &v
	return s
}

func (s *QuerySecurityCheckReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
