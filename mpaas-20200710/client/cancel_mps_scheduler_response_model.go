// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMpsSchedulerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelMpsSchedulerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelMpsSchedulerResponse
	GetStatusCode() *int32
	SetBody(v *CancelMpsSchedulerResponseBody) *CancelMpsSchedulerResponse
	GetBody() *CancelMpsSchedulerResponseBody
}

type CancelMpsSchedulerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelMpsSchedulerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelMpsSchedulerResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelMpsSchedulerResponse) GoString() string {
	return s.String()
}

func (s *CancelMpsSchedulerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelMpsSchedulerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelMpsSchedulerResponse) GetBody() *CancelMpsSchedulerResponseBody {
	return s.Body
}

func (s *CancelMpsSchedulerResponse) SetHeaders(v map[string]*string) *CancelMpsSchedulerResponse {
	s.Headers = v
	return s
}

func (s *CancelMpsSchedulerResponse) SetStatusCode(v int32) *CancelMpsSchedulerResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelMpsSchedulerResponse) SetBody(v *CancelMpsSchedulerResponseBody) *CancelMpsSchedulerResponse {
	s.Body = v
	return s
}

func (s *CancelMpsSchedulerResponse) Validate() error {
	return dara.Validate(s)
}
