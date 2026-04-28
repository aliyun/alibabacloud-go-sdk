// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportBizAlertInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReportBizAlertInfoResponseBody
	GetCode() *string
	SetData(v *ReportBizAlertInfoResponseBodyData) *ReportBizAlertInfoResponseBody
	GetData() *ReportBizAlertInfoResponseBodyData
	SetMessage(v string) *ReportBizAlertInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReportBizAlertInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReportBizAlertInfoResponseBody
	GetSuccess() *bool
}

type ReportBizAlertInfoResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReportBizAlertInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportBizAlertInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportBizAlertInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReportBizAlertInfoResponseBody) GetData() *ReportBizAlertInfoResponseBodyData {
	return s.Data
}

func (s *ReportBizAlertInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReportBizAlertInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportBizAlertInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReportBizAlertInfoResponseBody) SetCode(v string) *ReportBizAlertInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ReportBizAlertInfoResponseBody) SetData(v *ReportBizAlertInfoResponseBodyData) *ReportBizAlertInfoResponseBody {
	s.Data = v
	return s
}

func (s *ReportBizAlertInfoResponseBody) SetMessage(v string) *ReportBizAlertInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ReportBizAlertInfoResponseBody) SetRequestId(v string) *ReportBizAlertInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportBizAlertInfoResponseBody) SetSuccess(v bool) *ReportBizAlertInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ReportBizAlertInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReportBizAlertInfoResponseBodyData struct {
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ReportBizAlertInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReportBizAlertInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReportBizAlertInfoResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *ReportBizAlertInfoResponseBodyData) SetResult(v string) *ReportBizAlertInfoResponseBodyData {
	s.Result = &v
	return s
}

func (s *ReportBizAlertInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
