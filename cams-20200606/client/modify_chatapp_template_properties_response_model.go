// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChatappTemplatePropertiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyChatappTemplatePropertiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyChatappTemplatePropertiesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyChatappTemplatePropertiesResponseBody) *ModifyChatappTemplatePropertiesResponse
	GetBody() *ModifyChatappTemplatePropertiesResponseBody
}

type ModifyChatappTemplatePropertiesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyChatappTemplatePropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyChatappTemplatePropertiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplatePropertiesResponse) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplatePropertiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyChatappTemplatePropertiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyChatappTemplatePropertiesResponse) GetBody() *ModifyChatappTemplatePropertiesResponseBody {
	return s.Body
}

func (s *ModifyChatappTemplatePropertiesResponse) SetHeaders(v map[string]*string) *ModifyChatappTemplatePropertiesResponse {
	s.Headers = v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponse) SetStatusCode(v int32) *ModifyChatappTemplatePropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponse) SetBody(v *ModifyChatappTemplatePropertiesResponseBody) *ModifyChatappTemplatePropertiesResponse {
	s.Body = v
	return s
}

func (s *ModifyChatappTemplatePropertiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
