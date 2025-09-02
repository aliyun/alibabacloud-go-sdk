// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTriggerNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunTriggerNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunTriggerNodeResponse
	GetStatusCode() *int32
	SetBody(v *RunTriggerNodeResponseBody) *RunTriggerNodeResponse
	GetBody() *RunTriggerNodeResponseBody
}

type RunTriggerNodeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunTriggerNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTriggerNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RunTriggerNodeResponse) GoString() string {
	return s.String()
}

func (s *RunTriggerNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunTriggerNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTriggerNodeResponse) GetBody() *RunTriggerNodeResponseBody {
	return s.Body
}

func (s *RunTriggerNodeResponse) SetHeaders(v map[string]*string) *RunTriggerNodeResponse {
	s.Headers = v
	return s
}

func (s *RunTriggerNodeResponse) SetStatusCode(v int32) *RunTriggerNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTriggerNodeResponse) SetBody(v *RunTriggerNodeResponseBody) *RunTriggerNodeResponse {
	s.Body = v
	return s
}

func (s *RunTriggerNodeResponse) Validate() error {
	return dara.Validate(s)
}
