// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhitelistTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWhitelistTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWhitelistTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWhitelistTemplateResponseBody) *ModifyWhitelistTemplateResponse
	GetBody() *ModifyWhitelistTemplateResponseBody
}

type ModifyWhitelistTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWhitelistTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWhitelistTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhitelistTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyWhitelistTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWhitelistTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWhitelistTemplateResponse) GetBody() *ModifyWhitelistTemplateResponseBody {
	return s.Body
}

func (s *ModifyWhitelistTemplateResponse) SetHeaders(v map[string]*string) *ModifyWhitelistTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyWhitelistTemplateResponse) SetStatusCode(v int32) *ModifyWhitelistTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWhitelistTemplateResponse) SetBody(v *ModifyWhitelistTemplateResponseBody) *ModifyWhitelistTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyWhitelistTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
