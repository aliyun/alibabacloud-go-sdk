// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNluModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNluModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNluModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListNluModelsResponseBody) *ListNluModelsResponse
	GetBody() *ListNluModelsResponseBody
}

type ListNluModelsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNluModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNluModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNluModelsResponse) GoString() string {
	return s.String()
}

func (s *ListNluModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNluModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNluModelsResponse) GetBody() *ListNluModelsResponseBody {
	return s.Body
}

func (s *ListNluModelsResponse) SetHeaders(v map[string]*string) *ListNluModelsResponse {
	s.Headers = v
	return s
}

func (s *ListNluModelsResponse) SetStatusCode(v int32) *ListNluModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNluModelsResponse) SetBody(v *ListNluModelsResponseBody) *ListNluModelsResponse {
	s.Body = v
	return s
}

func (s *ListNluModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
