// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttackTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAttackTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAttackTargetResponse
	GetStatusCode() *int32
	SetBody(v *CreateAttackTargetResponseBody) *CreateAttackTargetResponse
	GetBody() *CreateAttackTargetResponseBody
}

type CreateAttackTargetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAttackTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAttackTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateAttackTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAttackTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAttackTargetResponse) GetBody() *CreateAttackTargetResponseBody {
	return s.Body
}

func (s *CreateAttackTargetResponse) SetHeaders(v map[string]*string) *CreateAttackTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateAttackTargetResponse) SetStatusCode(v int32) *CreateAttackTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAttackTargetResponse) SetBody(v *CreateAttackTargetResponseBody) *CreateAttackTargetResponse {
	s.Body = v
	return s
}

func (s *CreateAttackTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
