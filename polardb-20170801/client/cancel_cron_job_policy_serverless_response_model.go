// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCronJobPolicyServerlessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCronJobPolicyServerlessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCronJobPolicyServerlessResponse
	GetStatusCode() *int32
	SetBody(v *CancelCronJobPolicyServerlessResponseBody) *CancelCronJobPolicyServerlessResponse
	GetBody() *CancelCronJobPolicyServerlessResponseBody
}

type CancelCronJobPolicyServerlessResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCronJobPolicyServerlessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCronJobPolicyServerlessResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCronJobPolicyServerlessResponse) GoString() string {
	return s.String()
}

func (s *CancelCronJobPolicyServerlessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCronJobPolicyServerlessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCronJobPolicyServerlessResponse) GetBody() *CancelCronJobPolicyServerlessResponseBody {
	return s.Body
}

func (s *CancelCronJobPolicyServerlessResponse) SetHeaders(v map[string]*string) *CancelCronJobPolicyServerlessResponse {
	s.Headers = v
	return s
}

func (s *CancelCronJobPolicyServerlessResponse) SetStatusCode(v int32) *CancelCronJobPolicyServerlessResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCronJobPolicyServerlessResponse) SetBody(v *CancelCronJobPolicyServerlessResponseBody) *CancelCronJobPolicyServerlessResponse {
	s.Body = v
	return s
}

func (s *CancelCronJobPolicyServerlessResponse) Validate() error {
	return dara.Validate(s)
}
