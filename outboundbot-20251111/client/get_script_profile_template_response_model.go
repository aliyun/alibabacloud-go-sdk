// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptProfileTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScriptProfileTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScriptProfileTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetScriptProfileTemplateResponseBody) *GetScriptProfileTemplateResponse
	GetBody() *GetScriptProfileTemplateResponseBody
}

type GetScriptProfileTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScriptProfileTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScriptProfileTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScriptProfileTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetScriptProfileTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScriptProfileTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScriptProfileTemplateResponse) GetBody() *GetScriptProfileTemplateResponseBody {
	return s.Body
}

func (s *GetScriptProfileTemplateResponse) SetHeaders(v map[string]*string) *GetScriptProfileTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetScriptProfileTemplateResponse) SetStatusCode(v int32) *GetScriptProfileTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScriptProfileTemplateResponse) SetBody(v *GetScriptProfileTemplateResponseBody) *GetScriptProfileTemplateResponse {
	s.Body = v
	return s
}

func (s *GetScriptProfileTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
