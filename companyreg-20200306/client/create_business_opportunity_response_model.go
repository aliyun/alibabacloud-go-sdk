// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBusinessOpportunityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBusinessOpportunityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBusinessOpportunityResponse
	GetStatusCode() *int32
	SetBody(v *CreateBusinessOpportunityResponseBody) *CreateBusinessOpportunityResponse
	GetBody() *CreateBusinessOpportunityResponseBody
}

type CreateBusinessOpportunityResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBusinessOpportunityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBusinessOpportunityResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBusinessOpportunityResponse) GoString() string {
	return s.String()
}

func (s *CreateBusinessOpportunityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBusinessOpportunityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBusinessOpportunityResponse) GetBody() *CreateBusinessOpportunityResponseBody {
	return s.Body
}

func (s *CreateBusinessOpportunityResponse) SetHeaders(v map[string]*string) *CreateBusinessOpportunityResponse {
	s.Headers = v
	return s
}

func (s *CreateBusinessOpportunityResponse) SetStatusCode(v int32) *CreateBusinessOpportunityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBusinessOpportunityResponse) SetBody(v *CreateBusinessOpportunityResponseBody) *CreateBusinessOpportunityResponse {
	s.Body = v
	return s
}

func (s *CreateBusinessOpportunityResponse) Validate() error {
	return dara.Validate(s)
}
