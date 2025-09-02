// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateDISyncInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateDISyncInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateDISyncInstanceResponse
	GetStatusCode() *int32
	SetBody(v *TerminateDISyncInstanceResponseBody) *TerminateDISyncInstanceResponse
	GetBody() *TerminateDISyncInstanceResponseBody
}

type TerminateDISyncInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateDISyncInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateDISyncInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateDISyncInstanceResponse) GoString() string {
	return s.String()
}

func (s *TerminateDISyncInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateDISyncInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateDISyncInstanceResponse) GetBody() *TerminateDISyncInstanceResponseBody {
	return s.Body
}

func (s *TerminateDISyncInstanceResponse) SetHeaders(v map[string]*string) *TerminateDISyncInstanceResponse {
	s.Headers = v
	return s
}

func (s *TerminateDISyncInstanceResponse) SetStatusCode(v int32) *TerminateDISyncInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateDISyncInstanceResponse) SetBody(v *TerminateDISyncInstanceResponseBody) *TerminateDISyncInstanceResponse {
	s.Body = v
	return s
}

func (s *TerminateDISyncInstanceResponse) Validate() error {
	return dara.Validate(s)
}
