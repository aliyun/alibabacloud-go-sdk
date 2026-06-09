// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmAppInstanceResponseBody) *ConfirmAppInstanceResponse
	GetBody() *ConfirmAppInstanceResponseBody
}

type ConfirmAppInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConfirmAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmAppInstanceResponse) GetBody() *ConfirmAppInstanceResponseBody {
	return s.Body
}

func (s *ConfirmAppInstanceResponse) SetHeaders(v map[string]*string) *ConfirmAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConfirmAppInstanceResponse) SetStatusCode(v int32) *ConfirmAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmAppInstanceResponse) SetBody(v *ConfirmAppInstanceResponseBody) *ConfirmAppInstanceResponse {
	s.Body = v
	return s
}

func (s *ConfirmAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
