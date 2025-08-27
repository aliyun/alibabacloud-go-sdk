// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWaitApplyInvoiceTaskDetailQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WaitApplyInvoiceTaskDetailQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WaitApplyInvoiceTaskDetailQueryResponse
	GetStatusCode() *int32
	SetBody(v *WaitApplyInvoiceTaskDetailQueryResponseBody) *WaitApplyInvoiceTaskDetailQueryResponse
	GetBody() *WaitApplyInvoiceTaskDetailQueryResponseBody
}

type WaitApplyInvoiceTaskDetailQueryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WaitApplyInvoiceTaskDetailQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WaitApplyInvoiceTaskDetailQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s WaitApplyInvoiceTaskDetailQueryResponse) GoString() string {
	return s.String()
}

func (s *WaitApplyInvoiceTaskDetailQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WaitApplyInvoiceTaskDetailQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WaitApplyInvoiceTaskDetailQueryResponse) GetBody() *WaitApplyInvoiceTaskDetailQueryResponseBody {
	return s.Body
}

func (s *WaitApplyInvoiceTaskDetailQueryResponse) SetHeaders(v map[string]*string) *WaitApplyInvoiceTaskDetailQueryResponse {
	s.Headers = v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponse) SetStatusCode(v int32) *WaitApplyInvoiceTaskDetailQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponse) SetBody(v *WaitApplyInvoiceTaskDetailQueryResponseBody) *WaitApplyInvoiceTaskDetailQueryResponse {
	s.Body = v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponse) Validate() error {
	return dara.Validate(s)
}
