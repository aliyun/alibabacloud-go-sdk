// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIcebergNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIcebergNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIcebergNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *Namespace) *GetIcebergNamespaceResponse
	GetBody() *Namespace
}

type GetIcebergNamespaceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Namespace         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIcebergNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIcebergNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetIcebergNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIcebergNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIcebergNamespaceResponse) GetBody() *Namespace {
	return s.Body
}

func (s *GetIcebergNamespaceResponse) SetHeaders(v map[string]*string) *GetIcebergNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetIcebergNamespaceResponse) SetStatusCode(v int32) *GetIcebergNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIcebergNamespaceResponse) SetBody(v *Namespace) *GetIcebergNamespaceResponse {
	s.Body = v
	return s
}

func (s *GetIcebergNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
