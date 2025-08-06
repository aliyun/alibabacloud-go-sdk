// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackVodStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackVodStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackVodStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *RollbackVodStagingConfigResponseBody) *RollbackVodStagingConfigResponse
	GetBody() *RollbackVodStagingConfigResponseBody
}

type RollbackVodStagingConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackVodStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackVodStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackVodStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *RollbackVodStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackVodStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackVodStagingConfigResponse) GetBody() *RollbackVodStagingConfigResponseBody {
	return s.Body
}

func (s *RollbackVodStagingConfigResponse) SetHeaders(v map[string]*string) *RollbackVodStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *RollbackVodStagingConfigResponse) SetStatusCode(v int32) *RollbackVodStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackVodStagingConfigResponse) SetBody(v *RollbackVodStagingConfigResponseBody) *RollbackVodStagingConfigResponse {
	s.Body = v
	return s
}

func (s *RollbackVodStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
