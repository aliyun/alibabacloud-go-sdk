// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspaceQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspaceQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkspaceQuotaResponseBody) *GetWorkspaceQuotaResponse
	GetBody() *GetWorkspaceQuotaResponseBody
}

type GetWorkspaceQuotaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspaceQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspaceQuotaResponse) GetBody() *GetWorkspaceQuotaResponseBody {
	return s.Body
}

func (s *GetWorkspaceQuotaResponse) SetHeaders(v map[string]*string) *GetWorkspaceQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceQuotaResponse) SetStatusCode(v int32) *GetWorkspaceQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceQuotaResponse) SetBody(v *GetWorkspaceQuotaResponseBody) *GetWorkspaceQuotaResponse {
	s.Body = v
	return s
}

func (s *GetWorkspaceQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
