// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupportModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupportModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListSupportModelsResponseBody) *ListSupportModelsResponse
	GetBody() *ListSupportModelsResponseBody
}

type ListSupportModelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupportModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupportModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupportModelsResponse) GoString() string {
	return s.String()
}

func (s *ListSupportModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupportModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupportModelsResponse) GetBody() *ListSupportModelsResponseBody {
	return s.Body
}

func (s *ListSupportModelsResponse) SetHeaders(v map[string]*string) *ListSupportModelsResponse {
	s.Headers = v
	return s
}

func (s *ListSupportModelsResponse) SetStatusCode(v int32) *ListSupportModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportModelsResponse) SetBody(v *ListSupportModelsResponseBody) *ListSupportModelsResponse {
	s.Body = v
	return s
}

func (s *ListSupportModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
