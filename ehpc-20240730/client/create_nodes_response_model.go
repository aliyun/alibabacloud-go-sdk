// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNodesResponse
	GetStatusCode() *int32
	SetBody(v *CreateNodesResponseBody) *CreateNodesResponse
	GetBody() *CreateNodesResponseBody
}

type CreateNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNodesResponse) GoString() string {
	return s.String()
}

func (s *CreateNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNodesResponse) GetBody() *CreateNodesResponseBody {
	return s.Body
}

func (s *CreateNodesResponse) SetHeaders(v map[string]*string) *CreateNodesResponse {
	s.Headers = v
	return s
}

func (s *CreateNodesResponse) SetStatusCode(v int32) *CreateNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodesResponse) SetBody(v *CreateNodesResponseBody) *CreateNodesResponse {
	s.Body = v
	return s
}

func (s *CreateNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
