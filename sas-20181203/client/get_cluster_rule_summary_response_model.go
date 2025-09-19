// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterRuleSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterRuleSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterRuleSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterRuleSummaryResponseBody) *GetClusterRuleSummaryResponse
	GetBody() *GetClusterRuleSummaryResponseBody
}

type GetClusterRuleSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterRuleSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterRuleSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRuleSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetClusterRuleSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterRuleSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterRuleSummaryResponse) GetBody() *GetClusterRuleSummaryResponseBody {
	return s.Body
}

func (s *GetClusterRuleSummaryResponse) SetHeaders(v map[string]*string) *GetClusterRuleSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetClusterRuleSummaryResponse) SetStatusCode(v int32) *GetClusterRuleSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterRuleSummaryResponse) SetBody(v *GetClusterRuleSummaryResponseBody) *GetClusterRuleSummaryResponse {
	s.Body = v
	return s
}

func (s *GetClusterRuleSummaryResponse) Validate() error {
	return dara.Validate(s)
}
