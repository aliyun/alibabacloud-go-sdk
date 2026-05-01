// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigResourceGroupModelTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigResourceGroupModelTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigResourceGroupModelTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ConfigResourceGroupModelTemplateResponseBody) *ConfigResourceGroupModelTemplateResponse
	GetBody() *ConfigResourceGroupModelTemplateResponseBody
}

type ConfigResourceGroupModelTemplateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigResourceGroupModelTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigResourceGroupModelTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigResourceGroupModelTemplateResponse) GoString() string {
	return s.String()
}

func (s *ConfigResourceGroupModelTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigResourceGroupModelTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigResourceGroupModelTemplateResponse) GetBody() *ConfigResourceGroupModelTemplateResponseBody {
	return s.Body
}

func (s *ConfigResourceGroupModelTemplateResponse) SetHeaders(v map[string]*string) *ConfigResourceGroupModelTemplateResponse {
	s.Headers = v
	return s
}

func (s *ConfigResourceGroupModelTemplateResponse) SetStatusCode(v int32) *ConfigResourceGroupModelTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigResourceGroupModelTemplateResponse) SetBody(v *ConfigResourceGroupModelTemplateResponseBody) *ConfigResourceGroupModelTemplateResponse {
	s.Body = v
	return s
}

func (s *ConfigResourceGroupModelTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
