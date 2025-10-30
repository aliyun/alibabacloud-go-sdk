// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretNoDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySecretNoDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySecretNoDetailResponse
	GetStatusCode() *int32
	SetBody(v *QuerySecretNoDetailResponseBody) *QuerySecretNoDetailResponse
	GetBody() *QuerySecretNoDetailResponseBody
}

type QuerySecretNoDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecretNoDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecretNoDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoDetailResponse) GoString() string {
	return s.String()
}

func (s *QuerySecretNoDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySecretNoDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySecretNoDetailResponse) GetBody() *QuerySecretNoDetailResponseBody {
	return s.Body
}

func (s *QuerySecretNoDetailResponse) SetHeaders(v map[string]*string) *QuerySecretNoDetailResponse {
	s.Headers = v
	return s
}

func (s *QuerySecretNoDetailResponse) SetStatusCode(v int32) *QuerySecretNoDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecretNoDetailResponse) SetBody(v *QuerySecretNoDetailResponseBody) *QuerySecretNoDetailResponse {
	s.Body = v
	return s
}

func (s *QuerySecretNoDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
