// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappUserNameSuggestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWhatsappUserNameSuggestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWhatsappUserNameSuggestionsResponse
	GetStatusCode() *int32
	SetBody(v *GetWhatsappUserNameSuggestionsResponseBody) *GetWhatsappUserNameSuggestionsResponse
	GetBody() *GetWhatsappUserNameSuggestionsResponseBody
}

type GetWhatsappUserNameSuggestionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWhatsappUserNameSuggestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWhatsappUserNameSuggestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappUserNameSuggestionsResponse) GoString() string {
	return s.String()
}

func (s *GetWhatsappUserNameSuggestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWhatsappUserNameSuggestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWhatsappUserNameSuggestionsResponse) GetBody() *GetWhatsappUserNameSuggestionsResponseBody {
	return s.Body
}

func (s *GetWhatsappUserNameSuggestionsResponse) SetHeaders(v map[string]*string) *GetWhatsappUserNameSuggestionsResponse {
	s.Headers = v
	return s
}

func (s *GetWhatsappUserNameSuggestionsResponse) SetStatusCode(v int32) *GetWhatsappUserNameSuggestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWhatsappUserNameSuggestionsResponse) SetBody(v *GetWhatsappUserNameSuggestionsResponseBody) *GetWhatsappUserNameSuggestionsResponse {
	s.Body = v
	return s
}

func (s *GetWhatsappUserNameSuggestionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
