// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterInterceptionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterInterceptionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterInterceptionConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterInterceptionConfigResponseBody) *ListClusterInterceptionConfigResponse
	GetBody() *ListClusterInterceptionConfigResponseBody
}

type ListClusterInterceptionConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterInterceptionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterInterceptionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterInterceptionConfigResponse) GoString() string {
	return s.String()
}

func (s *ListClusterInterceptionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterInterceptionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterInterceptionConfigResponse) GetBody() *ListClusterInterceptionConfigResponseBody {
	return s.Body
}

func (s *ListClusterInterceptionConfigResponse) SetHeaders(v map[string]*string) *ListClusterInterceptionConfigResponse {
	s.Headers = v
	return s
}

func (s *ListClusterInterceptionConfigResponse) SetStatusCode(v int32) *ListClusterInterceptionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterInterceptionConfigResponse) SetBody(v *ListClusterInterceptionConfigResponseBody) *ListClusterInterceptionConfigResponse {
	s.Body = v
	return s
}

func (s *ListClusterInterceptionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
