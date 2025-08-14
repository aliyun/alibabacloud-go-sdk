// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocSendDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PocSendDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PocSendDataResponse
	GetStatusCode() *int32
	SetBody(v *PocSendDataResponseBody) *PocSendDataResponse
	GetBody() *PocSendDataResponseBody
}

type PocSendDataResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PocSendDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PocSendDataResponse) String() string {
	return dara.Prettify(s)
}

func (s PocSendDataResponse) GoString() string {
	return s.String()
}

func (s *PocSendDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PocSendDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PocSendDataResponse) GetBody() *PocSendDataResponseBody {
	return s.Body
}

func (s *PocSendDataResponse) SetHeaders(v map[string]*string) *PocSendDataResponse {
	s.Headers = v
	return s
}

func (s *PocSendDataResponse) SetStatusCode(v int32) *PocSendDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PocSendDataResponse) SetBody(v *PocSendDataResponseBody) *PocSendDataResponse {
	s.Body = v
	return s
}

func (s *PocSendDataResponse) Validate() error {
	return dara.Validate(s)
}
