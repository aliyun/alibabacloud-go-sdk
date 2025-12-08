// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealPersonVerificationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRealPersonVerificationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRealPersonVerificationResultResponse
	GetStatusCode() *int32
	SetBody(v *GetRealPersonVerificationResultResponseBody) *GetRealPersonVerificationResultResponse
	GetBody() *GetRealPersonVerificationResultResponseBody
}

type GetRealPersonVerificationResultResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealPersonVerificationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealPersonVerificationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRealPersonVerificationResultResponse) GoString() string {
	return s.String()
}

func (s *GetRealPersonVerificationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRealPersonVerificationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRealPersonVerificationResultResponse) GetBody() *GetRealPersonVerificationResultResponseBody {
	return s.Body
}

func (s *GetRealPersonVerificationResultResponse) SetHeaders(v map[string]*string) *GetRealPersonVerificationResultResponse {
	s.Headers = v
	return s
}

func (s *GetRealPersonVerificationResultResponse) SetStatusCode(v int32) *GetRealPersonVerificationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealPersonVerificationResultResponse) SetBody(v *GetRealPersonVerificationResultResponseBody) *GetRealPersonVerificationResultResponse {
	s.Body = v
	return s
}

func (s *GetRealPersonVerificationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
