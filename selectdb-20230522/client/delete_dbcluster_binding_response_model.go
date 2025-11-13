// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBClusterBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBClusterBindingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBClusterBindingResponseBody) *DeleteDBClusterBindingResponse
	GetBody() *DeleteDBClusterBindingResponseBody
}

type DeleteDBClusterBindingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBClusterBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBClusterBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterBindingResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBClusterBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBClusterBindingResponse) GetBody() *DeleteDBClusterBindingResponseBody {
	return s.Body
}

func (s *DeleteDBClusterBindingResponse) SetHeaders(v map[string]*string) *DeleteDBClusterBindingResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterBindingResponse) SetStatusCode(v int32) *DeleteDBClusterBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBClusterBindingResponse) SetBody(v *DeleteDBClusterBindingResponseBody) *DeleteDBClusterBindingResponse {
	s.Body = v
	return s
}

func (s *DeleteDBClusterBindingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
