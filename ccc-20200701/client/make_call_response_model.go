// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MakeCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MakeCallResponse
	GetStatusCode() *int32
	SetBody(v *MakeCallResponseBody) *MakeCallResponse
	GetBody() *MakeCallResponseBody
}

type MakeCallResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MakeCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MakeCallResponse) String() string {
	return dara.Prettify(s)
}

func (s MakeCallResponse) GoString() string {
	return s.String()
}

func (s *MakeCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MakeCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MakeCallResponse) GetBody() *MakeCallResponseBody {
	return s.Body
}

func (s *MakeCallResponse) SetHeaders(v map[string]*string) *MakeCallResponse {
	s.Headers = v
	return s
}

func (s *MakeCallResponse) SetStatusCode(v int32) *MakeCallResponse {
	s.StatusCode = &v
	return s
}

func (s *MakeCallResponse) SetBody(v *MakeCallResponseBody) *MakeCallResponse {
	s.Body = v
	return s
}

func (s *MakeCallResponse) Validate() error {
	return dara.Validate(s)
}
