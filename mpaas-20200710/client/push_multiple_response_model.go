// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMultipleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushMultipleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushMultipleResponse
	GetStatusCode() *int32
	SetBody(v *PushMultipleResponseBody) *PushMultipleResponse
	GetBody() *PushMultipleResponseBody
}

type PushMultipleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushMultipleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushMultipleResponse) String() string {
	return dara.Prettify(s)
}

func (s PushMultipleResponse) GoString() string {
	return s.String()
}

func (s *PushMultipleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushMultipleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushMultipleResponse) GetBody() *PushMultipleResponseBody {
	return s.Body
}

func (s *PushMultipleResponse) SetHeaders(v map[string]*string) *PushMultipleResponse {
	s.Headers = v
	return s
}

func (s *PushMultipleResponse) SetStatusCode(v int32) *PushMultipleResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMultipleResponse) SetBody(v *PushMultipleResponseBody) *PushMultipleResponse {
	s.Body = v
	return s
}

func (s *PushMultipleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
