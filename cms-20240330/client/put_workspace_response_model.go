// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *PutWorkspaceResponseBody) *PutWorkspaceResponse
	GetBody() *PutWorkspaceResponseBody
}

type PutWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s PutWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *PutWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutWorkspaceResponse) GetBody() *PutWorkspaceResponseBody {
	return s.Body
}

func (s *PutWorkspaceResponse) SetHeaders(v map[string]*string) *PutWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *PutWorkspaceResponse) SetStatusCode(v int32) *PutWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *PutWorkspaceResponse) SetBody(v *PutWorkspaceResponseBody) *PutWorkspaceResponse {
	s.Body = v
	return s
}

func (s *PutWorkspaceResponse) Validate() error {
	return dara.Validate(s)
}
