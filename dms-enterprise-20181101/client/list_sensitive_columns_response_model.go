// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSensitiveColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSensitiveColumnsResponse
	GetStatusCode() *int32
	SetBody(v *ListSensitiveColumnsResponseBody) *ListSensitiveColumnsResponse
	GetBody() *ListSensitiveColumnsResponseBody
}

type ListSensitiveColumnsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSensitiveColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSensitiveColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSensitiveColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSensitiveColumnsResponse) GetBody() *ListSensitiveColumnsResponseBody {
	return s.Body
}

func (s *ListSensitiveColumnsResponse) SetHeaders(v map[string]*string) *ListSensitiveColumnsResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveColumnsResponse) SetStatusCode(v int32) *ListSensitiveColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSensitiveColumnsResponse) SetBody(v *ListSensitiveColumnsResponseBody) *ListSensitiveColumnsResponse {
	s.Body = v
	return s
}

func (s *ListSensitiveColumnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
