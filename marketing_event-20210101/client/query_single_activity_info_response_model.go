// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySingleActivityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySingleActivityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySingleActivityInfoResponse
	GetStatusCode() *int32
	SetBody(v *QuerySingleActivityInfoResponseBody) *QuerySingleActivityInfoResponse
	GetBody() *QuerySingleActivityInfoResponseBody
}

type QuerySingleActivityInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySingleActivityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySingleActivityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleActivityInfoResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleActivityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySingleActivityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySingleActivityInfoResponse) GetBody() *QuerySingleActivityInfoResponseBody {
	return s.Body
}

func (s *QuerySingleActivityInfoResponse) SetHeaders(v map[string]*string) *QuerySingleActivityInfoResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleActivityInfoResponse) SetStatusCode(v int32) *QuerySingleActivityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySingleActivityInfoResponse) SetBody(v *QuerySingleActivityInfoResponseBody) *QuerySingleActivityInfoResponse {
	s.Body = v
	return s
}

func (s *QuerySingleActivityInfoResponse) Validate() error {
	return dara.Validate(s)
}
