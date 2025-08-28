// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIvrCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IvrCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IvrCallResponse
	GetStatusCode() *int32
	SetBody(v *IvrCallResponseBody) *IvrCallResponse
	GetBody() *IvrCallResponseBody
}

type IvrCallResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IvrCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IvrCallResponse) String() string {
	return dara.Prettify(s)
}

func (s IvrCallResponse) GoString() string {
	return s.String()
}

func (s *IvrCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IvrCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IvrCallResponse) GetBody() *IvrCallResponseBody {
	return s.Body
}

func (s *IvrCallResponse) SetHeaders(v map[string]*string) *IvrCallResponse {
	s.Headers = v
	return s
}

func (s *IvrCallResponse) SetStatusCode(v int32) *IvrCallResponse {
	s.StatusCode = &v
	return s
}

func (s *IvrCallResponse) SetBody(v *IvrCallResponseBody) *IvrCallResponse {
	s.Body = v
	return s
}

func (s *IvrCallResponse) Validate() error {
	return dara.Validate(s)
}
