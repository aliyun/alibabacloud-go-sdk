// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCreditUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserCreditUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserCreditUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetUserCreditUsageResponseBody) *GetUserCreditUsageResponse
	GetBody() *GetUserCreditUsageResponseBody
}

type GetUserCreditUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserCreditUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserCreditUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserCreditUsageResponse) GoString() string {
	return s.String()
}

func (s *GetUserCreditUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserCreditUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserCreditUsageResponse) GetBody() *GetUserCreditUsageResponseBody {
	return s.Body
}

func (s *GetUserCreditUsageResponse) SetHeaders(v map[string]*string) *GetUserCreditUsageResponse {
	s.Headers = v
	return s
}

func (s *GetUserCreditUsageResponse) SetStatusCode(v int32) *GetUserCreditUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserCreditUsageResponse) SetBody(v *GetUserCreditUsageResponseBody) *GetUserCreditUsageResponse {
	s.Body = v
	return s
}

func (s *GetUserCreditUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
