// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveNodeResponse
	GetStatusCode() *int32
	SetBody(v *MoveNodeResponseBody) *MoveNodeResponse
	GetBody() *MoveNodeResponseBody
}

type MoveNodeResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveNodeResponse) GoString() string {
	return s.String()
}

func (s *MoveNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveNodeResponse) GetBody() *MoveNodeResponseBody {
	return s.Body
}

func (s *MoveNodeResponse) SetHeaders(v map[string]*string) *MoveNodeResponse {
	s.Headers = v
	return s
}

func (s *MoveNodeResponse) SetStatusCode(v int32) *MoveNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveNodeResponse) SetBody(v *MoveNodeResponseBody) *MoveNodeResponse {
	s.Body = v
	return s
}

func (s *MoveNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
