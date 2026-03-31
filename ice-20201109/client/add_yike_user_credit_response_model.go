// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddYikeUserCreditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddYikeUserCreditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddYikeUserCreditResponse
	GetStatusCode() *int32
	SetBody(v *AddYikeUserCreditResponseBody) *AddYikeUserCreditResponse
	GetBody() *AddYikeUserCreditResponseBody
}

type AddYikeUserCreditResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddYikeUserCreditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddYikeUserCreditResponse) String() string {
	return dara.Prettify(s)
}

func (s AddYikeUserCreditResponse) GoString() string {
	return s.String()
}

func (s *AddYikeUserCreditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddYikeUserCreditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddYikeUserCreditResponse) GetBody() *AddYikeUserCreditResponseBody {
	return s.Body
}

func (s *AddYikeUserCreditResponse) SetHeaders(v map[string]*string) *AddYikeUserCreditResponse {
	s.Headers = v
	return s
}

func (s *AddYikeUserCreditResponse) SetStatusCode(v int32) *AddYikeUserCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *AddYikeUserCreditResponse) SetBody(v *AddYikeUserCreditResponseBody) *AddYikeUserCreditResponse {
	s.Body = v
	return s
}

func (s *AddYikeUserCreditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
