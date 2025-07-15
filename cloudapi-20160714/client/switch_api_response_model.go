// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchApiResponse
	GetStatusCode() *int32
	SetBody(v *SwitchApiResponseBody) *SwitchApiResponse
	GetBody() *SwitchApiResponseBody
}

type SwitchApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchApiResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchApiResponse) GoString() string {
	return s.String()
}

func (s *SwitchApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchApiResponse) GetBody() *SwitchApiResponseBody {
	return s.Body
}

func (s *SwitchApiResponse) SetHeaders(v map[string]*string) *SwitchApiResponse {
	s.Headers = v
	return s
}

func (s *SwitchApiResponse) SetStatusCode(v int32) *SwitchApiResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchApiResponse) SetBody(v *SwitchApiResponseBody) *SwitchApiResponse {
	s.Body = v
	return s
}

func (s *SwitchApiResponse) Validate() error {
	return dara.Validate(s)
}
