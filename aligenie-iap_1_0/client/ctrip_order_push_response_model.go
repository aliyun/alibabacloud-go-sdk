// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCtripOrderPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CtripOrderPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CtripOrderPushResponse
	GetStatusCode() *int32
	SetBody(v *CtripOrderPushResponseBody) *CtripOrderPushResponse
	GetBody() *CtripOrderPushResponseBody
}

type CtripOrderPushResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CtripOrderPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CtripOrderPushResponse) String() string {
	return dara.Prettify(s)
}

func (s CtripOrderPushResponse) GoString() string {
	return s.String()
}

func (s *CtripOrderPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CtripOrderPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CtripOrderPushResponse) GetBody() *CtripOrderPushResponseBody {
	return s.Body
}

func (s *CtripOrderPushResponse) SetHeaders(v map[string]*string) *CtripOrderPushResponse {
	s.Headers = v
	return s
}

func (s *CtripOrderPushResponse) SetStatusCode(v int32) *CtripOrderPushResponse {
	s.StatusCode = &v
	return s
}

func (s *CtripOrderPushResponse) SetBody(v *CtripOrderPushResponseBody) *CtripOrderPushResponse {
	s.Body = v
	return s
}

func (s *CtripOrderPushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
