// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActionPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateActionPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateActionPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateActionPlanResponseBody) *CreateActionPlanResponse
	GetBody() *CreateActionPlanResponseBody
}

type CreateActionPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateActionPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateActionPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateActionPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateActionPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateActionPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateActionPlanResponse) GetBody() *CreateActionPlanResponseBody {
	return s.Body
}

func (s *CreateActionPlanResponse) SetHeaders(v map[string]*string) *CreateActionPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateActionPlanResponse) SetStatusCode(v int32) *CreateActionPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateActionPlanResponse) SetBody(v *CreateActionPlanResponseBody) *CreateActionPlanResponse {
	s.Body = v
	return s
}

func (s *CreateActionPlanResponse) Validate() error {
	return dara.Validate(s)
}
