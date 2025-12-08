// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeUpgradeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcubeUpgradeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcubeUpgradeResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcubeUpgradeResourceResponseBody) *DeleteMcubeUpgradeResourceResponse
	GetBody() *DeleteMcubeUpgradeResourceResponseBody
}

type DeleteMcubeUpgradeResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcubeUpgradeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcubeUpgradeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeUpgradeResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcubeUpgradeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcubeUpgradeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcubeUpgradeResourceResponse) GetBody() *DeleteMcubeUpgradeResourceResponseBody {
	return s.Body
}

func (s *DeleteMcubeUpgradeResourceResponse) SetHeaders(v map[string]*string) *DeleteMcubeUpgradeResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponse) SetStatusCode(v int32) *DeleteMcubeUpgradeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponse) SetBody(v *DeleteMcubeUpgradeResourceResponseBody) *DeleteMcubeUpgradeResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteMcubeUpgradeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
