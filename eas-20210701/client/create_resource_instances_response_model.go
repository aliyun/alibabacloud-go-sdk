// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceInstancesResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceInstancesResponseBody) *CreateResourceInstancesResponse
	GetBody() *CreateResourceInstancesResponseBody
}

type CreateResourceInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceInstancesResponse) GetBody() *CreateResourceInstancesResponseBody {
	return s.Body
}

func (s *CreateResourceInstancesResponse) SetHeaders(v map[string]*string) *CreateResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceInstancesResponse) SetStatusCode(v int32) *CreateResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceInstancesResponse) SetBody(v *CreateResourceInstancesResponseBody) *CreateResourceInstancesResponse {
	s.Body = v
	return s
}

func (s *CreateResourceInstancesResponse) Validate() error {
	return dara.Validate(s)
}
