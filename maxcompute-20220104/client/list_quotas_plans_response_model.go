// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotasPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotasPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotasPlansResponseBody) *ListQuotasPlansResponse
	GetBody() *ListQuotasPlansResponseBody
}

type ListQuotasPlansResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotasPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotasPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasPlansResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotasPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotasPlansResponse) GetBody() *ListQuotasPlansResponseBody {
	return s.Body
}

func (s *ListQuotasPlansResponse) SetHeaders(v map[string]*string) *ListQuotasPlansResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasPlansResponse) SetStatusCode(v int32) *ListQuotasPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasPlansResponse) SetBody(v *ListQuotasPlansResponseBody) *ListQuotasPlansResponse {
	s.Body = v
	return s
}

func (s *ListQuotasPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
