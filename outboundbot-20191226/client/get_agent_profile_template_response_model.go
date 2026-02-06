// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentProfileTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentProfileTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentProfileTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentProfileTemplateResponseBody) *GetAgentProfileTemplateResponse
	GetBody() *GetAgentProfileTemplateResponseBody
}

type GetAgentProfileTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentProfileTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentProfileTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentProfileTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetAgentProfileTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentProfileTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentProfileTemplateResponse) GetBody() *GetAgentProfileTemplateResponseBody {
	return s.Body
}

func (s *GetAgentProfileTemplateResponse) SetHeaders(v map[string]*string) *GetAgentProfileTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetAgentProfileTemplateResponse) SetStatusCode(v int32) *GetAgentProfileTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentProfileTemplateResponse) SetBody(v *GetAgentProfileTemplateResponseBody) *GetAgentProfileTemplateResponse {
	s.Body = v
	return s
}

func (s *GetAgentProfileTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
