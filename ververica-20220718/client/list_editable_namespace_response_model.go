// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEditableNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEditableNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEditableNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *ListEditableNamespaceResponseBody) *ListEditableNamespaceResponse
	GetBody() *ListEditableNamespaceResponseBody
}

type ListEditableNamespaceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEditableNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEditableNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEditableNamespaceResponse) GoString() string {
	return s.String()
}

func (s *ListEditableNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEditableNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEditableNamespaceResponse) GetBody() *ListEditableNamespaceResponseBody {
	return s.Body
}

func (s *ListEditableNamespaceResponse) SetHeaders(v map[string]*string) *ListEditableNamespaceResponse {
	s.Headers = v
	return s
}

func (s *ListEditableNamespaceResponse) SetStatusCode(v int32) *ListEditableNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEditableNamespaceResponse) SetBody(v *ListEditableNamespaceResponseBody) *ListEditableNamespaceResponse {
	s.Body = v
	return s
}

func (s *ListEditableNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
