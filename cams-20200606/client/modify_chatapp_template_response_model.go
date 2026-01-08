// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChatappTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyChatappTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyChatappTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyChatappTemplateResponseBody) *ModifyChatappTemplateResponse
	GetBody() *ModifyChatappTemplateResponseBody
}

type ModifyChatappTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyChatappTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyChatappTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyChatappTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyChatappTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyChatappTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyChatappTemplateResponse) GetBody() *ModifyChatappTemplateResponseBody {
	return s.Body
}

func (s *ModifyChatappTemplateResponse) SetHeaders(v map[string]*string) *ModifyChatappTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyChatappTemplateResponse) SetStatusCode(v int32) *ModifyChatappTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyChatappTemplateResponse) SetBody(v *ModifyChatappTemplateResponseBody) *ModifyChatappTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyChatappTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
