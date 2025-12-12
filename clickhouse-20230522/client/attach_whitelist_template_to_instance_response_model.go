// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachWhitelistTemplateToInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachWhitelistTemplateToInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachWhitelistTemplateToInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AttachWhitelistTemplateToInstanceResponseBody) *AttachWhitelistTemplateToInstanceResponse
	GetBody() *AttachWhitelistTemplateToInstanceResponseBody
}

type AttachWhitelistTemplateToInstanceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachWhitelistTemplateToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachWhitelistTemplateToInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachWhitelistTemplateToInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachWhitelistTemplateToInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachWhitelistTemplateToInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachWhitelistTemplateToInstanceResponse) GetBody() *AttachWhitelistTemplateToInstanceResponseBody {
	return s.Body
}

func (s *AttachWhitelistTemplateToInstanceResponse) SetHeaders(v map[string]*string) *AttachWhitelistTemplateToInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponse) SetStatusCode(v int32) *AttachWhitelistTemplateToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponse) SetBody(v *AttachWhitelistTemplateToInstanceResponseBody) *AttachWhitelistTemplateToInstanceResponse {
	s.Body = v
	return s
}

func (s *AttachWhitelistTemplateToInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
