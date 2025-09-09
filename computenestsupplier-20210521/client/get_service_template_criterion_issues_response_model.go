// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTemplateCriterionIssuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceTemplateCriterionIssuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceTemplateCriterionIssuesResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceTemplateCriterionIssuesResponseBody) *GetServiceTemplateCriterionIssuesResponse
	GetBody() *GetServiceTemplateCriterionIssuesResponseBody
}

type GetServiceTemplateCriterionIssuesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTemplateCriterionIssuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTemplateCriterionIssuesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTemplateCriterionIssuesResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateCriterionIssuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceTemplateCriterionIssuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceTemplateCriterionIssuesResponse) GetBody() *GetServiceTemplateCriterionIssuesResponseBody {
	return s.Body
}

func (s *GetServiceTemplateCriterionIssuesResponse) SetHeaders(v map[string]*string) *GetServiceTemplateCriterionIssuesResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponse) SetStatusCode(v int32) *GetServiceTemplateCriterionIssuesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponse) SetBody(v *GetServiceTemplateCriterionIssuesResponseBody) *GetServiceTemplateCriterionIssuesResponse {
	s.Body = v
	return s
}

func (s *GetServiceTemplateCriterionIssuesResponse) Validate() error {
	return dara.Validate(s)
}
