// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportEdsAgentInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReportEdsAgentInfoResponseBody
	GetRequestId() *string
}

type ReportEdsAgentInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportEdsAgentInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportEdsAgentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportEdsAgentInfoResponseBody) SetRequestId(v string) *ReportEdsAgentInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportEdsAgentInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
