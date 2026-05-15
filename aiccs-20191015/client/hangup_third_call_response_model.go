// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupThirdCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HangupThirdCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HangupThirdCallResponse
	GetStatusCode() *int32
	SetBody(v *HangupThirdCallResponseBody) *HangupThirdCallResponse
	GetBody() *HangupThirdCallResponseBody
}

type HangupThirdCallResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HangupThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HangupThirdCallResponse) String() string {
	return dara.Prettify(s)
}

func (s HangupThirdCallResponse) GoString() string {
	return s.String()
}

func (s *HangupThirdCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HangupThirdCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HangupThirdCallResponse) GetBody() *HangupThirdCallResponseBody {
	return s.Body
}

func (s *HangupThirdCallResponse) SetHeaders(v map[string]*string) *HangupThirdCallResponse {
	s.Headers = v
	return s
}

func (s *HangupThirdCallResponse) SetStatusCode(v int32) *HangupThirdCallResponse {
	s.StatusCode = &v
	return s
}

func (s *HangupThirdCallResponse) SetBody(v *HangupThirdCallResponseBody) *HangupThirdCallResponse {
	s.Body = v
	return s
}

func (s *HangupThirdCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
