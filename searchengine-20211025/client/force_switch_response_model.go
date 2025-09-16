// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForceSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForceSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForceSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ForceSwitchResponseBody) *ForceSwitchResponse
	GetBody() *ForceSwitchResponseBody
}

type ForceSwitchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForceSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForceSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ForceSwitchResponse) GoString() string {
	return s.String()
}

func (s *ForceSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForceSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForceSwitchResponse) GetBody() *ForceSwitchResponseBody {
	return s.Body
}

func (s *ForceSwitchResponse) SetHeaders(v map[string]*string) *ForceSwitchResponse {
	s.Headers = v
	return s
}

func (s *ForceSwitchResponse) SetStatusCode(v int32) *ForceSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ForceSwitchResponse) SetBody(v *ForceSwitchResponseBody) *ForceSwitchResponse {
	s.Body = v
	return s
}

func (s *ForceSwitchResponse) Validate() error {
	return dara.Validate(s)
}
