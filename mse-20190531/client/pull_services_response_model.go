// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PullServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PullServicesResponse
	GetStatusCode() *int32
	SetBody(v *PullServicesResponseBody) *PullServicesResponse
	GetBody() *PullServicesResponseBody
}

type PullServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PullServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PullServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s PullServicesResponse) GoString() string {
	return s.String()
}

func (s *PullServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PullServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PullServicesResponse) GetBody() *PullServicesResponseBody {
	return s.Body
}

func (s *PullServicesResponse) SetHeaders(v map[string]*string) *PullServicesResponse {
	s.Headers = v
	return s
}

func (s *PullServicesResponse) SetStatusCode(v int32) *PullServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *PullServicesResponse) SetBody(v *PullServicesResponseBody) *PullServicesResponse {
	s.Body = v
	return s
}

func (s *PullServicesResponse) Validate() error {
	return dara.Validate(s)
}
