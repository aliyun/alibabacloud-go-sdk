// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodePoolComponentInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNodePoolComponentInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNodePoolComponentInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNodePoolComponentInstanceResponseBody) *DeleteNodePoolComponentInstanceResponse
	GetBody() *DeleteNodePoolComponentInstanceResponseBody
}

type DeleteNodePoolComponentInstanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodePoolComponentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodePoolComponentInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodePoolComponentInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodePoolComponentInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNodePoolComponentInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNodePoolComponentInstanceResponse) GetBody() *DeleteNodePoolComponentInstanceResponseBody {
	return s.Body
}

func (s *DeleteNodePoolComponentInstanceResponse) SetHeaders(v map[string]*string) *DeleteNodePoolComponentInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodePoolComponentInstanceResponse) SetStatusCode(v int32) *DeleteNodePoolComponentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodePoolComponentInstanceResponse) SetBody(v *DeleteNodePoolComponentInstanceResponseBody) *DeleteNodePoolComponentInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteNodePoolComponentInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
