// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsrModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAsrModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAsrModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListAsrModelsResponseBody) *ListAsrModelsResponse
	GetBody() *ListAsrModelsResponseBody
}

type ListAsrModelsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsrModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsrModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAsrModelsResponse) GoString() string {
	return s.String()
}

func (s *ListAsrModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAsrModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAsrModelsResponse) GetBody() *ListAsrModelsResponseBody {
	return s.Body
}

func (s *ListAsrModelsResponse) SetHeaders(v map[string]*string) *ListAsrModelsResponse {
	s.Headers = v
	return s
}

func (s *ListAsrModelsResponse) SetStatusCode(v int32) *ListAsrModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsrModelsResponse) SetBody(v *ListAsrModelsResponseBody) *ListAsrModelsResponse {
	s.Body = v
	return s
}

func (s *ListAsrModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
