// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAllExpirationDayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAllExpirationDayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAllExpirationDayResponse
	GetStatusCode() *int32
	SetBody(v *SetAllExpirationDayResponseBody) *SetAllExpirationDayResponse
	GetBody() *SetAllExpirationDayResponseBody
}

type SetAllExpirationDayResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAllExpirationDayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAllExpirationDayResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAllExpirationDayResponse) GoString() string {
	return s.String()
}

func (s *SetAllExpirationDayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAllExpirationDayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAllExpirationDayResponse) GetBody() *SetAllExpirationDayResponseBody {
	return s.Body
}

func (s *SetAllExpirationDayResponse) SetHeaders(v map[string]*string) *SetAllExpirationDayResponse {
	s.Headers = v
	return s
}

func (s *SetAllExpirationDayResponse) SetStatusCode(v int32) *SetAllExpirationDayResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAllExpirationDayResponse) SetBody(v *SetAllExpirationDayResponseBody) *SetAllExpirationDayResponse {
	s.Body = v
	return s
}

func (s *SetAllExpirationDayResponse) Validate() error {
	return dara.Validate(s)
}
