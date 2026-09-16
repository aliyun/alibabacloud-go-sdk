// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaNamespaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLumaNamespaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLumaNamespaceResponse
	GetStatusCode() *int32
	SetBody(v *GetLumaNamespaceResponseBody) *GetLumaNamespaceResponse
	GetBody() *GetLumaNamespaceResponseBody
}

type GetLumaNamespaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLumaNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLumaNamespaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLumaNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetLumaNamespaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLumaNamespaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLumaNamespaceResponse) GetBody() *GetLumaNamespaceResponseBody {
	return s.Body
}

func (s *GetLumaNamespaceResponse) SetHeaders(v map[string]*string) *GetLumaNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetLumaNamespaceResponse) SetStatusCode(v int32) *GetLumaNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLumaNamespaceResponse) SetBody(v *GetLumaNamespaceResponseBody) *GetLumaNamespaceResponse {
	s.Body = v
	return s
}

func (s *GetLumaNamespaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
