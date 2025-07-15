// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneGlobalReplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInterveneGlobalReplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInterveneGlobalReplyResponse
	GetStatusCode() *int32
	SetBody(v *GetInterveneGlobalReplyResponseBody) *GetInterveneGlobalReplyResponse
	GetBody() *GetInterveneGlobalReplyResponseBody
}

type GetInterveneGlobalReplyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneGlobalReplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneGlobalReplyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneGlobalReplyResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInterveneGlobalReplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInterveneGlobalReplyResponse) GetBody() *GetInterveneGlobalReplyResponseBody {
	return s.Body
}

func (s *GetInterveneGlobalReplyResponse) SetHeaders(v map[string]*string) *GetInterveneGlobalReplyResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneGlobalReplyResponse) SetStatusCode(v int32) *GetInterveneGlobalReplyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneGlobalReplyResponse) SetBody(v *GetInterveneGlobalReplyResponseBody) *GetInterveneGlobalReplyResponse {
	s.Body = v
	return s
}

func (s *GetInterveneGlobalReplyResponse) Validate() error {
	return dara.Validate(s)
}
