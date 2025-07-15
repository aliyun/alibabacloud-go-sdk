// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPredictiveCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPredictiveCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPredictiveCallResponse
	GetStatusCode() *int32
	SetBody(v *StartPredictiveCallResponseBody) *StartPredictiveCallResponse
	GetBody() *StartPredictiveCallResponseBody
}

type StartPredictiveCallResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPredictiveCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPredictiveCallResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPredictiveCallResponse) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPredictiveCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPredictiveCallResponse) GetBody() *StartPredictiveCallResponseBody {
	return s.Body
}

func (s *StartPredictiveCallResponse) SetHeaders(v map[string]*string) *StartPredictiveCallResponse {
	s.Headers = v
	return s
}

func (s *StartPredictiveCallResponse) SetStatusCode(v int32) *StartPredictiveCallResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPredictiveCallResponse) SetBody(v *StartPredictiveCallResponseBody) *StartPredictiveCallResponse {
	s.Body = v
	return s
}

func (s *StartPredictiveCallResponse) Validate() error {
	return dara.Validate(s)
}
