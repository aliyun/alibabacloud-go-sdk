// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseSupportPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterpriseSupportPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterpriseSupportPlanResponse
	GetStatusCode() *int32
	SetBody(v *ListEnterpriseSupportPlanResponseBody) *ListEnterpriseSupportPlanResponse
	GetBody() *ListEnterpriseSupportPlanResponseBody
}

type ListEnterpriseSupportPlanResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseSupportPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseSupportPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterpriseSupportPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterpriseSupportPlanResponse) GetBody() *ListEnterpriseSupportPlanResponseBody {
	return s.Body
}

func (s *ListEnterpriseSupportPlanResponse) SetHeaders(v map[string]*string) *ListEnterpriseSupportPlanResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseSupportPlanResponse) SetStatusCode(v int32) *ListEnterpriseSupportPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanResponse) SetBody(v *ListEnterpriseSupportPlanResponseBody) *ListEnterpriseSupportPlanResponse {
	s.Body = v
	return s
}

func (s *ListEnterpriseSupportPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
