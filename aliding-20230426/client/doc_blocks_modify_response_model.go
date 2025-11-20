// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocBlocksModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocBlocksModifyResponse
	GetStatusCode() *int32
	SetBody(v *DocBlocksModifyResponseBody) *DocBlocksModifyResponse
	GetBody() *DocBlocksModifyResponseBody
}

type DocBlocksModifyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocBlocksModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocBlocksModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksModifyResponse) GoString() string {
	return s.String()
}

func (s *DocBlocksModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocBlocksModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocBlocksModifyResponse) GetBody() *DocBlocksModifyResponseBody {
	return s.Body
}

func (s *DocBlocksModifyResponse) SetHeaders(v map[string]*string) *DocBlocksModifyResponse {
	s.Headers = v
	return s
}

func (s *DocBlocksModifyResponse) SetStatusCode(v int32) *DocBlocksModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DocBlocksModifyResponse) SetBody(v *DocBlocksModifyResponseBody) *DocBlocksModifyResponse {
	s.Body = v
	return s
}

func (s *DocBlocksModifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
