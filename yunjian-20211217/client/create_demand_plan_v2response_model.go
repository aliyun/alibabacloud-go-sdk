// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemandPlanV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDemandPlanV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDemandPlanV2Response
	GetStatusCode() *int32
	SetBody(v *CreateDemandPlanV2ResponseBody) *CreateDemandPlanV2Response
	GetBody() *CreateDemandPlanV2ResponseBody
}

type CreateDemandPlanV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDemandPlanV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDemandPlanV2Response) String() string {
	return dara.Prettify(s)
}

func (s CreateDemandPlanV2Response) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDemandPlanV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDemandPlanV2Response) GetBody() *CreateDemandPlanV2ResponseBody {
	return s.Body
}

func (s *CreateDemandPlanV2Response) SetHeaders(v map[string]*string) *CreateDemandPlanV2Response {
	s.Headers = v
	return s
}

func (s *CreateDemandPlanV2Response) SetStatusCode(v int32) *CreateDemandPlanV2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateDemandPlanV2Response) SetBody(v *CreateDemandPlanV2ResponseBody) *CreateDemandPlanV2Response {
	s.Body = v
	return s
}

func (s *CreateDemandPlanV2Response) Validate() error {
	return dara.Validate(s)
}
