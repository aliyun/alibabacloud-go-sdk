// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawVpcPublishedRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WithdrawVpcPublishedRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WithdrawVpcPublishedRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *WithdrawVpcPublishedRouteEntriesResponseBody) *WithdrawVpcPublishedRouteEntriesResponse
	GetBody() *WithdrawVpcPublishedRouteEntriesResponseBody
}

type WithdrawVpcPublishedRouteEntriesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawVpcPublishedRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawVpcPublishedRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s WithdrawVpcPublishedRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *WithdrawVpcPublishedRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WithdrawVpcPublishedRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WithdrawVpcPublishedRouteEntriesResponse) GetBody() *WithdrawVpcPublishedRouteEntriesResponseBody {
	return s.Body
}

func (s *WithdrawVpcPublishedRouteEntriesResponse) SetHeaders(v map[string]*string) *WithdrawVpcPublishedRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesResponse) SetStatusCode(v int32) *WithdrawVpcPublishedRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesResponse) SetBody(v *WithdrawVpcPublishedRouteEntriesResponseBody) *WithdrawVpcPublishedRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *WithdrawVpcPublishedRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
