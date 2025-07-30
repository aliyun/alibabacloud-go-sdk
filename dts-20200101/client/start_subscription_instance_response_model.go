// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSubscriptionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSubscriptionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSubscriptionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartSubscriptionInstanceResponseBody) *StartSubscriptionInstanceResponse
	GetBody() *StartSubscriptionInstanceResponseBody
}

type StartSubscriptionInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSubscriptionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSubscriptionInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSubscriptionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSubscriptionInstanceResponse) GetBody() *StartSubscriptionInstanceResponseBody {
	return s.Body
}

func (s *StartSubscriptionInstanceResponse) SetHeaders(v map[string]*string) *StartSubscriptionInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartSubscriptionInstanceResponse) SetStatusCode(v int32) *StartSubscriptionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSubscriptionInstanceResponse) SetBody(v *StartSubscriptionInstanceResponseBody) *StartSubscriptionInstanceResponse {
	s.Body = v
	return s
}

func (s *StartSubscriptionInstanceResponse) Validate() error {
	return dara.Validate(s)
}
