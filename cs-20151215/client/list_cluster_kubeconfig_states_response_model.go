// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterKubeconfigStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterKubeconfigStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterKubeconfigStatesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterKubeconfigStatesResponseBody) *ListClusterKubeconfigStatesResponse
	GetBody() *ListClusterKubeconfigStatesResponseBody
}

type ListClusterKubeconfigStatesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterKubeconfigStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterKubeconfigStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterKubeconfigStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterKubeconfigStatesResponse) GetBody() *ListClusterKubeconfigStatesResponseBody {
	return s.Body
}

func (s *ListClusterKubeconfigStatesResponse) SetHeaders(v map[string]*string) *ListClusterKubeconfigStatesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterKubeconfigStatesResponse) SetStatusCode(v int32) *ListClusterKubeconfigStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponse) SetBody(v *ListClusterKubeconfigStatesResponseBody) *ListClusterKubeconfigStatesResponse {
	s.Body = v
	return s
}

func (s *ListClusterKubeconfigStatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
