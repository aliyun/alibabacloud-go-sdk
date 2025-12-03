// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainExtensionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDomainExtensionResponseBody) *DeleteDomainExtensionResponse
	GetBody() *DeleteDomainExtensionResponseBody
}

type DeleteDomainExtensionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainExtensionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainExtensionResponse) GetBody() *DeleteDomainExtensionResponseBody {
	return s.Body
}

func (s *DeleteDomainExtensionResponse) SetHeaders(v map[string]*string) *DeleteDomainExtensionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainExtensionResponse) SetStatusCode(v int32) *DeleteDomainExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainExtensionResponse) SetBody(v *DeleteDomainExtensionResponseBody) *DeleteDomainExtensionResponse {
	s.Body = v
	return s
}

func (s *DeleteDomainExtensionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
