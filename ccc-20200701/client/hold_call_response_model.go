// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHoldCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HoldCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HoldCallResponse
	GetStatusCode() *int32
	SetBody(v *HoldCallResponseBody) *HoldCallResponse
	GetBody() *HoldCallResponseBody
}

type HoldCallResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HoldCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HoldCallResponse) String() string {
	return dara.Prettify(s)
}

func (s HoldCallResponse) GoString() string {
	return s.String()
}

func (s *HoldCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HoldCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HoldCallResponse) GetBody() *HoldCallResponseBody {
	return s.Body
}

func (s *HoldCallResponse) SetHeaders(v map[string]*string) *HoldCallResponse {
	s.Headers = v
	return s
}

func (s *HoldCallResponse) SetStatusCode(v int32) *HoldCallResponse {
	s.StatusCode = &v
	return s
}

func (s *HoldCallResponse) SetBody(v *HoldCallResponseBody) *HoldCallResponse {
	s.Body = v
	return s
}

func (s *HoldCallResponse) Validate() error {
	return dara.Validate(s)
}
