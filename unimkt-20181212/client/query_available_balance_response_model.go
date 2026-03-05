// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableBalanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAvailableBalanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAvailableBalanceResponse
	GetStatusCode() *int32
	SetBody(v *QueryAvailableBalanceResponseBody) *QueryAvailableBalanceResponse
	GetBody() *QueryAvailableBalanceResponseBody
}

type QueryAvailableBalanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvailableBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvailableBalanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableBalanceResponse) GoString() string {
	return s.String()
}

func (s *QueryAvailableBalanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAvailableBalanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAvailableBalanceResponse) GetBody() *QueryAvailableBalanceResponseBody {
	return s.Body
}

func (s *QueryAvailableBalanceResponse) SetHeaders(v map[string]*string) *QueryAvailableBalanceResponse {
	s.Headers = v
	return s
}

func (s *QueryAvailableBalanceResponse) SetStatusCode(v int32) *QueryAvailableBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvailableBalanceResponse) SetBody(v *QueryAvailableBalanceResponseBody) *QueryAvailableBalanceResponse {
	s.Body = v
	return s
}

func (s *QueryAvailableBalanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
