// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendAsyncModifyLoginEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResendAsyncModifyLoginEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResendAsyncModifyLoginEmailResponse
	GetStatusCode() *int32
	SetBody(v *ResendAsyncModifyLoginEmailResponseBody) *ResendAsyncModifyLoginEmailResponse
	GetBody() *ResendAsyncModifyLoginEmailResponseBody
}

type ResendAsyncModifyLoginEmailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendAsyncModifyLoginEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendAsyncModifyLoginEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s ResendAsyncModifyLoginEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendAsyncModifyLoginEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResendAsyncModifyLoginEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResendAsyncModifyLoginEmailResponse) GetBody() *ResendAsyncModifyLoginEmailResponseBody {
	return s.Body
}

func (s *ResendAsyncModifyLoginEmailResponse) SetHeaders(v map[string]*string) *ResendAsyncModifyLoginEmailResponse {
	s.Headers = v
	return s
}

func (s *ResendAsyncModifyLoginEmailResponse) SetStatusCode(v int32) *ResendAsyncModifyLoginEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendAsyncModifyLoginEmailResponse) SetBody(v *ResendAsyncModifyLoginEmailResponseBody) *ResendAsyncModifyLoginEmailResponse {
	s.Body = v
	return s
}

func (s *ResendAsyncModifyLoginEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
