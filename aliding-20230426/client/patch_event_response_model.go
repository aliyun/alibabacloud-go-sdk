// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PatchEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PatchEventResponse
	GetStatusCode() *int32
	SetBody(v *PatchEventResponseBody) *PatchEventResponse
	GetBody() *PatchEventResponseBody
}

type PatchEventResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PatchEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PatchEventResponse) String() string {
	return dara.Prettify(s)
}

func (s PatchEventResponse) GoString() string {
	return s.String()
}

func (s *PatchEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PatchEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PatchEventResponse) GetBody() *PatchEventResponseBody {
	return s.Body
}

func (s *PatchEventResponse) SetHeaders(v map[string]*string) *PatchEventResponse {
	s.Headers = v
	return s
}

func (s *PatchEventResponse) SetStatusCode(v int32) *PatchEventResponse {
	s.StatusCode = &v
	return s
}

func (s *PatchEventResponse) SetBody(v *PatchEventResponseBody) *PatchEventResponse {
	s.Body = v
	return s
}

func (s *PatchEventResponse) Validate() error {
	return dara.Validate(s)
}
