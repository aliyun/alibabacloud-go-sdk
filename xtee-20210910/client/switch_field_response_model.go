// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchFieldResponse
	GetStatusCode() *int32
	SetBody(v *SwitchFieldResponseBody) *SwitchFieldResponse
	GetBody() *SwitchFieldResponseBody
}

type SwitchFieldResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchFieldResponse) GoString() string {
	return s.String()
}

func (s *SwitchFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchFieldResponse) GetBody() *SwitchFieldResponseBody {
	return s.Body
}

func (s *SwitchFieldResponse) SetHeaders(v map[string]*string) *SwitchFieldResponse {
	s.Headers = v
	return s
}

func (s *SwitchFieldResponse) SetStatusCode(v int32) *SwitchFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchFieldResponse) SetBody(v *SwitchFieldResponseBody) *SwitchFieldResponse {
	s.Body = v
	return s
}

func (s *SwitchFieldResponse) Validate() error {
	return dara.Validate(s)
}
