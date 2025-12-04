// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmartCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmartCallResponse
	GetStatusCode() *int32
	SetBody(v *SmartCallResponseBody) *SmartCallResponse
	GetBody() *SmartCallResponseBody
}

type SmartCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartCallResponse) String() string {
	return dara.Prettify(s)
}

func (s SmartCallResponse) GoString() string {
	return s.String()
}

func (s *SmartCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmartCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmartCallResponse) GetBody() *SmartCallResponseBody {
	return s.Body
}

func (s *SmartCallResponse) SetHeaders(v map[string]*string) *SmartCallResponse {
	s.Headers = v
	return s
}

func (s *SmartCallResponse) SetStatusCode(v int32) *SmartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartCallResponse) SetBody(v *SmartCallResponseBody) *SmartCallResponse {
	s.Body = v
	return s
}

func (s *SmartCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
