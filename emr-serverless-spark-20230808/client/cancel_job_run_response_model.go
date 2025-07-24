// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelJobRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelJobRunResponse
	GetStatusCode() *int32
	SetBody(v *CancelJobRunResponseBody) *CancelJobRunResponse
	GetBody() *CancelJobRunResponseBody
}

type CancelJobRunResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelJobRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelJobRunResponse) GoString() string {
	return s.String()
}

func (s *CancelJobRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelJobRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelJobRunResponse) GetBody() *CancelJobRunResponseBody {
	return s.Body
}

func (s *CancelJobRunResponse) SetHeaders(v map[string]*string) *CancelJobRunResponse {
	s.Headers = v
	return s
}

func (s *CancelJobRunResponse) SetStatusCode(v int32) *CancelJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelJobRunResponse) SetBody(v *CancelJobRunResponseBody) *CancelJobRunResponse {
	s.Body = v
	return s
}

func (s *CancelJobRunResponse) Validate() error {
	return dara.Validate(s)
}
