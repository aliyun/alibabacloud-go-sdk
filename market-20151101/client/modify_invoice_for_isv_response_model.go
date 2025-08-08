// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInvoiceForIsvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInvoiceForIsvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInvoiceForIsvResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInvoiceForIsvResponseBody) *ModifyInvoiceForIsvResponse
	GetBody() *ModifyInvoiceForIsvResponseBody
}

type ModifyInvoiceForIsvResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInvoiceForIsvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInvoiceForIsvResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInvoiceForIsvResponse) GoString() string {
	return s.String()
}

func (s *ModifyInvoiceForIsvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInvoiceForIsvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInvoiceForIsvResponse) GetBody() *ModifyInvoiceForIsvResponseBody {
	return s.Body
}

func (s *ModifyInvoiceForIsvResponse) SetHeaders(v map[string]*string) *ModifyInvoiceForIsvResponse {
	s.Headers = v
	return s
}

func (s *ModifyInvoiceForIsvResponse) SetStatusCode(v int32) *ModifyInvoiceForIsvResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInvoiceForIsvResponse) SetBody(v *ModifyInvoiceForIsvResponseBody) *ModifyInvoiceForIsvResponse {
	s.Body = v
	return s
}

func (s *ModifyInvoiceForIsvResponse) Validate() error {
	return dara.Validate(s)
}
