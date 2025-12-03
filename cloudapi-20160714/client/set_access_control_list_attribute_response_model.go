// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessControlListAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAccessControlListAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAccessControlListAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetAccessControlListAttributeResponseBody) *SetAccessControlListAttributeResponse
	GetBody() *SetAccessControlListAttributeResponseBody
}

type SetAccessControlListAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccessControlListAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccessControlListAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAccessControlListAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetAccessControlListAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAccessControlListAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAccessControlListAttributeResponse) GetBody() *SetAccessControlListAttributeResponseBody {
	return s.Body
}

func (s *SetAccessControlListAttributeResponse) SetHeaders(v map[string]*string) *SetAccessControlListAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetAccessControlListAttributeResponse) SetStatusCode(v int32) *SetAccessControlListAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessControlListAttributeResponse) SetBody(v *SetAccessControlListAttributeResponseBody) *SetAccessControlListAttributeResponse {
	s.Body = v
	return s
}

func (s *SetAccessControlListAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
