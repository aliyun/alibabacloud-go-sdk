// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudSiemCustomizeRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudSiemCustomizeRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudSiemCustomizeRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudSiemCustomizeRulesResponseBody) *ListCloudSiemCustomizeRulesResponse
	GetBody() *ListCloudSiemCustomizeRulesResponseBody
}

type ListCloudSiemCustomizeRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudSiemCustomizeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudSiemCustomizeRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudSiemCustomizeRulesResponse) GetBody() *ListCloudSiemCustomizeRulesResponseBody {
	return s.Body
}

func (s *ListCloudSiemCustomizeRulesResponse) SetHeaders(v map[string]*string) *ListCloudSiemCustomizeRulesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponse) SetStatusCode(v int32) *ListCloudSiemCustomizeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponse) SetBody(v *ListCloudSiemCustomizeRulesResponseBody) *ListCloudSiemCustomizeRulesResponse {
	s.Body = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
