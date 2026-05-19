// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientRuleDashboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileProtectClientRuleDashboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileProtectClientRuleDashboardResponse
	GetStatusCode() *int32
	SetBody(v *GetFileProtectClientRuleDashboardResponseBody) *GetFileProtectClientRuleDashboardResponse
	GetBody() *GetFileProtectClientRuleDashboardResponseBody
}

type GetFileProtectClientRuleDashboardResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileProtectClientRuleDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileProtectClientRuleDashboardResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientRuleDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientRuleDashboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileProtectClientRuleDashboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileProtectClientRuleDashboardResponse) GetBody() *GetFileProtectClientRuleDashboardResponseBody {
	return s.Body
}

func (s *GetFileProtectClientRuleDashboardResponse) SetHeaders(v map[string]*string) *GetFileProtectClientRuleDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponse) SetStatusCode(v int32) *GetFileProtectClientRuleDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponse) SetBody(v *GetFileProtectClientRuleDashboardResponseBody) *GetFileProtectClientRuleDashboardResponse {
	s.Body = v
	return s
}

func (s *GetFileProtectClientRuleDashboardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
