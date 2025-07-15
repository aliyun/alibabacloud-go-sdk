// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MuteCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MuteCallResponse
	GetStatusCode() *int32
	SetBody(v *MuteCallResponseBody) *MuteCallResponse
	GetBody() *MuteCallResponseBody
}

type MuteCallResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MuteCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MuteCallResponse) String() string {
	return dara.Prettify(s)
}

func (s MuteCallResponse) GoString() string {
	return s.String()
}

func (s *MuteCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MuteCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MuteCallResponse) GetBody() *MuteCallResponseBody {
	return s.Body
}

func (s *MuteCallResponse) SetHeaders(v map[string]*string) *MuteCallResponse {
	s.Headers = v
	return s
}

func (s *MuteCallResponse) SetStatusCode(v int32) *MuteCallResponse {
	s.StatusCode = &v
	return s
}

func (s *MuteCallResponse) SetBody(v *MuteCallResponseBody) *MuteCallResponse {
	s.Body = v
	return s
}

func (s *MuteCallResponse) Validate() error {
	return dara.Validate(s)
}
