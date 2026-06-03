// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonBillingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCommonBillingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCommonBillingResponse
	GetStatusCode() *int32
	SetBody(v *QueryCommonBillingResponseBody) *QueryCommonBillingResponse
	GetBody() *QueryCommonBillingResponseBody
}

type QueryCommonBillingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCommonBillingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCommonBillingResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonBillingResponse) GoString() string {
	return s.String()
}

func (s *QueryCommonBillingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCommonBillingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCommonBillingResponse) GetBody() *QueryCommonBillingResponseBody {
	return s.Body
}

func (s *QueryCommonBillingResponse) SetHeaders(v map[string]*string) *QueryCommonBillingResponse {
	s.Headers = v
	return s
}

func (s *QueryCommonBillingResponse) SetStatusCode(v int32) *QueryCommonBillingResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCommonBillingResponse) SetBody(v *QueryCommonBillingResponseBody) *QueryCommonBillingResponse {
	s.Body = v
	return s
}

func (s *QueryCommonBillingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
