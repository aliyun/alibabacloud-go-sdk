// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyLastUsedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyLastUsedResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyLastUsedResponseBody) *GetAccessKeyLastUsedResponse
	GetBody() *GetAccessKeyLastUsedResponseBody
}

type GetAccessKeyLastUsedResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyLastUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyLastUsedResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyLastUsedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyLastUsedResponse) GetBody() *GetAccessKeyLastUsedResponseBody {
	return s.Body
}

func (s *GetAccessKeyLastUsedResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedResponse) SetBody(v *GetAccessKeyLastUsedResponseBody) *GetAccessKeyLastUsedResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyLastUsedResponse) Validate() error {
	return dara.Validate(s)
}
