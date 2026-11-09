// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGroupIdResponseBody) *DeleteGroupIdResponse
	GetBody() *DeleteGroupIdResponseBody
}

type DeleteGroupIdResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGroupIdResponse) GetBody() *DeleteGroupIdResponseBody {
	return s.Body
}

func (s *DeleteGroupIdResponse) SetHeaders(v map[string]*string) *DeleteGroupIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupIdResponse) SetStatusCode(v int32) *DeleteGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupIdResponse) SetBody(v *DeleteGroupIdResponseBody) *DeleteGroupIdResponse {
	s.Body = v
	return s
}

func (s *DeleteGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
