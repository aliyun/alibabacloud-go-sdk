// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKvNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKvNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *GetKvNamespaceResponseBody) *GetKvNamespaceResponse
	GetBody() *GetKvNamespaceResponseBody
}

type GetKvNamespaceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKvNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetKvNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKvNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKvNamespaceResponse) GetBody() *GetKvNamespaceResponseBody {
	return s.Body
}

func (s *GetKvNamespaceResponse) SetHeaders(v map[string]*string) *GetKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetKvNamespaceResponse) SetStatusCode(v int32) *GetKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKvNamespaceResponse) SetBody(v *GetKvNamespaceResponseBody) *GetKvNamespaceResponse {
	s.Body = v
	return s
}

func (s *GetKvNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
