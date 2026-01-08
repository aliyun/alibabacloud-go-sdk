// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGroupByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGroupByIdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGroupByIdResponseBody) *DeleteGroupByIdResponse
	GetBody() *DeleteGroupByIdResponseBody
}

type DeleteGroupByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGroupByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGroupByIdResponse) GetBody() *DeleteGroupByIdResponseBody {
	return s.Body
}

func (s *DeleteGroupByIdResponse) SetHeaders(v map[string]*string) *DeleteGroupByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupByIdResponse) SetStatusCode(v int32) *DeleteGroupByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupByIdResponse) SetBody(v *DeleteGroupByIdResponseBody) *DeleteGroupByIdResponse {
	s.Body = v
	return s
}

func (s *DeleteGroupByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
