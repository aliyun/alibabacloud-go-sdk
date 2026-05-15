// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangUpDoubleCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HangUpDoubleCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HangUpDoubleCallResponse
	GetStatusCode() *int32
	SetBody(v *HangUpDoubleCallResponseBody) *HangUpDoubleCallResponse
	GetBody() *HangUpDoubleCallResponseBody
}

type HangUpDoubleCallResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HangUpDoubleCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HangUpDoubleCallResponse) String() string {
	return dara.Prettify(s)
}

func (s HangUpDoubleCallResponse) GoString() string {
	return s.String()
}

func (s *HangUpDoubleCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HangUpDoubleCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HangUpDoubleCallResponse) GetBody() *HangUpDoubleCallResponseBody {
	return s.Body
}

func (s *HangUpDoubleCallResponse) SetHeaders(v map[string]*string) *HangUpDoubleCallResponse {
	s.Headers = v
	return s
}

func (s *HangUpDoubleCallResponse) SetStatusCode(v int32) *HangUpDoubleCallResponse {
	s.StatusCode = &v
	return s
}

func (s *HangUpDoubleCallResponse) SetBody(v *HangUpDoubleCallResponseBody) *HangUpDoubleCallResponse {
	s.Body = v
	return s
}

func (s *HangUpDoubleCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
