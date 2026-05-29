// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceWhiteListResponseBody) *ListInstanceWhiteListResponse
	GetBody() *ListInstanceWhiteListResponseBody
}

type ListInstanceWhiteListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceWhiteListResponse) GetBody() *ListInstanceWhiteListResponseBody {
	return s.Body
}

func (s *ListInstanceWhiteListResponse) SetHeaders(v map[string]*string) *ListInstanceWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceWhiteListResponse) SetStatusCode(v int32) *ListInstanceWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceWhiteListResponse) SetBody(v *ListInstanceWhiteListResponseBody) *ListInstanceWhiteListResponse {
	s.Body = v
	return s
}

func (s *ListInstanceWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
