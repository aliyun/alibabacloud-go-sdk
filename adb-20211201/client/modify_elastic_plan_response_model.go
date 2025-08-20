// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElasticPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElasticPlanResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElasticPlanResponseBody) *ModifyElasticPlanResponse
	GetBody() *ModifyElasticPlanResponseBody
}

type ModifyElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElasticPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElasticPlanResponse) GetBody() *ModifyElasticPlanResponseBody {
	return s.Body
}

func (s *ModifyElasticPlanResponse) SetHeaders(v map[string]*string) *ModifyElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticPlanResponse) SetStatusCode(v int32) *ModifyElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticPlanResponse) SetBody(v *ModifyElasticPlanResponseBody) *ModifyElasticPlanResponse {
	s.Body = v
	return s
}

func (s *ModifyElasticPlanResponse) Validate() error {
	return dara.Validate(s)
}
