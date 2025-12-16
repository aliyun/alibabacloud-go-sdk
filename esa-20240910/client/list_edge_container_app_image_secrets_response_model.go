// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppImageSecretsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeContainerAppImageSecretsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeContainerAppImageSecretsResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeContainerAppImageSecretsResponseBody) *ListEdgeContainerAppImageSecretsResponse
	GetBody() *ListEdgeContainerAppImageSecretsResponseBody
}

type ListEdgeContainerAppImageSecretsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeContainerAppImageSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeContainerAppImageSecretsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppImageSecretsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppImageSecretsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeContainerAppImageSecretsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeContainerAppImageSecretsResponse) GetBody() *ListEdgeContainerAppImageSecretsResponseBody {
	return s.Body
}

func (s *ListEdgeContainerAppImageSecretsResponse) SetHeaders(v map[string]*string) *ListEdgeContainerAppImageSecretsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeContainerAppImageSecretsResponse) SetStatusCode(v int32) *ListEdgeContainerAppImageSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeContainerAppImageSecretsResponse) SetBody(v *ListEdgeContainerAppImageSecretsResponseBody) *ListEdgeContainerAppImageSecretsResponse {
	s.Body = v
	return s
}

func (s *ListEdgeContainerAppImageSecretsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
