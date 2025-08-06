// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *GetAppPoliciesResponseBody) *GetAppPoliciesResponse
	GetBody() *GetAppPoliciesResponseBody
}

type GetAppPoliciesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppPoliciesResponse) GoString() string {
	return s.String()
}

func (s *GetAppPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppPoliciesResponse) GetBody() *GetAppPoliciesResponseBody {
	return s.Body
}

func (s *GetAppPoliciesResponse) SetHeaders(v map[string]*string) *GetAppPoliciesResponse {
	s.Headers = v
	return s
}

func (s *GetAppPoliciesResponse) SetStatusCode(v int32) *GetAppPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppPoliciesResponse) SetBody(v *GetAppPoliciesResponseBody) *GetAppPoliciesResponse {
	s.Body = v
	return s
}

func (s *GetAppPoliciesResponse) Validate() error {
	return dara.Validate(s)
}
