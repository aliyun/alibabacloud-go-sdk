// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableJobResponse
	GetStatusCode() *int32
	SetBody(v *DisableJobResponseBody) *DisableJobResponse
	GetBody() *DisableJobResponseBody
}

type DisableJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableJobResponse) GoString() string {
	return s.String()
}

func (s *DisableJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableJobResponse) GetBody() *DisableJobResponseBody {
	return s.Body
}

func (s *DisableJobResponse) SetHeaders(v map[string]*string) *DisableJobResponse {
	s.Headers = v
	return s
}

func (s *DisableJobResponse) SetStatusCode(v int32) *DisableJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableJobResponse) SetBody(v *DisableJobResponseBody) *DisableJobResponse {
	s.Body = v
	return s
}

func (s *DisableJobResponse) Validate() error {
	return dara.Validate(s)
}
