// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNodesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNodesResponseBody) *DeleteNodesResponse
	GetBody() *DeleteNodesResponseBody
}

type DeleteNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNodesResponse) GetBody() *DeleteNodesResponseBody {
	return s.Body
}

func (s *DeleteNodesResponse) SetHeaders(v map[string]*string) *DeleteNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodesResponse) SetStatusCode(v int32) *DeleteNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodesResponse) SetBody(v *DeleteNodesResponseBody) *DeleteNodesResponse {
	s.Body = v
	return s
}

func (s *DeleteNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
