// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviousAndNextControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviousAndNextControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviousAndNextControlResponse
	GetStatusCode() *int32
	SetBody(v *PreviousAndNextControlResponseBody) *PreviousAndNextControlResponse
	GetBody() *PreviousAndNextControlResponseBody
}

type PreviousAndNextControlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviousAndNextControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviousAndNextControlResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviousAndNextControlResponse) GoString() string {
	return s.String()
}

func (s *PreviousAndNextControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviousAndNextControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviousAndNextControlResponse) GetBody() *PreviousAndNextControlResponseBody {
	return s.Body
}

func (s *PreviousAndNextControlResponse) SetHeaders(v map[string]*string) *PreviousAndNextControlResponse {
	s.Headers = v
	return s
}

func (s *PreviousAndNextControlResponse) SetStatusCode(v int32) *PreviousAndNextControlResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviousAndNextControlResponse) SetBody(v *PreviousAndNextControlResponseBody) *PreviousAndNextControlResponse {
	s.Body = v
	return s
}

func (s *PreviousAndNextControlResponse) Validate() error {
	return dara.Validate(s)
}
