// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupBroadcastTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetupBroadcastTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetupBroadcastTablesResponse
	GetStatusCode() *int32
	SetBody(v *SetupBroadcastTablesResponseBody) *SetupBroadcastTablesResponse
	GetBody() *SetupBroadcastTablesResponseBody
}

type SetupBroadcastTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetupBroadcastTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetupBroadcastTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetupBroadcastTablesResponse) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetupBroadcastTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetupBroadcastTablesResponse) GetBody() *SetupBroadcastTablesResponseBody {
	return s.Body
}

func (s *SetupBroadcastTablesResponse) SetHeaders(v map[string]*string) *SetupBroadcastTablesResponse {
	s.Headers = v
	return s
}

func (s *SetupBroadcastTablesResponse) SetStatusCode(v int32) *SetupBroadcastTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetupBroadcastTablesResponse) SetBody(v *SetupBroadcastTablesResponseBody) *SetupBroadcastTablesResponse {
	s.Body = v
	return s
}

func (s *SetupBroadcastTablesResponse) Validate() error {
	return dara.Validate(s)
}
