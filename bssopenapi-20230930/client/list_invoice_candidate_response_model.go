// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInvoiceCandidateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInvoiceCandidateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInvoiceCandidateResponse
	GetStatusCode() *int32
	SetBody(v *ListInvoiceCandidateResponseBody) *ListInvoiceCandidateResponse
	GetBody() *ListInvoiceCandidateResponseBody
}

type ListInvoiceCandidateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInvoiceCandidateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInvoiceCandidateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceCandidateResponse) GoString() string {
	return s.String()
}

func (s *ListInvoiceCandidateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInvoiceCandidateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInvoiceCandidateResponse) GetBody() *ListInvoiceCandidateResponseBody {
	return s.Body
}

func (s *ListInvoiceCandidateResponse) SetHeaders(v map[string]*string) *ListInvoiceCandidateResponse {
	s.Headers = v
	return s
}

func (s *ListInvoiceCandidateResponse) SetStatusCode(v int32) *ListInvoiceCandidateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInvoiceCandidateResponse) SetBody(v *ListInvoiceCandidateResponseBody) *ListInvoiceCandidateResponse {
	s.Body = v
	return s
}

func (s *ListInvoiceCandidateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
