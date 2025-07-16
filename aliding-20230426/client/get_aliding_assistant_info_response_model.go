// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlidingAssistantInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlidingAssistantInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlidingAssistantInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAlidingAssistantInfoResponseBody) *GetAlidingAssistantInfoResponse
	GetBody() *GetAlidingAssistantInfoResponseBody
}

type GetAlidingAssistantInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlidingAssistantInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlidingAssistantInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlidingAssistantInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAlidingAssistantInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlidingAssistantInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlidingAssistantInfoResponse) GetBody() *GetAlidingAssistantInfoResponseBody {
	return s.Body
}

func (s *GetAlidingAssistantInfoResponse) SetHeaders(v map[string]*string) *GetAlidingAssistantInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAlidingAssistantInfoResponse) SetStatusCode(v int32) *GetAlidingAssistantInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlidingAssistantInfoResponse) SetBody(v *GetAlidingAssistantInfoResponseBody) *GetAlidingAssistantInfoResponse {
	s.Body = v
	return s
}

func (s *GetAlidingAssistantInfoResponse) Validate() error {
	return dara.Validate(s)
}
