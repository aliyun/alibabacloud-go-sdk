// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HangupCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HangupCallResponse
	GetStatusCode() *int32
	SetBody(v *HangupCallResponseBody) *HangupCallResponse
	GetBody() *HangupCallResponseBody
}

type HangupCallResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HangupCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HangupCallResponse) String() string {
	return dara.Prettify(s)
}

func (s HangupCallResponse) GoString() string {
	return s.String()
}

func (s *HangupCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HangupCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HangupCallResponse) GetBody() *HangupCallResponseBody {
	return s.Body
}

func (s *HangupCallResponse) SetHeaders(v map[string]*string) *HangupCallResponse {
	s.Headers = v
	return s
}

func (s *HangupCallResponse) SetStatusCode(v int32) *HangupCallResponse {
	s.StatusCode = &v
	return s
}

func (s *HangupCallResponse) SetBody(v *HangupCallResponseBody) *HangupCallResponse {
	s.Body = v
	return s
}

func (s *HangupCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
