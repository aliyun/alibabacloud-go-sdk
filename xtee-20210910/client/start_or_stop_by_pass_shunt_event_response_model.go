// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartOrStopByPassShuntEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartOrStopByPassShuntEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartOrStopByPassShuntEventResponse
	GetStatusCode() *int32
	SetBody(v *StartOrStopByPassShuntEventResponseBody) *StartOrStopByPassShuntEventResponse
	GetBody() *StartOrStopByPassShuntEventResponseBody
}

type StartOrStopByPassShuntEventResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartOrStopByPassShuntEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartOrStopByPassShuntEventResponse) String() string {
	return dara.Prettify(s)
}

func (s StartOrStopByPassShuntEventResponse) GoString() string {
	return s.String()
}

func (s *StartOrStopByPassShuntEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartOrStopByPassShuntEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartOrStopByPassShuntEventResponse) GetBody() *StartOrStopByPassShuntEventResponseBody {
	return s.Body
}

func (s *StartOrStopByPassShuntEventResponse) SetHeaders(v map[string]*string) *StartOrStopByPassShuntEventResponse {
	s.Headers = v
	return s
}

func (s *StartOrStopByPassShuntEventResponse) SetStatusCode(v int32) *StartOrStopByPassShuntEventResponse {
	s.StatusCode = &v
	return s
}

func (s *StartOrStopByPassShuntEventResponse) SetBody(v *StartOrStopByPassShuntEventResponseBody) *StartOrStopByPassShuntEventResponse {
	s.Body = v
	return s
}

func (s *StartOrStopByPassShuntEventResponse) Validate() error {
	return dara.Validate(s)
}
