// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateARMServerInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateARMServerInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateARMServerInstancesResponse
	GetStatusCode() *int32
	SetBody(v *CreateARMServerInstancesResponseBody) *CreateARMServerInstancesResponse
	GetBody() *CreateARMServerInstancesResponseBody
}

type CreateARMServerInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateARMServerInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateARMServerInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateARMServerInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateARMServerInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateARMServerInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateARMServerInstancesResponse) GetBody() *CreateARMServerInstancesResponseBody {
	return s.Body
}

func (s *CreateARMServerInstancesResponse) SetHeaders(v map[string]*string) *CreateARMServerInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateARMServerInstancesResponse) SetStatusCode(v int32) *CreateARMServerInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateARMServerInstancesResponse) SetBody(v *CreateARMServerInstancesResponseBody) *CreateARMServerInstancesResponse {
	s.Body = v
	return s
}

func (s *CreateARMServerInstancesResponse) Validate() error {
	return dara.Validate(s)
}
