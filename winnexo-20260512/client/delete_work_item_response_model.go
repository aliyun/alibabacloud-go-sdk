// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkItemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkItemResponseBody) *DeleteWorkItemResponse
	GetBody() *DeleteWorkItemResponseBody
}

type DeleteWorkItemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkItemResponse) GetBody() *DeleteWorkItemResponseBody {
	return s.Body
}

func (s *DeleteWorkItemResponse) SetHeaders(v map[string]*string) *DeleteWorkItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkItemResponse) SetStatusCode(v int32) *DeleteWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkItemResponse) SetBody(v *DeleteWorkItemResponseBody) *DeleteWorkItemResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
