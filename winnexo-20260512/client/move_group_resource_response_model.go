// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveGroupResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveGroupResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveGroupResourceResponse
	GetStatusCode() *int32
	SetBody(v *MoveGroupResourceResponseBody) *MoveGroupResourceResponse
	GetBody() *MoveGroupResourceResponseBody
}

type MoveGroupResourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveGroupResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveGroupResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveGroupResourceResponse) GoString() string {
	return s.String()
}

func (s *MoveGroupResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveGroupResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveGroupResourceResponse) GetBody() *MoveGroupResourceResponseBody {
	return s.Body
}

func (s *MoveGroupResourceResponse) SetHeaders(v map[string]*string) *MoveGroupResourceResponse {
	s.Headers = v
	return s
}

func (s *MoveGroupResourceResponse) SetStatusCode(v int32) *MoveGroupResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveGroupResourceResponse) SetBody(v *MoveGroupResourceResponseBody) *MoveGroupResourceResponse {
	s.Body = v
	return s
}

func (s *MoveGroupResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
