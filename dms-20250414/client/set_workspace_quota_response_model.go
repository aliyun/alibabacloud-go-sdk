// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkspaceQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWorkspaceQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWorkspaceQuotaResponse
	GetStatusCode() *int32
	SetBody(v *SetWorkspaceQuotaResponseBody) *SetWorkspaceQuotaResponse
	GetBody() *SetWorkspaceQuotaResponseBody
}

type SetWorkspaceQuotaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWorkspaceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWorkspaceQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWorkspaceQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetWorkspaceQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWorkspaceQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWorkspaceQuotaResponse) GetBody() *SetWorkspaceQuotaResponseBody {
	return s.Body
}

func (s *SetWorkspaceQuotaResponse) SetHeaders(v map[string]*string) *SetWorkspaceQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetWorkspaceQuotaResponse) SetStatusCode(v int32) *SetWorkspaceQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWorkspaceQuotaResponse) SetBody(v *SetWorkspaceQuotaResponseBody) *SetWorkspaceQuotaResponse {
	s.Body = v
	return s
}

func (s *SetWorkspaceQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
