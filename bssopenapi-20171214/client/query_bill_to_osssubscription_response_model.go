// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBillToOSSSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBillToOSSSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBillToOSSSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *QueryBillToOSSSubscriptionResponseBody) *QueryBillToOSSSubscriptionResponse
	GetBody() *QueryBillToOSSSubscriptionResponseBody
}

type QueryBillToOSSSubscriptionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBillToOSSSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBillToOSSSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBillToOSSSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBillToOSSSubscriptionResponse) GetBody() *QueryBillToOSSSubscriptionResponseBody {
	return s.Body
}

func (s *QueryBillToOSSSubscriptionResponse) SetHeaders(v map[string]*string) *QueryBillToOSSSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *QueryBillToOSSSubscriptionResponse) SetStatusCode(v int32) *QueryBillToOSSSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponse) SetBody(v *QueryBillToOSSSubscriptionResponseBody) *QueryBillToOSSSubscriptionResponse {
	s.Body = v
	return s
}

func (s *QueryBillToOSSSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
