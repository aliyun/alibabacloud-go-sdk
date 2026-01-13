// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEngineConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEngineConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListEngineConfigsResponseBody) *ListEngineConfigsResponse
	GetBody() *ListEngineConfigsResponseBody
}

type ListEngineConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEngineConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEngineConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEngineConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListEngineConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEngineConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEngineConfigsResponse) GetBody() *ListEngineConfigsResponseBody {
	return s.Body
}

func (s *ListEngineConfigsResponse) SetHeaders(v map[string]*string) *ListEngineConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListEngineConfigsResponse) SetStatusCode(v int32) *ListEngineConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEngineConfigsResponse) SetBody(v *ListEngineConfigsResponseBody) *ListEngineConfigsResponse {
	s.Body = v
	return s
}

func (s *ListEngineConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
