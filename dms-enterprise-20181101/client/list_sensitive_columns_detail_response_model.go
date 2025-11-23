// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnsDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSensitiveColumnsDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSensitiveColumnsDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListSensitiveColumnsDetailResponseBody) *ListSensitiveColumnsDetailResponse
	GetBody() *ListSensitiveColumnsDetailResponseBody
}

type ListSensitiveColumnsDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSensitiveColumnsDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSensitiveColumnsDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsDetailResponse) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSensitiveColumnsDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSensitiveColumnsDetailResponse) GetBody() *ListSensitiveColumnsDetailResponseBody {
	return s.Body
}

func (s *ListSensitiveColumnsDetailResponse) SetHeaders(v map[string]*string) *ListSensitiveColumnsDetailResponse {
	s.Headers = v
	return s
}

func (s *ListSensitiveColumnsDetailResponse) SetStatusCode(v int32) *ListSensitiveColumnsDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponse) SetBody(v *ListSensitiveColumnsDetailResponseBody) *ListSensitiveColumnsDetailResponse {
	s.Body = v
	return s
}

func (s *ListSensitiveColumnsDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
