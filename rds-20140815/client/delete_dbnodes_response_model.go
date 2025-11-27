// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBNodesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBNodesResponseBody) *DeleteDBNodesResponse
	GetBody() *DeleteDBNodesResponseBody
}

type DeleteDBNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBNodesResponse) GetBody() *DeleteDBNodesResponseBody {
	return s.Body
}

func (s *DeleteDBNodesResponse) SetHeaders(v map[string]*string) *DeleteDBNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBNodesResponse) SetStatusCode(v int32) *DeleteDBNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBNodesResponse) SetBody(v *DeleteDBNodesResponseBody) *DeleteDBNodesResponse {
	s.Body = v
	return s
}

func (s *DeleteDBNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
