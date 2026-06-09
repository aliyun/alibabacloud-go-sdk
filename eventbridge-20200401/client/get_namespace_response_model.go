// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *GetNamespaceResponseBody) *GetNamespaceResponse
	GetBody() *GetNamespaceResponseBody
}

type GetNamespaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNamespaceResponse) GetBody() *GetNamespaceResponseBody {
	return s.Body
}

func (s *GetNamespaceResponse) SetHeaders(v map[string]*string) *GetNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetNamespaceResponse) SetStatusCode(v int32) *GetNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNamespaceResponse) SetBody(v *GetNamespaceResponseBody) *GetNamespaceResponse {
	s.Body = v
	return s
}

func (s *GetNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
