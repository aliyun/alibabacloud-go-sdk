// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFCTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFCTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFCTriggerResponse
	GetStatusCode() *int32
	SetBody(v *AddFCTriggerResponseBody) *AddFCTriggerResponse
	GetBody() *AddFCTriggerResponseBody
}

type AddFCTriggerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFCTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *AddFCTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFCTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFCTriggerResponse) GetBody() *AddFCTriggerResponseBody {
	return s.Body
}

func (s *AddFCTriggerResponse) SetHeaders(v map[string]*string) *AddFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *AddFCTriggerResponse) SetStatusCode(v int32) *AddFCTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFCTriggerResponse) SetBody(v *AddFCTriggerResponseBody) *AddFCTriggerResponse {
	s.Body = v
	return s
}

func (s *AddFCTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
