// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIImageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIImageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIImageInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListAIImageInfoResponseBody) *ListAIImageInfoResponse
	GetBody() *ListAIImageInfoResponseBody
}

type ListAIImageInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIImageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIImageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIImageInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAIImageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIImageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIImageInfoResponse) GetBody() *ListAIImageInfoResponseBody {
	return s.Body
}

func (s *ListAIImageInfoResponse) SetHeaders(v map[string]*string) *ListAIImageInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAIImageInfoResponse) SetStatusCode(v int32) *ListAIImageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIImageInfoResponse) SetBody(v *ListAIImageInfoResponseBody) *ListAIImageInfoResponse {
	s.Body = v
	return s
}

func (s *ListAIImageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
