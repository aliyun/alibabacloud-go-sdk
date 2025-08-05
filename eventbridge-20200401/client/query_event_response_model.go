// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEventResponse
	GetStatusCode() *int32
	SetBody(v *QueryEventResponseBody) *QueryEventResponse
	GetBody() *QueryEventResponseBody
}

type QueryEventResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEventResponse) GoString() string {
	return s.String()
}

func (s *QueryEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEventResponse) GetBody() *QueryEventResponseBody {
	return s.Body
}

func (s *QueryEventResponse) SetHeaders(v map[string]*string) *QueryEventResponse {
	s.Headers = v
	return s
}

func (s *QueryEventResponse) SetStatusCode(v int32) *QueryEventResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventResponse) SetBody(v *QueryEventResponseBody) *QueryEventResponse {
	s.Body = v
	return s
}

func (s *QueryEventResponse) Validate() error {
	return dara.Validate(s)
}
