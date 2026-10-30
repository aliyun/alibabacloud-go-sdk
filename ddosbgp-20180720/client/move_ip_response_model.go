// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveIpResponse
	GetStatusCode() *int32
	SetBody(v *MoveIpResponseBody) *MoveIpResponse
	GetBody() *MoveIpResponseBody
}

type MoveIpResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveIpResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveIpResponse) GoString() string {
	return s.String()
}

func (s *MoveIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveIpResponse) GetBody() *MoveIpResponseBody {
	return s.Body
}

func (s *MoveIpResponse) SetHeaders(v map[string]*string) *MoveIpResponse {
	s.Headers = v
	return s
}

func (s *MoveIpResponse) SetStatusCode(v int32) *MoveIpResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveIpResponse) SetBody(v *MoveIpResponseBody) *MoveIpResponse {
	s.Body = v
	return s
}

func (s *MoveIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
