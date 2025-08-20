// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteElasticPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteElasticPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteElasticPlanResponseBody) *DeleteElasticPlanResponse
	GetBody() *DeleteElasticPlanResponseBody
}

type DeleteElasticPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteElasticPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteElasticPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteElasticPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteElasticPlanResponse) GetBody() *DeleteElasticPlanResponseBody {
	return s.Body
}

func (s *DeleteElasticPlanResponse) SetHeaders(v map[string]*string) *DeleteElasticPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteElasticPlanResponse) SetStatusCode(v int32) *DeleteElasticPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteElasticPlanResponse) SetBody(v *DeleteElasticPlanResponseBody) *DeleteElasticPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteElasticPlanResponse) Validate() error {
	return dara.Validate(s)
}
