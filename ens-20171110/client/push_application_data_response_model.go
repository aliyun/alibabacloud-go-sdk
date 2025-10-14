// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushApplicationDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushApplicationDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushApplicationDataResponse
	GetStatusCode() *int32
	SetBody(v *PushApplicationDataResponseBody) *PushApplicationDataResponse
	GetBody() *PushApplicationDataResponseBody
}

type PushApplicationDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushApplicationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushApplicationDataResponse) String() string {
	return dara.Prettify(s)
}

func (s PushApplicationDataResponse) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushApplicationDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushApplicationDataResponse) GetBody() *PushApplicationDataResponseBody {
	return s.Body
}

func (s *PushApplicationDataResponse) SetHeaders(v map[string]*string) *PushApplicationDataResponse {
	s.Headers = v
	return s
}

func (s *PushApplicationDataResponse) SetStatusCode(v int32) *PushApplicationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushApplicationDataResponse) SetBody(v *PushApplicationDataResponseBody) *PushApplicationDataResponse {
	s.Body = v
	return s
}

func (s *PushApplicationDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
