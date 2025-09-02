// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDISyncInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDISyncInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDISyncInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartDISyncInstanceResponseBody) *StartDISyncInstanceResponse
	GetBody() *StartDISyncInstanceResponseBody
}

type StartDISyncInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDISyncInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDISyncInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDISyncInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartDISyncInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDISyncInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDISyncInstanceResponse) GetBody() *StartDISyncInstanceResponseBody {
	return s.Body
}

func (s *StartDISyncInstanceResponse) SetHeaders(v map[string]*string) *StartDISyncInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartDISyncInstanceResponse) SetStatusCode(v int32) *StartDISyncInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDISyncInstanceResponse) SetBody(v *StartDISyncInstanceResponseBody) *StartDISyncInstanceResponse {
	s.Body = v
	return s
}

func (s *StartDISyncInstanceResponse) Validate() error {
	return dara.Validate(s)
}
