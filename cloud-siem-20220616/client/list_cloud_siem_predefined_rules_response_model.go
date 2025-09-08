// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudSiemPredefinedRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudSiemPredefinedRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudSiemPredefinedRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudSiemPredefinedRulesResponseBody) *ListCloudSiemPredefinedRulesResponse
	GetBody() *ListCloudSiemPredefinedRulesResponseBody
}

type ListCloudSiemPredefinedRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudSiemPredefinedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudSiemPredefinedRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemPredefinedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudSiemPredefinedRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudSiemPredefinedRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudSiemPredefinedRulesResponse) GetBody() *ListCloudSiemPredefinedRulesResponseBody {
	return s.Body
}

func (s *ListCloudSiemPredefinedRulesResponse) SetHeaders(v map[string]*string) *ListCloudSiemPredefinedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponse) SetStatusCode(v int32) *ListCloudSiemPredefinedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponse) SetBody(v *ListCloudSiemPredefinedRulesResponseBody) *ListCloudSiemPredefinedRulesResponse {
	s.Body = v
	return s
}

func (s *ListCloudSiemPredefinedRulesResponse) Validate() error {
	return dara.Validate(s)
}
