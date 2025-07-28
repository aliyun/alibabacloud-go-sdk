// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemandPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDemandPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDemandPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateDemandPlanResponseBody) *CreateDemandPlanResponse
	GetBody() *CreateDemandPlanResponseBody
}

type CreateDemandPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDemandPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDemandPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDemandPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDemandPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDemandPlanResponse) GetBody() *CreateDemandPlanResponseBody {
	return s.Body
}

func (s *CreateDemandPlanResponse) SetHeaders(v map[string]*string) *CreateDemandPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDemandPlanResponse) SetStatusCode(v int32) *CreateDemandPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDemandPlanResponse) SetBody(v *CreateDemandPlanResponseBody) *CreateDemandPlanResponse {
	s.Body = v
	return s
}

func (s *CreateDemandPlanResponse) Validate() error {
	return dara.Validate(s)
}
