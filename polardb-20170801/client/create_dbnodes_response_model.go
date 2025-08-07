// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBNodesResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBNodesResponseBody) *CreateDBNodesResponse
	GetBody() *CreateDBNodesResponseBody
}

type CreateDBNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBNodesResponse) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBNodesResponse) GetBody() *CreateDBNodesResponseBody {
	return s.Body
}

func (s *CreateDBNodesResponse) SetHeaders(v map[string]*string) *CreateDBNodesResponse {
	s.Headers = v
	return s
}

func (s *CreateDBNodesResponse) SetStatusCode(v int32) *CreateDBNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBNodesResponse) SetBody(v *CreateDBNodesResponseBody) *CreateDBNodesResponse {
	s.Body = v
	return s
}

func (s *CreateDBNodesResponse) Validate() error {
	return dara.Validate(s)
}
