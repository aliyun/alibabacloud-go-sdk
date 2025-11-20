// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocBlocksDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocBlocksDeleteResponse
	GetStatusCode() *int32
	SetBody(v *DocBlocksDeleteResponseBody) *DocBlocksDeleteResponse
	GetBody() *DocBlocksDeleteResponseBody
}

type DocBlocksDeleteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocBlocksDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocBlocksDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksDeleteResponse) GoString() string {
	return s.String()
}

func (s *DocBlocksDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocBlocksDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocBlocksDeleteResponse) GetBody() *DocBlocksDeleteResponseBody {
	return s.Body
}

func (s *DocBlocksDeleteResponse) SetHeaders(v map[string]*string) *DocBlocksDeleteResponse {
	s.Headers = v
	return s
}

func (s *DocBlocksDeleteResponse) SetStatusCode(v int32) *DocBlocksDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *DocBlocksDeleteResponse) SetBody(v *DocBlocksDeleteResponseBody) *DocBlocksDeleteResponse {
	s.Body = v
	return s
}

func (s *DocBlocksDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
