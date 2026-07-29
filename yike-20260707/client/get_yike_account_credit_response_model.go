// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAccountCreditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeAccountCreditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeAccountCreditResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeAccountCreditResponseBody) *GetYikeAccountCreditResponse
	GetBody() *GetYikeAccountCreditResponseBody
}

type GetYikeAccountCreditResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeAccountCreditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeAccountCreditResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAccountCreditResponse) GoString() string {
	return s.String()
}

func (s *GetYikeAccountCreditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeAccountCreditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeAccountCreditResponse) GetBody() *GetYikeAccountCreditResponseBody {
	return s.Body
}

func (s *GetYikeAccountCreditResponse) SetHeaders(v map[string]*string) *GetYikeAccountCreditResponse {
	s.Headers = v
	return s
}

func (s *GetYikeAccountCreditResponse) SetStatusCode(v int32) *GetYikeAccountCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeAccountCreditResponse) SetBody(v *GetYikeAccountCreditResponseBody) *GetYikeAccountCreditResponse {
	s.Body = v
	return s
}

func (s *GetYikeAccountCreditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
