// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawPublishedRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WithdrawPublishedRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WithdrawPublishedRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *WithdrawPublishedRouteEntriesResponseBody) *WithdrawPublishedRouteEntriesResponse
	GetBody() *WithdrawPublishedRouteEntriesResponseBody
}

type WithdrawPublishedRouteEntriesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawPublishedRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawPublishedRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s WithdrawPublishedRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *WithdrawPublishedRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WithdrawPublishedRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WithdrawPublishedRouteEntriesResponse) GetBody() *WithdrawPublishedRouteEntriesResponseBody {
	return s.Body
}

func (s *WithdrawPublishedRouteEntriesResponse) SetHeaders(v map[string]*string) *WithdrawPublishedRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *WithdrawPublishedRouteEntriesResponse) SetStatusCode(v int32) *WithdrawPublishedRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawPublishedRouteEntriesResponse) SetBody(v *WithdrawPublishedRouteEntriesResponseBody) *WithdrawPublishedRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *WithdrawPublishedRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
