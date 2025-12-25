// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNamespaceResponseBody) *DeleteNamespaceResponse
	GetBody() *DeleteNamespaceResponseBody
}

type DeleteNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNamespaceResponse) GetBody() *DeleteNamespaceResponseBody {
	return s.Body
}

func (s *DeleteNamespaceResponse) SetHeaders(v map[string]*string) *DeleteNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNamespaceResponse) SetStatusCode(v int32) *DeleteNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNamespaceResponse) SetBody(v *DeleteNamespaceResponseBody) *DeleteNamespaceResponse {
	s.Body = v
	return s
}

func (s *DeleteNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
