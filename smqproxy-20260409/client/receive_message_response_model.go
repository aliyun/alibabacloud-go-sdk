// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReceiveMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReceiveMessageResponse
	GetStatusCode() *int32
	SetBody(v *ReceiveMessageResponseBody) *ReceiveMessageResponse
	GetBody() *ReceiveMessageResponseBody
}

type ReceiveMessageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReceiveMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReceiveMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s ReceiveMessageResponse) GoString() string {
	return s.String()
}

func (s *ReceiveMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReceiveMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReceiveMessageResponse) GetBody() *ReceiveMessageResponseBody {
	return s.Body
}

func (s *ReceiveMessageResponse) SetHeaders(v map[string]*string) *ReceiveMessageResponse {
	s.Headers = v
	return s
}

func (s *ReceiveMessageResponse) SetStatusCode(v int32) *ReceiveMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReceiveMessageResponse) SetBody(v *ReceiveMessageResponseBody) *ReceiveMessageResponse {
	s.Body = v
	return s
}

func (s *ReceiveMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
