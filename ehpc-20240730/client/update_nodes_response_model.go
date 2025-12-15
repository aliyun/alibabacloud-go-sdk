// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodesResponseBody) *UpdateNodesResponse
	GetBody() *UpdateNodesResponseBody
}

type UpdateNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodesResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodesResponse) GetBody() *UpdateNodesResponseBody {
	return s.Body
}

func (s *UpdateNodesResponse) SetHeaders(v map[string]*string) *UpdateNodesResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodesResponse) SetStatusCode(v int32) *UpdateNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodesResponse) SetBody(v *UpdateNodesResponseBody) *UpdateNodesResponse {
	s.Body = v
	return s
}

func (s *UpdateNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
