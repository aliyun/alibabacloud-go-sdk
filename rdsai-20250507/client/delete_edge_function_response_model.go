// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEdgeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEdgeFunctionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEdgeFunctionResponseBody) *DeleteEdgeFunctionResponse
	GetBody() *DeleteEdgeFunctionResponseBody
}

type DeleteEdgeFunctionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEdgeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEdgeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteEdgeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEdgeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEdgeFunctionResponse) GetBody() *DeleteEdgeFunctionResponseBody {
	return s.Body
}

func (s *DeleteEdgeFunctionResponse) SetHeaders(v map[string]*string) *DeleteEdgeFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteEdgeFunctionResponse) SetStatusCode(v int32) *DeleteEdgeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEdgeFunctionResponse) SetBody(v *DeleteEdgeFunctionResponseBody) *DeleteEdgeFunctionResponse {
	s.Body = v
	return s
}

func (s *DeleteEdgeFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
