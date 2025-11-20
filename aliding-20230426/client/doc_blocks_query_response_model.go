// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocBlocksQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocBlocksQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocBlocksQueryResponse
	GetStatusCode() *int32
	SetBody(v *DocBlocksQueryResponseBody) *DocBlocksQueryResponse
	GetBody() *DocBlocksQueryResponseBody
}

type DocBlocksQueryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocBlocksQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocBlocksQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s DocBlocksQueryResponse) GoString() string {
	return s.String()
}

func (s *DocBlocksQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocBlocksQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocBlocksQueryResponse) GetBody() *DocBlocksQueryResponseBody {
	return s.Body
}

func (s *DocBlocksQueryResponse) SetHeaders(v map[string]*string) *DocBlocksQueryResponse {
	s.Headers = v
	return s
}

func (s *DocBlocksQueryResponse) SetStatusCode(v int32) *DocBlocksQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DocBlocksQueryResponse) SetBody(v *DocBlocksQueryResponseBody) *DocBlocksQueryResponse {
	s.Body = v
	return s
}

func (s *DocBlocksQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
