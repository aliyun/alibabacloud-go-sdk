// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAttackTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAttackTargetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAttackTargetResponseBody) *DeleteAttackTargetResponse
	GetBody() *DeleteAttackTargetResponseBody
}

type DeleteAttackTargetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAttackTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAttackTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteAttackTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAttackTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAttackTargetResponse) GetBody() *DeleteAttackTargetResponseBody {
	return s.Body
}

func (s *DeleteAttackTargetResponse) SetHeaders(v map[string]*string) *DeleteAttackTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteAttackTargetResponse) SetStatusCode(v int32) *DeleteAttackTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAttackTargetResponse) SetBody(v *DeleteAttackTargetResponseBody) *DeleteAttackTargetResponse {
	s.Body = v
	return s
}

func (s *DeleteAttackTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
