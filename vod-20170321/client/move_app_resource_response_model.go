// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveAppResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveAppResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveAppResourceResponse
	GetStatusCode() *int32
	SetBody(v *MoveAppResourceResponseBody) *MoveAppResourceResponse
	GetBody() *MoveAppResourceResponseBody
}

type MoveAppResourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveAppResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveAppResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveAppResourceResponse) GoString() string {
	return s.String()
}

func (s *MoveAppResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveAppResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveAppResourceResponse) GetBody() *MoveAppResourceResponseBody {
	return s.Body
}

func (s *MoveAppResourceResponse) SetHeaders(v map[string]*string) *MoveAppResourceResponse {
	s.Headers = v
	return s
}

func (s *MoveAppResourceResponse) SetStatusCode(v int32) *MoveAppResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveAppResourceResponse) SetBody(v *MoveAppResourceResponseBody) *MoveAppResourceResponse {
	s.Body = v
	return s
}

func (s *MoveAppResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
