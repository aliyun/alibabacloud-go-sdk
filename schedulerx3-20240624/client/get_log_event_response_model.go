// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogEventResponse
	GetStatusCode() *int32
	SetBody(v *GetLogEventResponseBody) *GetLogEventResponse
	GetBody() *GetLogEventResponseBody
}

type GetLogEventResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogEventResponse) GoString() string {
	return s.String()
}

func (s *GetLogEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogEventResponse) GetBody() *GetLogEventResponseBody {
	return s.Body
}

func (s *GetLogEventResponse) SetHeaders(v map[string]*string) *GetLogEventResponse {
	s.Headers = v
	return s
}

func (s *GetLogEventResponse) SetStatusCode(v int32) *GetLogEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogEventResponse) SetBody(v *GetLogEventResponseBody) *GetLogEventResponse {
	s.Body = v
	return s
}

func (s *GetLogEventResponse) Validate() error {
	return dara.Validate(s)
}
