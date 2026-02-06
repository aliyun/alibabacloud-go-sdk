// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConcurrentConversationQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConcurrentConversationQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConcurrentConversationQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetConcurrentConversationQuotaResponseBody) *GetConcurrentConversationQuotaResponse
	GetBody() *GetConcurrentConversationQuotaResponseBody
}

type GetConcurrentConversationQuotaResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConcurrentConversationQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConcurrentConversationQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConcurrentConversationQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetConcurrentConversationQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConcurrentConversationQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConcurrentConversationQuotaResponse) GetBody() *GetConcurrentConversationQuotaResponseBody {
	return s.Body
}

func (s *GetConcurrentConversationQuotaResponse) SetHeaders(v map[string]*string) *GetConcurrentConversationQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetConcurrentConversationQuotaResponse) SetStatusCode(v int32) *GetConcurrentConversationQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConcurrentConversationQuotaResponse) SetBody(v *GetConcurrentConversationQuotaResponseBody) *GetConcurrentConversationQuotaResponse {
	s.Body = v
	return s
}

func (s *GetConcurrentConversationQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
