// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSensitiveColumnInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSensitiveColumnInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListSensitiveColumnInfoResponseBody) *ListSensitiveColumnInfoResponse
	GetBody() *ListSensitiveColumnInfoResponseBody
}

type ListSensitiveColumnInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSensitiveColumnInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSensitiveColumnInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSensitiveColumnInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSensitiveColumnInfoResponse) GetBody() *ListSensitiveColumnInfoResponseBody {
	return s.Body
}

func (s *ListSensitiveColumnInfoResponse) SetHeaders(v map[string]*string) *ListSensitiveColumnInfoResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveColumnInfoResponse) SetStatusCode(v int32) *ListSensitiveColumnInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSensitiveColumnInfoResponse) SetBody(v *ListSensitiveColumnInfoResponseBody) *ListSensitiveColumnInfoResponse {
	s.Body = v
	return s
}

func (s *ListSensitiveColumnInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
