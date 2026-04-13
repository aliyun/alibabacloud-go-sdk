// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDetectConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDetectConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListDetectConfigsResponseBody) *ListDetectConfigsResponse
	GetBody() *ListDetectConfigsResponseBody
}

type ListDetectConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDetectConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDetectConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListDetectConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDetectConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDetectConfigsResponse) GetBody() *ListDetectConfigsResponseBody {
	return s.Body
}

func (s *ListDetectConfigsResponse) SetHeaders(v map[string]*string) *ListDetectConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListDetectConfigsResponse) SetStatusCode(v int32) *ListDetectConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDetectConfigsResponse) SetBody(v *ListDetectConfigsResponseBody) *ListDetectConfigsResponse {
	s.Body = v
	return s
}

func (s *ListDetectConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
