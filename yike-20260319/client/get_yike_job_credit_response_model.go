// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeJobCreditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeJobCreditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeJobCreditResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeJobCreditResponseBody) *GetYikeJobCreditResponse
	GetBody() *GetYikeJobCreditResponseBody
}

type GetYikeJobCreditResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeJobCreditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeJobCreditResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeJobCreditResponse) GoString() string {
	return s.String()
}

func (s *GetYikeJobCreditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeJobCreditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeJobCreditResponse) GetBody() *GetYikeJobCreditResponseBody {
	return s.Body
}

func (s *GetYikeJobCreditResponse) SetHeaders(v map[string]*string) *GetYikeJobCreditResponse {
	s.Headers = v
	return s
}

func (s *GetYikeJobCreditResponse) SetStatusCode(v int32) *GetYikeJobCreditResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeJobCreditResponse) SetBody(v *GetYikeJobCreditResponseBody) *GetYikeJobCreditResponse {
	s.Body = v
	return s
}

func (s *GetYikeJobCreditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
