// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVehicle5ItemQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Vehicle5ItemQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Vehicle5ItemQueryResponse
	GetStatusCode() *int32
	SetBody(v *Vehicle5ItemQueryResponseBody) *Vehicle5ItemQueryResponse
	GetBody() *Vehicle5ItemQueryResponseBody
}

type Vehicle5ItemQueryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Vehicle5ItemQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Vehicle5ItemQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s Vehicle5ItemQueryResponse) GoString() string {
	return s.String()
}

func (s *Vehicle5ItemQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Vehicle5ItemQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Vehicle5ItemQueryResponse) GetBody() *Vehicle5ItemQueryResponseBody {
	return s.Body
}

func (s *Vehicle5ItemQueryResponse) SetHeaders(v map[string]*string) *Vehicle5ItemQueryResponse {
	s.Headers = v
	return s
}

func (s *Vehicle5ItemQueryResponse) SetStatusCode(v int32) *Vehicle5ItemQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *Vehicle5ItemQueryResponse) SetBody(v *Vehicle5ItemQueryResponseBody) *Vehicle5ItemQueryResponse {
	s.Body = v
	return s
}

func (s *Vehicle5ItemQueryResponse) Validate() error {
	return dara.Validate(s)
}
