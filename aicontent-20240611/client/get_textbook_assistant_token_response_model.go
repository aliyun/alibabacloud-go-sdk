// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextbookAssistantTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTextbookAssistantTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTextbookAssistantTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetTextbookAssistantTokenResponseBody) *GetTextbookAssistantTokenResponse
	GetBody() *GetTextbookAssistantTokenResponseBody
}

type GetTextbookAssistantTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextbookAssistantTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextbookAssistantTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTextbookAssistantTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTextbookAssistantTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTextbookAssistantTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTextbookAssistantTokenResponse) GetBody() *GetTextbookAssistantTokenResponseBody {
	return s.Body
}

func (s *GetTextbookAssistantTokenResponse) SetHeaders(v map[string]*string) *GetTextbookAssistantTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTextbookAssistantTokenResponse) SetStatusCode(v int32) *GetTextbookAssistantTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextbookAssistantTokenResponse) SetBody(v *GetTextbookAssistantTokenResponseBody) *GetTextbookAssistantTokenResponse {
	s.Body = v
	return s
}

func (s *GetTextbookAssistantTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
