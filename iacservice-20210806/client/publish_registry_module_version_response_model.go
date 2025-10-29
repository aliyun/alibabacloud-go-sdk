// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRegistryModuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishRegistryModuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishRegistryModuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *PublishRegistryModuleVersionResponseBody) *PublishRegistryModuleVersionResponse
	GetBody() *PublishRegistryModuleVersionResponseBody
}

type PublishRegistryModuleVersionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishRegistryModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishRegistryModuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishRegistryModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishRegistryModuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishRegistryModuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishRegistryModuleVersionResponse) GetBody() *PublishRegistryModuleVersionResponseBody {
	return s.Body
}

func (s *PublishRegistryModuleVersionResponse) SetHeaders(v map[string]*string) *PublishRegistryModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishRegistryModuleVersionResponse) SetStatusCode(v int32) *PublishRegistryModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishRegistryModuleVersionResponse) SetBody(v *PublishRegistryModuleVersionResponseBody) *PublishRegistryModuleVersionResponse {
	s.Body = v
	return s
}

func (s *PublishRegistryModuleVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
