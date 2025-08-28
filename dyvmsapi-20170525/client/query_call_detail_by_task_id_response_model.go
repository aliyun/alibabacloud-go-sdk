// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallDetailByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCallDetailByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCallDetailByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryCallDetailByTaskIdResponseBody) *QueryCallDetailByTaskIdResponse
	GetBody() *QueryCallDetailByTaskIdResponseBody
}

type QueryCallDetailByTaskIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallDetailByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallDetailByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCallDetailByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCallDetailByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCallDetailByTaskIdResponse) GetBody() *QueryCallDetailByTaskIdResponseBody {
	return s.Body
}

func (s *QueryCallDetailByTaskIdResponse) SetHeaders(v map[string]*string) *QueryCallDetailByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *QueryCallDetailByTaskIdResponse) SetStatusCode(v int32) *QueryCallDetailByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponse) SetBody(v *QueryCallDetailByTaskIdResponseBody) *QueryCallDetailByTaskIdResponse {
	s.Body = v
	return s
}

func (s *QueryCallDetailByTaskIdResponse) Validate() error {
	return dara.Validate(s)
}
