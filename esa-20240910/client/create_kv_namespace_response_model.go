// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKvNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKvNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKvNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateKvNamespaceResponseBody) *CreateKvNamespaceResponse
	GetBody() *CreateKvNamespaceResponseBody
}

type CreateKvNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKvNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKvNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKvNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateKvNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKvNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKvNamespaceResponse) GetBody() *CreateKvNamespaceResponseBody {
	return s.Body
}

func (s *CreateKvNamespaceResponse) SetHeaders(v map[string]*string) *CreateKvNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateKvNamespaceResponse) SetStatusCode(v int32) *CreateKvNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKvNamespaceResponse) SetBody(v *CreateKvNamespaceResponseBody) *CreateKvNamespaceResponse {
	s.Body = v
	return s
}

func (s *CreateKvNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
