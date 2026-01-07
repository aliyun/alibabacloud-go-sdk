// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactQueryDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccountContactQueryDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccountContactQueryDetailResponse
	GetStatusCode() *int32
	SetBody(v *AccountContactQueryDetailResponseBody) *AccountContactQueryDetailResponse
	GetBody() *AccountContactQueryDetailResponseBody
}

type AccountContactQueryDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountContactQueryDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccountContactQueryDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s AccountContactQueryDetailResponse) GoString() string {
	return s.String()
}

func (s *AccountContactQueryDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccountContactQueryDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccountContactQueryDetailResponse) GetBody() *AccountContactQueryDetailResponseBody {
	return s.Body
}

func (s *AccountContactQueryDetailResponse) SetHeaders(v map[string]*string) *AccountContactQueryDetailResponse {
	s.Headers = v
	return s
}

func (s *AccountContactQueryDetailResponse) SetStatusCode(v int32) *AccountContactQueryDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *AccountContactQueryDetailResponse) SetBody(v *AccountContactQueryDetailResponseBody) *AccountContactQueryDetailResponse {
	s.Body = v
	return s
}

func (s *AccountContactQueryDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
