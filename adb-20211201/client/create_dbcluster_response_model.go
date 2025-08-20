// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBClusterResponseBody) *CreateDBClusterResponse
	GetBody() *CreateDBClusterResponseBody
}

type CreateDBClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBClusterResponse) GetBody() *CreateDBClusterResponseBody {
	return s.Body
}

func (s *CreateDBClusterResponse) SetHeaders(v map[string]*string) *CreateDBClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterResponse) SetStatusCode(v int32) *CreateDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBClusterResponse) SetBody(v *CreateDBClusterResponseBody) *CreateDBClusterResponse {
	s.Body = v
	return s
}

func (s *CreateDBClusterResponse) Validate() error {
	return dara.Validate(s)
}
