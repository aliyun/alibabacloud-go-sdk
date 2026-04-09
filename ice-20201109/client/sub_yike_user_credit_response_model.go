// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubYikeUserCreditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubYikeUserCreditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubYikeUserCreditResponse
	GetStatusCode() *int32
	SetBody(v *SubYikeUserCreditResponseBody) *SubYikeUserCreditResponse
	GetBody() *SubYikeUserCreditResponseBody
}

type SubYikeUserCreditResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubYikeUserCreditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubYikeUserCreditResponse) String() string {
	return dara.Prettify(s)
}

func (s SubYikeUserCreditResponse) GoString() string {
	return s.String()
}

func (s *SubYikeUserCreditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubYikeUserCreditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubYikeUserCreditResponse) GetBody() *SubYikeUserCreditResponseBody {
	return s.Body
}

func (s *SubYikeUserCreditResponse) SetHeaders(v map[string]*string) *SubYikeUserCreditResponse {
	s.Headers = v
	return s
}

func (s *SubYikeUserCreditResponse) SetStatusCode(v int32) *SubYikeUserCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *SubYikeUserCreditResponse) SetBody(v *SubYikeUserCreditResponseBody) *SubYikeUserCreditResponse {
	s.Body = v
	return s
}

func (s *SubYikeUserCreditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
