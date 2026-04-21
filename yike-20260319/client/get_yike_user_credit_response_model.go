// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeUserCreditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeUserCreditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeUserCreditResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeUserCreditResponseBody) *GetYikeUserCreditResponse
	GetBody() *GetYikeUserCreditResponseBody
}

type GetYikeUserCreditResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeUserCreditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeUserCreditResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeUserCreditResponse) GoString() string {
	return s.String()
}

func (s *GetYikeUserCreditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeUserCreditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeUserCreditResponse) GetBody() *GetYikeUserCreditResponseBody {
	return s.Body
}

func (s *GetYikeUserCreditResponse) SetHeaders(v map[string]*string) *GetYikeUserCreditResponse {
	s.Headers = v
	return s
}

func (s *GetYikeUserCreditResponse) SetStatusCode(v int32) *GetYikeUserCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeUserCreditResponse) SetBody(v *GetYikeUserCreditResponseBody) *GetYikeUserCreditResponse {
	s.Body = v
	return s
}

func (s *GetYikeUserCreditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
