// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventTracesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEventTracesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEventTracesResponse
	GetStatusCode() *int32
	SetBody(v *QueryEventTracesResponseBody) *QueryEventTracesResponse
	GetBody() *QueryEventTracesResponseBody
}

type QueryEventTracesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEventTracesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEventTracesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEventTracesResponse) GoString() string {
	return s.String()
}

func (s *QueryEventTracesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEventTracesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEventTracesResponse) GetBody() *QueryEventTracesResponseBody {
	return s.Body
}

func (s *QueryEventTracesResponse) SetHeaders(v map[string]*string) *QueryEventTracesResponse {
	s.Headers = v
	return s
}

func (s *QueryEventTracesResponse) SetStatusCode(v int32) *QueryEventTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEventTracesResponse) SetBody(v *QueryEventTracesResponseBody) *QueryEventTracesResponse {
	s.Body = v
	return s
}

func (s *QueryEventTracesResponse) Validate() error {
	return dara.Validate(s)
}
