// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveElasticPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveElasticPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveElasticPlanResponse
	GetStatusCode() *int32
	SetBody(v *RemoveElasticPlanResponseBody) *RemoveElasticPlanResponse
	GetBody() *RemoveElasticPlanResponseBody
}

type RemoveElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveElasticPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *RemoveElasticPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveElasticPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveElasticPlanResponse) GetBody() *RemoveElasticPlanResponseBody {
	return s.Body
}

func (s *RemoveElasticPlanResponse) SetHeaders(v map[string]*string) *RemoveElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *RemoveElasticPlanResponse) SetStatusCode(v int32) *RemoveElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveElasticPlanResponse) SetBody(v *RemoveElasticPlanResponseBody) *RemoveElasticPlanResponse {
	s.Body = v
	return s
}

func (s *RemoveElasticPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
