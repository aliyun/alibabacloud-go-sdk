// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkMonitorAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MarkMonitorAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MarkMonitorAccountsResponse
	GetStatusCode() *int32
	SetBody(v *MarkMonitorAccountsResponseBody) *MarkMonitorAccountsResponse
	GetBody() *MarkMonitorAccountsResponseBody
}

type MarkMonitorAccountsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkMonitorAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MarkMonitorAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s MarkMonitorAccountsResponse) GoString() string {
	return s.String()
}

func (s *MarkMonitorAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MarkMonitorAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MarkMonitorAccountsResponse) GetBody() *MarkMonitorAccountsResponseBody {
	return s.Body
}

func (s *MarkMonitorAccountsResponse) SetHeaders(v map[string]*string) *MarkMonitorAccountsResponse {
	s.Headers = v
	return s
}

func (s *MarkMonitorAccountsResponse) SetStatusCode(v int32) *MarkMonitorAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkMonitorAccountsResponse) SetBody(v *MarkMonitorAccountsResponseBody) *MarkMonitorAccountsResponse {
	s.Body = v
	return s
}

func (s *MarkMonitorAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
