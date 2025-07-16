// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *RollbackStagingConfigResponseBody) *RollbackStagingConfigResponse
	GetBody() *RollbackStagingConfigResponseBody
}

type RollbackStagingConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *RollbackStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackStagingConfigResponse) GetBody() *RollbackStagingConfigResponseBody {
	return s.Body
}

func (s *RollbackStagingConfigResponse) SetHeaders(v map[string]*string) *RollbackStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *RollbackStagingConfigResponse) SetStatusCode(v int32) *RollbackStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackStagingConfigResponse) SetBody(v *RollbackStagingConfigResponseBody) *RollbackStagingConfigResponse {
	s.Body = v
	return s
}

func (s *RollbackStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
