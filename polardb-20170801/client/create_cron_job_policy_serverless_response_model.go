// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCronJobPolicyServerlessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCronJobPolicyServerlessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCronJobPolicyServerlessResponse
	GetStatusCode() *int32
	SetBody(v *CreateCronJobPolicyServerlessResponseBody) *CreateCronJobPolicyServerlessResponse
	GetBody() *CreateCronJobPolicyServerlessResponseBody
}

type CreateCronJobPolicyServerlessResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCronJobPolicyServerlessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCronJobPolicyServerlessResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCronJobPolicyServerlessResponse) GoString() string {
	return s.String()
}

func (s *CreateCronJobPolicyServerlessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCronJobPolicyServerlessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCronJobPolicyServerlessResponse) GetBody() *CreateCronJobPolicyServerlessResponseBody {
	return s.Body
}

func (s *CreateCronJobPolicyServerlessResponse) SetHeaders(v map[string]*string) *CreateCronJobPolicyServerlessResponse {
	s.Headers = v
	return s
}

func (s *CreateCronJobPolicyServerlessResponse) SetStatusCode(v int32) *CreateCronJobPolicyServerlessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCronJobPolicyServerlessResponse) SetBody(v *CreateCronJobPolicyServerlessResponseBody) *CreateCronJobPolicyServerlessResponse {
	s.Body = v
	return s
}

func (s *CreateCronJobPolicyServerlessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
