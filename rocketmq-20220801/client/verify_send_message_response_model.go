// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySendMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifySendMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifySendMessageResponse
	GetStatusCode() *int32
	SetBody(v *VerifySendMessageResponseBody) *VerifySendMessageResponse
	GetBody() *VerifySendMessageResponseBody
}

type VerifySendMessageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifySendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifySendMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifySendMessageResponse) GoString() string {
	return s.String()
}

func (s *VerifySendMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifySendMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifySendMessageResponse) GetBody() *VerifySendMessageResponseBody {
	return s.Body
}

func (s *VerifySendMessageResponse) SetHeaders(v map[string]*string) *VerifySendMessageResponse {
	s.Headers = v
	return s
}

func (s *VerifySendMessageResponse) SetStatusCode(v int32) *VerifySendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySendMessageResponse) SetBody(v *VerifySendMessageResponseBody) *VerifySendMessageResponse {
	s.Body = v
	return s
}

func (s *VerifySendMessageResponse) Validate() error {
	return dara.Validate(s)
}
