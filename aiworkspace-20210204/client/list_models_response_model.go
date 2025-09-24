// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListModelsResponseBody) *ListModelsResponse
	GetBody() *ListModelsResponseBody
}

type ListModelsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponse) GoString() string {
	return s.String()
}

func (s *ListModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelsResponse) GetBody() *ListModelsResponseBody {
	return s.Body
}

func (s *ListModelsResponse) SetHeaders(v map[string]*string) *ListModelsResponse {
	s.Headers = v
	return s
}

func (s *ListModelsResponse) SetStatusCode(v int32) *ListModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelsResponse) SetBody(v *ListModelsResponseBody) *ListModelsResponse {
	s.Body = v
	return s
}

func (s *ListModelsResponse) Validate() error {
	return dara.Validate(s)
}
