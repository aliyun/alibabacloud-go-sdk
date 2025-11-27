// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActionPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetActionPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetActionPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetActionPlanResponseBody) *GetActionPlanResponse
	GetBody() *GetActionPlanResponseBody
}

type GetActionPlanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetActionPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetActionPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetActionPlanResponse) GoString() string {
	return s.String()
}

func (s *GetActionPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetActionPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetActionPlanResponse) GetBody() *GetActionPlanResponseBody {
	return s.Body
}

func (s *GetActionPlanResponse) SetHeaders(v map[string]*string) *GetActionPlanResponse {
	s.Headers = v
	return s
}

func (s *GetActionPlanResponse) SetStatusCode(v int32) *GetActionPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetActionPlanResponse) SetBody(v *GetActionPlanResponseBody) *GetActionPlanResponse {
	s.Body = v
	return s
}

func (s *GetActionPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
