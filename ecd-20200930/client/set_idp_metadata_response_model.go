// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdpMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetIdpMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetIdpMetadataResponse
	GetStatusCode() *int32
	SetBody(v *SetIdpMetadataResponseBody) *SetIdpMetadataResponse
	GetBody() *SetIdpMetadataResponseBody
}

type SetIdpMetadataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetIdpMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetIdpMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s SetIdpMetadataResponse) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetIdpMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetIdpMetadataResponse) GetBody() *SetIdpMetadataResponseBody {
	return s.Body
}

func (s *SetIdpMetadataResponse) SetHeaders(v map[string]*string) *SetIdpMetadataResponse {
	s.Headers = v
	return s
}

func (s *SetIdpMetadataResponse) SetStatusCode(v int32) *SetIdpMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *SetIdpMetadataResponse) SetBody(v *SetIdpMetadataResponseBody) *SetIdpMetadataResponse {
	s.Body = v
	return s
}

func (s *SetIdpMetadataResponse) Validate() error {
	return dara.Validate(s)
}
