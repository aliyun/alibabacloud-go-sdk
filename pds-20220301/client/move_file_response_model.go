// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveFileResponse
	GetStatusCode() *int32
	SetBody(v *MoveFileResponseBody) *MoveFileResponse
	GetBody() *MoveFileResponseBody
}

type MoveFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveFileResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveFileResponse) GoString() string {
	return s.String()
}

func (s *MoveFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveFileResponse) GetBody() *MoveFileResponseBody {
	return s.Body
}

func (s *MoveFileResponse) SetHeaders(v map[string]*string) *MoveFileResponse {
	s.Headers = v
	return s
}

func (s *MoveFileResponse) SetStatusCode(v int32) *MoveFileResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveFileResponse) SetBody(v *MoveFileResponseBody) *MoveFileResponse {
	s.Body = v
	return s
}

func (s *MoveFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
