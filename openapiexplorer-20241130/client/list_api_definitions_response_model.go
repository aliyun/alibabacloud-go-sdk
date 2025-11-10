// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiDefinitionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiDefinitionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiDefinitionsResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *ListApiDefinitionsResponse
	GetBody() map[string]interface{}
}

type ListApiDefinitionsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiDefinitionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListApiDefinitionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiDefinitionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiDefinitionsResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *ListApiDefinitionsResponse) SetHeaders(v map[string]*string) *ListApiDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListApiDefinitionsResponse) SetStatusCode(v int32) *ListApiDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiDefinitionsResponse) SetBody(v map[string]interface{}) *ListApiDefinitionsResponse {
	s.Body = v
	return s
}

func (s *ListApiDefinitionsResponse) Validate() error {
	return dara.Validate(s)
}
