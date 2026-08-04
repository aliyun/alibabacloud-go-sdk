// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountProfileInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountProfileInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountProfileInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountProfileInfoResponseBody) *QueryAccountProfileInfoResponse
	GetBody() *QueryAccountProfileInfoResponseBody
}

type QueryAccountProfileInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountProfileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountProfileInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountProfileInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountProfileInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountProfileInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountProfileInfoResponse) GetBody() *QueryAccountProfileInfoResponseBody {
	return s.Body
}

func (s *QueryAccountProfileInfoResponse) SetHeaders(v map[string]*string) *QueryAccountProfileInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountProfileInfoResponse) SetStatusCode(v int32) *QueryAccountProfileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountProfileInfoResponse) SetBody(v *QueryAccountProfileInfoResponseBody) *QueryAccountProfileInfoResponse {
	s.Body = v
	return s
}

func (s *QueryAccountProfileInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
