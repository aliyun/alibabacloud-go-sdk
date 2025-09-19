// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRegistryExtraResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageRegistryExtraResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageRegistryExtraResponse
	GetStatusCode() *int32
	SetBody(v *ListImageRegistryExtraResponseBody) *ListImageRegistryExtraResponse
	GetBody() *ListImageRegistryExtraResponseBody
}

type ListImageRegistryExtraResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageRegistryExtraResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageRegistryExtraResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageRegistryExtraResponse) GoString() string {
	return s.String()
}

func (s *ListImageRegistryExtraResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageRegistryExtraResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageRegistryExtraResponse) GetBody() *ListImageRegistryExtraResponseBody {
	return s.Body
}

func (s *ListImageRegistryExtraResponse) SetHeaders(v map[string]*string) *ListImageRegistryExtraResponse {
	s.Headers = v
	return s
}

func (s *ListImageRegistryExtraResponse) SetStatusCode(v int32) *ListImageRegistryExtraResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageRegistryExtraResponse) SetBody(v *ListImageRegistryExtraResponseBody) *ListImageRegistryExtraResponse {
	s.Body = v
	return s
}

func (s *ListImageRegistryExtraResponse) Validate() error {
	return dara.Validate(s)
}
