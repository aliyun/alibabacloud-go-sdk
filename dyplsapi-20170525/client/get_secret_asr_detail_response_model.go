// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretAsrDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecretAsrDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecretAsrDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSecretAsrDetailResponseBody) *GetSecretAsrDetailResponse
	GetBody() *GetSecretAsrDetailResponseBody
}

type GetSecretAsrDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretAsrDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretAsrDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecretAsrDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSecretAsrDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecretAsrDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecretAsrDetailResponse) GetBody() *GetSecretAsrDetailResponseBody {
	return s.Body
}

func (s *GetSecretAsrDetailResponse) SetHeaders(v map[string]*string) *GetSecretAsrDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSecretAsrDetailResponse) SetStatusCode(v int32) *GetSecretAsrDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretAsrDetailResponse) SetBody(v *GetSecretAsrDetailResponseBody) *GetSecretAsrDetailResponse {
	s.Body = v
	return s
}

func (s *GetSecretAsrDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
