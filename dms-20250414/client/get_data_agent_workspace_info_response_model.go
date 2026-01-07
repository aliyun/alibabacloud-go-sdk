// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentWorkspaceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataAgentWorkspaceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataAgentWorkspaceInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDataAgentWorkspaceInfoResponseBody) *GetDataAgentWorkspaceInfoResponse
	GetBody() *GetDataAgentWorkspaceInfoResponseBody
}

type GetDataAgentWorkspaceInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataAgentWorkspaceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataAgentWorkspaceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentWorkspaceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDataAgentWorkspaceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataAgentWorkspaceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataAgentWorkspaceInfoResponse) GetBody() *GetDataAgentWorkspaceInfoResponseBody {
	return s.Body
}

func (s *GetDataAgentWorkspaceInfoResponse) SetHeaders(v map[string]*string) *GetDataAgentWorkspaceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponse) SetStatusCode(v int32) *GetDataAgentWorkspaceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponse) SetBody(v *GetDataAgentWorkspaceInfoResponseBody) *GetDataAgentWorkspaceInfoResponse {
	s.Body = v
	return s
}

func (s *GetDataAgentWorkspaceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
