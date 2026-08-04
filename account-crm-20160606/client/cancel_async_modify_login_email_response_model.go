// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncModifyLoginEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAsyncModifyLoginEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAsyncModifyLoginEmailResponse
	GetStatusCode() *int32
	SetBody(v *CancelAsyncModifyLoginEmailResponseBody) *CancelAsyncModifyLoginEmailResponse
	GetBody() *CancelAsyncModifyLoginEmailResponseBody
}

type CancelAsyncModifyLoginEmailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAsyncModifyLoginEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAsyncModifyLoginEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncModifyLoginEmailResponse) GoString() string {
	return s.String()
}

func (s *CancelAsyncModifyLoginEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAsyncModifyLoginEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAsyncModifyLoginEmailResponse) GetBody() *CancelAsyncModifyLoginEmailResponseBody {
	return s.Body
}

func (s *CancelAsyncModifyLoginEmailResponse) SetHeaders(v map[string]*string) *CancelAsyncModifyLoginEmailResponse {
	s.Headers = v
	return s
}

func (s *CancelAsyncModifyLoginEmailResponse) SetStatusCode(v int32) *CancelAsyncModifyLoginEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAsyncModifyLoginEmailResponse) SetBody(v *CancelAsyncModifyLoginEmailResponseBody) *CancelAsyncModifyLoginEmailResponse {
	s.Body = v
	return s
}

func (s *CancelAsyncModifyLoginEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
