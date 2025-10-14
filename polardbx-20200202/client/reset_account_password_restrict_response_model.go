// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordRestrictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAccountPasswordRestrictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAccountPasswordRestrictResponse
	GetStatusCode() *int32
	SetBody(v *ResetAccountPasswordRestrictResponseBody) *ResetAccountPasswordRestrictResponse
	GetBody() *ResetAccountPasswordRestrictResponseBody
}

type ResetAccountPasswordRestrictResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAccountPasswordRestrictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAccountPasswordRestrictResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordRestrictResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRestrictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAccountPasswordRestrictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAccountPasswordRestrictResponse) GetBody() *ResetAccountPasswordRestrictResponseBody {
	return s.Body
}

func (s *ResetAccountPasswordRestrictResponse) SetHeaders(v map[string]*string) *ResetAccountPasswordRestrictResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountPasswordRestrictResponse) SetStatusCode(v int32) *ResetAccountPasswordRestrictResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAccountPasswordRestrictResponse) SetBody(v *ResetAccountPasswordRestrictResponseBody) *ResetAccountPasswordRestrictResponse {
	s.Body = v
	return s
}

func (s *ResetAccountPasswordRestrictResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
