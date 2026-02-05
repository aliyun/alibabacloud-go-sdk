// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseSupportPlanSimpleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnterpriseSupportPlanSimpleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnterpriseSupportPlanSimpleResponse
	GetStatusCode() *int32
	SetBody(v *ListEnterpriseSupportPlanSimpleResponseBody) *ListEnterpriseSupportPlanSimpleResponse
	GetBody() *ListEnterpriseSupportPlanSimpleResponseBody
}

type ListEnterpriseSupportPlanSimpleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnterpriseSupportPlanSimpleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnterpriseSupportPlanSimpleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseSupportPlanSimpleResponse) GoString() string {
	return s.String()
}

func (s *ListEnterpriseSupportPlanSimpleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnterpriseSupportPlanSimpleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnterpriseSupportPlanSimpleResponse) GetBody() *ListEnterpriseSupportPlanSimpleResponseBody {
	return s.Body
}

func (s *ListEnterpriseSupportPlanSimpleResponse) SetHeaders(v map[string]*string) *ListEnterpriseSupportPlanSimpleResponse {
	s.Headers = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponse) SetStatusCode(v int32) *ListEnterpriseSupportPlanSimpleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponse) SetBody(v *ListEnterpriseSupportPlanSimpleResponseBody) *ListEnterpriseSupportPlanSimpleResponse {
	s.Body = v
	return s
}

func (s *ListEnterpriseSupportPlanSimpleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
