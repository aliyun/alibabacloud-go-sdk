// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutDcdnKvNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutDcdnKvNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *PutDcdnKvNamespaceResponseBody) *PutDcdnKvNamespaceResponse
	GetBody() *PutDcdnKvNamespaceResponseBody
}

type PutDcdnKvNamespaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDcdnKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDcdnKvNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *PutDcdnKvNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutDcdnKvNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutDcdnKvNamespaceResponse) GetBody() *PutDcdnKvNamespaceResponseBody {
	return s.Body
}

func (s *PutDcdnKvNamespaceResponse) SetHeaders(v map[string]*string) *PutDcdnKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *PutDcdnKvNamespaceResponse) SetStatusCode(v int32) *PutDcdnKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDcdnKvNamespaceResponse) SetBody(v *PutDcdnKvNamespaceResponseBody) *PutDcdnKvNamespaceResponse {
	s.Body = v
	return s
}

func (s *PutDcdnKvNamespaceResponse) Validate() error {
	return dara.Validate(s)
}
