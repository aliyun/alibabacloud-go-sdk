// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEngineNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEngineNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *ListEngineNamespacesResponseBody) *ListEngineNamespacesResponse
	GetBody() *ListEngineNamespacesResponseBody
}

type ListEngineNamespacesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEngineNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEngineNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEngineNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListEngineNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEngineNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEngineNamespacesResponse) GetBody() *ListEngineNamespacesResponseBody {
	return s.Body
}

func (s *ListEngineNamespacesResponse) SetHeaders(v map[string]*string) *ListEngineNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListEngineNamespacesResponse) SetStatusCode(v int32) *ListEngineNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEngineNamespacesResponse) SetBody(v *ListEngineNamespacesResponseBody) *ListEngineNamespacesResponse {
	s.Body = v
	return s
}

func (s *ListEngineNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
