// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnMarkMonitorAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnMarkMonitorAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnMarkMonitorAccountsResponse
	GetStatusCode() *int32
	SetBody(v *UnMarkMonitorAccountsResponseBody) *UnMarkMonitorAccountsResponse
	GetBody() *UnMarkMonitorAccountsResponseBody
}

type UnMarkMonitorAccountsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnMarkMonitorAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnMarkMonitorAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s UnMarkMonitorAccountsResponse) GoString() string {
	return s.String()
}

func (s *UnMarkMonitorAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnMarkMonitorAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnMarkMonitorAccountsResponse) GetBody() *UnMarkMonitorAccountsResponseBody {
	return s.Body
}

func (s *UnMarkMonitorAccountsResponse) SetHeaders(v map[string]*string) *UnMarkMonitorAccountsResponse {
	s.Headers = v
	return s
}

func (s *UnMarkMonitorAccountsResponse) SetStatusCode(v int32) *UnMarkMonitorAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *UnMarkMonitorAccountsResponse) SetBody(v *UnMarkMonitorAccountsResponseBody) *UnMarkMonitorAccountsResponse {
	s.Body = v
	return s
}

func (s *UnMarkMonitorAccountsResponse) Validate() error {
	return dara.Validate(s)
}
