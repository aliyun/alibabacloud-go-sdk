// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstancesResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstancesResponseBody) *CreateInstancesResponse
	GetBody() *CreateInstancesResponseBody
}

type CreateInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstancesResponse) GetBody() *CreateInstancesResponseBody {
	return s.Body
}

func (s *CreateInstancesResponse) SetHeaders(v map[string]*string) *CreateInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateInstancesResponse) SetStatusCode(v int32) *CreateInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstancesResponse) SetBody(v *CreateInstancesResponseBody) *CreateInstancesResponse {
	s.Body = v
	return s
}

func (s *CreateInstancesResponse) Validate() error {
	return dara.Validate(s)
}
