// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeUpgradeTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMcubeUpgradeTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMcubeUpgradeTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMcubeUpgradeTaskInfoResponseBody) *GetMcubeUpgradeTaskInfoResponse
	GetBody() *GetMcubeUpgradeTaskInfoResponseBody
}

type GetMcubeUpgradeTaskInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMcubeUpgradeTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMcubeUpgradeTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradeTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradeTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcubeUpgradeTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMcubeUpgradeTaskInfoResponse) GetBody() *GetMcubeUpgradeTaskInfoResponseBody {
	return s.Body
}

func (s *GetMcubeUpgradeTaskInfoResponse) SetHeaders(v map[string]*string) *GetMcubeUpgradeTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponse) SetStatusCode(v int32) *GetMcubeUpgradeTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponse) SetBody(v *GetMcubeUpgradeTaskInfoResponseBody) *GetMcubeUpgradeTaskInfoResponse {
	s.Body = v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
