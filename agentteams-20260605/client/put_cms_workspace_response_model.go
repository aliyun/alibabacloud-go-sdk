// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCmsWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutCmsWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutCmsWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *PutCmsWorkspaceResponseBody) *PutCmsWorkspaceResponse
	GetBody() *PutCmsWorkspaceResponseBody
}

type PutCmsWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutCmsWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutCmsWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s PutCmsWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *PutCmsWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutCmsWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutCmsWorkspaceResponse) GetBody() *PutCmsWorkspaceResponseBody {
	return s.Body
}

func (s *PutCmsWorkspaceResponse) SetHeaders(v map[string]*string) *PutCmsWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *PutCmsWorkspaceResponse) SetStatusCode(v int32) *PutCmsWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *PutCmsWorkspaceResponse) SetBody(v *PutCmsWorkspaceResponseBody) *PutCmsWorkspaceResponse {
	s.Body = v
	return s
}

func (s *PutCmsWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
