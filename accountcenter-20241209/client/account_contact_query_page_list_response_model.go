// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactQueryPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccountContactQueryPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccountContactQueryPageListResponse
	GetStatusCode() *int32
	SetBody(v *AccountContactQueryPageListResponseBody) *AccountContactQueryPageListResponse
	GetBody() *AccountContactQueryPageListResponseBody
}

type AccountContactQueryPageListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountContactQueryPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccountContactQueryPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s AccountContactQueryPageListResponse) GoString() string {
	return s.String()
}

func (s *AccountContactQueryPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccountContactQueryPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccountContactQueryPageListResponse) GetBody() *AccountContactQueryPageListResponseBody {
	return s.Body
}

func (s *AccountContactQueryPageListResponse) SetHeaders(v map[string]*string) *AccountContactQueryPageListResponse {
	s.Headers = v
	return s
}

func (s *AccountContactQueryPageListResponse) SetStatusCode(v int32) *AccountContactQueryPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *AccountContactQueryPageListResponse) SetBody(v *AccountContactQueryPageListResponseBody) *AccountContactQueryPageListResponse {
	s.Body = v
	return s
}

func (s *AccountContactQueryPageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
