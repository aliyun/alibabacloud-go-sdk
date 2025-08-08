// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPushSchedulerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPushSchedulerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPushSchedulerResponse
	GetStatusCode() *int32
	SetBody(v *CancelPushSchedulerResponseBody) *CancelPushSchedulerResponse
	GetBody() *CancelPushSchedulerResponseBody
}

type CancelPushSchedulerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPushSchedulerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPushSchedulerResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPushSchedulerResponse) GoString() string {
	return s.String()
}

func (s *CancelPushSchedulerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPushSchedulerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPushSchedulerResponse) GetBody() *CancelPushSchedulerResponseBody {
	return s.Body
}

func (s *CancelPushSchedulerResponse) SetHeaders(v map[string]*string) *CancelPushSchedulerResponse {
	s.Headers = v
	return s
}

func (s *CancelPushSchedulerResponse) SetStatusCode(v int32) *CancelPushSchedulerResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPushSchedulerResponse) SetBody(v *CancelPushSchedulerResponseBody) *CancelPushSchedulerResponse {
	s.Body = v
	return s
}

func (s *CancelPushSchedulerResponse) Validate() error {
	return dara.Validate(s)
}
