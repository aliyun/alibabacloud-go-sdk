// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSendMailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SingleSendMailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SingleSendMailResponse
	GetStatusCode() *int32
	SetBody(v *SingleSendMailResponseBody) *SingleSendMailResponse
	GetBody() *SingleSendMailResponseBody
}

type SingleSendMailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleSendMailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleSendMailResponse) String() string {
	return dara.Prettify(s)
}

func (s SingleSendMailResponse) GoString() string {
	return s.String()
}

func (s *SingleSendMailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SingleSendMailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SingleSendMailResponse) GetBody() *SingleSendMailResponseBody {
	return s.Body
}

func (s *SingleSendMailResponse) SetHeaders(v map[string]*string) *SingleSendMailResponse {
	s.Headers = v
	return s
}

func (s *SingleSendMailResponse) SetStatusCode(v int32) *SingleSendMailResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleSendMailResponse) SetBody(v *SingleSendMailResponseBody) *SingleSendMailResponse {
	s.Body = v
	return s
}

func (s *SingleSendMailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
