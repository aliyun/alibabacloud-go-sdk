// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCronJobPolicyServerlessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCronJobPolicyServerlessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCronJobPolicyServerlessResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCronJobPolicyServerlessResponseBody) *DescribeCronJobPolicyServerlessResponse
	GetBody() *DescribeCronJobPolicyServerlessResponseBody
}

type DescribeCronJobPolicyServerlessResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCronJobPolicyServerlessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCronJobPolicyServerlessResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCronJobPolicyServerlessResponse) GoString() string {
	return s.String()
}

func (s *DescribeCronJobPolicyServerlessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCronJobPolicyServerlessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCronJobPolicyServerlessResponse) GetBody() *DescribeCronJobPolicyServerlessResponseBody {
	return s.Body
}

func (s *DescribeCronJobPolicyServerlessResponse) SetHeaders(v map[string]*string) *DescribeCronJobPolicyServerlessResponse {
	s.Headers = v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponse) SetStatusCode(v int32) *DescribeCronJobPolicyServerlessResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponse) SetBody(v *DescribeCronJobPolicyServerlessResponseBody) *DescribeCronJobPolicyServerlessResponse {
	s.Body = v
	return s
}

func (s *DescribeCronJobPolicyServerlessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
