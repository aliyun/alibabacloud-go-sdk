// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAiAppByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAiAppByPageResponse
	GetStatusCode() *int32
	SetBody(v *ListAiAppByPageResponseBody) *ListAiAppByPageResponse
	GetBody() *ListAiAppByPageResponseBody
}

type ListAiAppByPageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAiAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAiAppByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppByPageResponse) GoString() string {
	return s.String()
}

func (s *ListAiAppByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAiAppByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAiAppByPageResponse) GetBody() *ListAiAppByPageResponseBody {
	return s.Body
}

func (s *ListAiAppByPageResponse) SetHeaders(v map[string]*string) *ListAiAppByPageResponse {
	s.Headers = v
	return s
}

func (s *ListAiAppByPageResponse) SetStatusCode(v int32) *ListAiAppByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAiAppByPageResponse) SetBody(v *ListAiAppByPageResponseBody) *ListAiAppByPageResponse {
	s.Body = v
	return s
}

func (s *ListAiAppByPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
