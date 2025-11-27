// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmNotifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmNotifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmNotifyResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmNotifyResponseBody) *ConfirmNotifyResponse
	GetBody() *ConfirmNotifyResponseBody
}

type ConfirmNotifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmNotifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmNotifyResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmNotifyResponse) GoString() string {
	return s.String()
}

func (s *ConfirmNotifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmNotifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmNotifyResponse) GetBody() *ConfirmNotifyResponseBody {
	return s.Body
}

func (s *ConfirmNotifyResponse) SetHeaders(v map[string]*string) *ConfirmNotifyResponse {
	s.Headers = v
	return s
}

func (s *ConfirmNotifyResponse) SetStatusCode(v int32) *ConfirmNotifyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmNotifyResponse) SetBody(v *ConfirmNotifyResponseBody) *ConfirmNotifyResponse {
	s.Body = v
	return s
}

func (s *ConfirmNotifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
