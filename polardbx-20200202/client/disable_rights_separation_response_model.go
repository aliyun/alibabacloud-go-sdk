// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRightsSeparationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableRightsSeparationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableRightsSeparationResponse
	GetStatusCode() *int32
	SetBody(v *DisableRightsSeparationResponseBody) *DisableRightsSeparationResponse
	GetBody() *DisableRightsSeparationResponseBody
}

type DisableRightsSeparationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableRightsSeparationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableRightsSeparationResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableRightsSeparationResponse) GoString() string {
	return s.String()
}

func (s *DisableRightsSeparationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableRightsSeparationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableRightsSeparationResponse) GetBody() *DisableRightsSeparationResponseBody {
	return s.Body
}

func (s *DisableRightsSeparationResponse) SetHeaders(v map[string]*string) *DisableRightsSeparationResponse {
	s.Headers = v
	return s
}

func (s *DisableRightsSeparationResponse) SetStatusCode(v int32) *DisableRightsSeparationResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableRightsSeparationResponse) SetBody(v *DisableRightsSeparationResponseBody) *DisableRightsSeparationResponse {
	s.Body = v
	return s
}

func (s *DisableRightsSeparationResponse) Validate() error {
	return dara.Validate(s)
}
