// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMonitorAccountResponseBody) *DeleteMonitorAccountResponse
	GetBody() *DeleteMonitorAccountResponseBody
}

type DeleteMonitorAccountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMonitorAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMonitorAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorAccountResponse) GetBody() *DeleteMonitorAccountResponseBody {
	return s.Body
}

func (s *DeleteMonitorAccountResponse) SetHeaders(v map[string]*string) *DeleteMonitorAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorAccountResponse) SetStatusCode(v int32) *DeleteMonitorAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorAccountResponse) SetBody(v *DeleteMonitorAccountResponseBody) *DeleteMonitorAccountResponse {
	s.Body = v
	return s
}

func (s *DeleteMonitorAccountResponse) Validate() error {
	return dara.Validate(s)
}
