// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirtualHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirtualHostResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirtualHostResponseBody) *CreateVirtualHostResponse
	GetBody() *CreateVirtualHostResponseBody
}

type CreateVirtualHostResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualHostResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualHostResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirtualHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirtualHostResponse) GetBody() *CreateVirtualHostResponseBody {
	return s.Body
}

func (s *CreateVirtualHostResponse) SetHeaders(v map[string]*string) *CreateVirtualHostResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualHostResponse) SetStatusCode(v int32) *CreateVirtualHostResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualHostResponse) SetBody(v *CreateVirtualHostResponseBody) *CreateVirtualHostResponse {
	s.Body = v
	return s
}

func (s *CreateVirtualHostResponse) Validate() error {
	return dara.Validate(s)
}
