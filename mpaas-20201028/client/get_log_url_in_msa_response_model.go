// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogUrlInMsaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogUrlInMsaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogUrlInMsaResponse
	GetStatusCode() *int32
	SetBody(v *GetLogUrlInMsaResponseBody) *GetLogUrlInMsaResponse
	GetBody() *GetLogUrlInMsaResponseBody
}

type GetLogUrlInMsaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogUrlInMsaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogUrlInMsaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogUrlInMsaResponse) GoString() string {
	return s.String()
}

func (s *GetLogUrlInMsaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogUrlInMsaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogUrlInMsaResponse) GetBody() *GetLogUrlInMsaResponseBody {
	return s.Body
}

func (s *GetLogUrlInMsaResponse) SetHeaders(v map[string]*string) *GetLogUrlInMsaResponse {
	s.Headers = v
	return s
}

func (s *GetLogUrlInMsaResponse) SetStatusCode(v int32) *GetLogUrlInMsaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogUrlInMsaResponse) SetBody(v *GetLogUrlInMsaResponseBody) *GetLogUrlInMsaResponse {
	s.Body = v
	return s
}

func (s *GetLogUrlInMsaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
