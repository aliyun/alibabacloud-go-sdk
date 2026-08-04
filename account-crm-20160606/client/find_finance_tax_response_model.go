// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindFinanceTaxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindFinanceTaxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindFinanceTaxResponse
	GetStatusCode() *int32
	SetBody(v *FindFinanceTaxResponseBody) *FindFinanceTaxResponse
	GetBody() *FindFinanceTaxResponseBody
}

type FindFinanceTaxResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindFinanceTaxResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindFinanceTaxResponse) String() string {
	return dara.Prettify(s)
}

func (s FindFinanceTaxResponse) GoString() string {
	return s.String()
}

func (s *FindFinanceTaxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindFinanceTaxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindFinanceTaxResponse) GetBody() *FindFinanceTaxResponseBody {
	return s.Body
}

func (s *FindFinanceTaxResponse) SetHeaders(v map[string]*string) *FindFinanceTaxResponse {
	s.Headers = v
	return s
}

func (s *FindFinanceTaxResponse) SetStatusCode(v int32) *FindFinanceTaxResponse {
	s.StatusCode = &v
	return s
}

func (s *FindFinanceTaxResponse) SetBody(v *FindFinanceTaxResponseBody) *FindFinanceTaxResponse {
	s.Body = v
	return s
}

func (s *FindFinanceTaxResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
