// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartApsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartApsJobResponse
	GetStatusCode() *int32
	SetBody(v *StartApsJobResponseBody) *StartApsJobResponse
	GetBody() *StartApsJobResponseBody
}

type StartApsJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartApsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartApsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartApsJobResponse) GoString() string {
	return s.String()
}

func (s *StartApsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartApsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartApsJobResponse) GetBody() *StartApsJobResponseBody {
	return s.Body
}

func (s *StartApsJobResponse) SetHeaders(v map[string]*string) *StartApsJobResponse {
	s.Headers = v
	return s
}

func (s *StartApsJobResponse) SetStatusCode(v int32) *StartApsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartApsJobResponse) SetBody(v *StartApsJobResponseBody) *StartApsJobResponse {
	s.Body = v
	return s
}

func (s *StartApsJobResponse) Validate() error {
	return dara.Validate(s)
}
