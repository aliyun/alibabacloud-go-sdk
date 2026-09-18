// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackTargetResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackTargetResponseBody) *GetAttackTargetResponse
	GetBody() *GetAttackTargetResponseBody
}

type GetAttackTargetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackTargetResponse) GoString() string {
	return s.String()
}

func (s *GetAttackTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackTargetResponse) GetBody() *GetAttackTargetResponseBody {
	return s.Body
}

func (s *GetAttackTargetResponse) SetHeaders(v map[string]*string) *GetAttackTargetResponse {
	s.Headers = v
	return s
}

func (s *GetAttackTargetResponse) SetStatusCode(v int32) *GetAttackTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackTargetResponse) SetBody(v *GetAttackTargetResponseBody) *GetAttackTargetResponse {
	s.Body = v
	return s
}

func (s *GetAttackTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
