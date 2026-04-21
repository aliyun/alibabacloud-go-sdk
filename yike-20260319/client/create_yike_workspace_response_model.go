// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateYikeWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateYikeWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateYikeWorkspaceResponseBody) *CreateYikeWorkspaceResponse
	GetBody() *CreateYikeWorkspaceResponseBody
}

type CreateYikeWorkspaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYikeWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYikeWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateYikeWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateYikeWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateYikeWorkspaceResponse) GetBody() *CreateYikeWorkspaceResponseBody {
	return s.Body
}

func (s *CreateYikeWorkspaceResponse) SetHeaders(v map[string]*string) *CreateYikeWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateYikeWorkspaceResponse) SetStatusCode(v int32) *CreateYikeWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYikeWorkspaceResponse) SetBody(v *CreateYikeWorkspaceResponseBody) *CreateYikeWorkspaceResponse {
	s.Body = v
	return s
}

func (s *CreateYikeWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
