// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKubernetesSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKubernetesSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKubernetesSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetKubernetesSourceResponseBody) *GetKubernetesSourceResponse
	GetBody() *GetKubernetesSourceResponseBody
}

type GetKubernetesSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKubernetesSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKubernetesSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKubernetesSourceResponse) GoString() string {
	return s.String()
}

func (s *GetKubernetesSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKubernetesSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKubernetesSourceResponse) GetBody() *GetKubernetesSourceResponseBody {
	return s.Body
}

func (s *GetKubernetesSourceResponse) SetHeaders(v map[string]*string) *GetKubernetesSourceResponse {
	s.Headers = v
	return s
}

func (s *GetKubernetesSourceResponse) SetStatusCode(v int32) *GetKubernetesSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKubernetesSourceResponse) SetBody(v *GetKubernetesSourceResponseBody) *GetKubernetesSourceResponse {
	s.Body = v
	return s
}

func (s *GetKubernetesSourceResponse) Validate() error {
	return dara.Validate(s)
}
