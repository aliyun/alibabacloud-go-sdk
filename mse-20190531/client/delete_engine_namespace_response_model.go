// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEngineNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEngineNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEngineNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEngineNamespaceResponseBody) *DeleteEngineNamespaceResponse
	GetBody() *DeleteEngineNamespaceResponseBody
}

type DeleteEngineNamespaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEngineNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEngineNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEngineNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEngineNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEngineNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEngineNamespaceResponse) GetBody() *DeleteEngineNamespaceResponseBody {
	return s.Body
}

func (s *DeleteEngineNamespaceResponse) SetHeaders(v map[string]*string) *DeleteEngineNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEngineNamespaceResponse) SetStatusCode(v int32) *DeleteEngineNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEngineNamespaceResponse) SetBody(v *DeleteEngineNamespaceResponseBody) *DeleteEngineNamespaceResponse {
	s.Body = v
	return s
}

func (s *DeleteEngineNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
