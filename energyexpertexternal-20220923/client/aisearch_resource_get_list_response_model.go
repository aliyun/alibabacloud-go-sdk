// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceGetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AISearchResourceGetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AISearchResourceGetListResponse
	GetStatusCode() *int32
	SetBody(v *AISearchResourceGetListResponseBody) *AISearchResourceGetListResponse
	GetBody() *AISearchResourceGetListResponseBody
}

type AISearchResourceGetListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AISearchResourceGetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AISearchResourceGetListResponse) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceGetListResponse) GoString() string {
	return s.String()
}

func (s *AISearchResourceGetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AISearchResourceGetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AISearchResourceGetListResponse) GetBody() *AISearchResourceGetListResponseBody {
	return s.Body
}

func (s *AISearchResourceGetListResponse) SetHeaders(v map[string]*string) *AISearchResourceGetListResponse {
	s.Headers = v
	return s
}

func (s *AISearchResourceGetListResponse) SetStatusCode(v int32) *AISearchResourceGetListResponse {
	s.StatusCode = &v
	return s
}

func (s *AISearchResourceGetListResponse) SetBody(v *AISearchResourceGetListResponseBody) *AISearchResourceGetListResponse {
	s.Body = v
	return s
}

func (s *AISearchResourceGetListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
