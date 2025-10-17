// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCronJobPolicyServerlessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCronJobPolicyServerlessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCronJobPolicyServerlessResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCronJobPolicyServerlessResponseBody) *ModifyCronJobPolicyServerlessResponse
	GetBody() *ModifyCronJobPolicyServerlessResponseBody
}

type ModifyCronJobPolicyServerlessResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCronJobPolicyServerlessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCronJobPolicyServerlessResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCronJobPolicyServerlessResponse) GoString() string {
	return s.String()
}

func (s *ModifyCronJobPolicyServerlessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCronJobPolicyServerlessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCronJobPolicyServerlessResponse) GetBody() *ModifyCronJobPolicyServerlessResponseBody {
	return s.Body
}

func (s *ModifyCronJobPolicyServerlessResponse) SetHeaders(v map[string]*string) *ModifyCronJobPolicyServerlessResponse {
	s.Headers = v
	return s
}

func (s *ModifyCronJobPolicyServerlessResponse) SetStatusCode(v int32) *ModifyCronJobPolicyServerlessResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessResponse) SetBody(v *ModifyCronJobPolicyServerlessResponseBody) *ModifyCronJobPolicyServerlessResponse {
	s.Body = v
	return s
}

func (s *ModifyCronJobPolicyServerlessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
