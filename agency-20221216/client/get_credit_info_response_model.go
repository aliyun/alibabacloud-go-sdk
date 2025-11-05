// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreditInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCreditInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCreditInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetCreditInfoResponseBody) *GetCreditInfoResponse
	GetBody() *GetCreditInfoResponseBody
}

type GetCreditInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreditInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCreditInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCreditInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCreditInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCreditInfoResponse) GetBody() *GetCreditInfoResponseBody {
	return s.Body
}

func (s *GetCreditInfoResponse) SetHeaders(v map[string]*string) *GetCreditInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCreditInfoResponse) SetStatusCode(v int32) *GetCreditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreditInfoResponse) SetBody(v *GetCreditInfoResponseBody) *GetCreditInfoResponse {
	s.Body = v
	return s
}

func (s *GetCreditInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
