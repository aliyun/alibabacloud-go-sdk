// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicalModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogicalModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogicalModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListLogicalModelsResponseBody) *ListLogicalModelsResponse
	GetBody() *ListLogicalModelsResponseBody
}

type ListLogicalModelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogicalModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogicalModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsResponse) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogicalModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogicalModelsResponse) GetBody() *ListLogicalModelsResponseBody {
	return s.Body
}

func (s *ListLogicalModelsResponse) SetHeaders(v map[string]*string) *ListLogicalModelsResponse {
	s.Headers = v
	return s
}

func (s *ListLogicalModelsResponse) SetStatusCode(v int32) *ListLogicalModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogicalModelsResponse) SetBody(v *ListLogicalModelsResponseBody) *ListLogicalModelsResponse {
	s.Body = v
	return s
}

func (s *ListLogicalModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
