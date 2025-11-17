// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccelerationOfWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccelerationOfWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccelerationOfWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *ListAccelerationOfWorkspaceResponseBody) *ListAccelerationOfWorkspaceResponse
	GetBody() *ListAccelerationOfWorkspaceResponseBody
}

type ListAccelerationOfWorkspaceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccelerationOfWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccelerationOfWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccelerationOfWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *ListAccelerationOfWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccelerationOfWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccelerationOfWorkspaceResponse) GetBody() *ListAccelerationOfWorkspaceResponseBody {
	return s.Body
}

func (s *ListAccelerationOfWorkspaceResponse) SetHeaders(v map[string]*string) *ListAccelerationOfWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *ListAccelerationOfWorkspaceResponse) SetStatusCode(v int32) *ListAccelerationOfWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccelerationOfWorkspaceResponse) SetBody(v *ListAccelerationOfWorkspaceResponseBody) *ListAccelerationOfWorkspaceResponse {
	s.Body = v
	return s
}

func (s *ListAccelerationOfWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
