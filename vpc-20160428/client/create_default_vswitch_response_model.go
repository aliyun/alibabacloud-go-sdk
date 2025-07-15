// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefaultVSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDefaultVSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDefaultVSwitchResponse
	GetStatusCode() *int32
	SetBody(v *CreateDefaultVSwitchResponseBody) *CreateDefaultVSwitchResponse
	GetBody() *CreateDefaultVSwitchResponseBody
}

type CreateDefaultVSwitchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefaultVSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefaultVSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDefaultVSwitchResponse) GoString() string {
	return s.String()
}

func (s *CreateDefaultVSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDefaultVSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDefaultVSwitchResponse) GetBody() *CreateDefaultVSwitchResponseBody {
	return s.Body
}

func (s *CreateDefaultVSwitchResponse) SetHeaders(v map[string]*string) *CreateDefaultVSwitchResponse {
	s.Headers = v
	return s
}

func (s *CreateDefaultVSwitchResponse) SetStatusCode(v int32) *CreateDefaultVSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefaultVSwitchResponse) SetBody(v *CreateDefaultVSwitchResponseBody) *CreateDefaultVSwitchResponse {
	s.Body = v
	return s
}

func (s *CreateDefaultVSwitchResponse) Validate() error {
	return dara.Validate(s)
}
