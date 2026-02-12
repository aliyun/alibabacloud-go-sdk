// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTrendGroupOutputTpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTrendGroupOutputTpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTrendGroupOutputTpsResponse
	GetStatusCode() *int32
	SetBody(v *OnsTrendGroupOutputTpsResponseBody) *OnsTrendGroupOutputTpsResponse
	GetBody() *OnsTrendGroupOutputTpsResponseBody
}

type OnsTrendGroupOutputTpsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTrendGroupOutputTpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTrendGroupOutputTpsResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendGroupOutputTpsResponse) GoString() string {
	return s.String()
}

func (s *OnsTrendGroupOutputTpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTrendGroupOutputTpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTrendGroupOutputTpsResponse) GetBody() *OnsTrendGroupOutputTpsResponseBody {
	return s.Body
}

func (s *OnsTrendGroupOutputTpsResponse) SetHeaders(v map[string]*string) *OnsTrendGroupOutputTpsResponse {
	s.Headers = v
	return s
}

func (s *OnsTrendGroupOutputTpsResponse) SetStatusCode(v int32) *OnsTrendGroupOutputTpsResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTrendGroupOutputTpsResponse) SetBody(v *OnsTrendGroupOutputTpsResponseBody) *OnsTrendGroupOutputTpsResponse {
	s.Body = v
	return s
}

func (s *OnsTrendGroupOutputTpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
