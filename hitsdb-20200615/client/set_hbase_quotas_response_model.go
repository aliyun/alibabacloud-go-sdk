// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHBaseQuotasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetHBaseQuotasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetHBaseQuotasResponse
	GetStatusCode() *int32
	SetBody(v *SetHBaseQuotasResponseBody) *SetHBaseQuotasResponse
	GetBody() *SetHBaseQuotasResponseBody
}

type SetHBaseQuotasResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetHBaseQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetHBaseQuotasResponse) String() string {
	return dara.Prettify(s)
}

func (s SetHBaseQuotasResponse) GoString() string {
	return s.String()
}

func (s *SetHBaseQuotasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetHBaseQuotasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetHBaseQuotasResponse) GetBody() *SetHBaseQuotasResponseBody {
	return s.Body
}

func (s *SetHBaseQuotasResponse) SetHeaders(v map[string]*string) *SetHBaseQuotasResponse {
	s.Headers = v
	return s
}

func (s *SetHBaseQuotasResponse) SetStatusCode(v int32) *SetHBaseQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *SetHBaseQuotasResponse) SetBody(v *SetHBaseQuotasResponseBody) *SetHBaseQuotasResponse {
	s.Body = v
	return s
}

func (s *SetHBaseQuotasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
