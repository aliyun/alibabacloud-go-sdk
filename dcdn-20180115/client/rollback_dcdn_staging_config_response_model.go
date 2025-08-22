// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackDcdnStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackDcdnStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackDcdnStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *RollbackDcdnStagingConfigResponseBody) *RollbackDcdnStagingConfigResponse
	GetBody() *RollbackDcdnStagingConfigResponseBody
}

type RollbackDcdnStagingConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackDcdnStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackDcdnStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackDcdnStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *RollbackDcdnStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackDcdnStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackDcdnStagingConfigResponse) GetBody() *RollbackDcdnStagingConfigResponseBody {
	return s.Body
}

func (s *RollbackDcdnStagingConfigResponse) SetHeaders(v map[string]*string) *RollbackDcdnStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *RollbackDcdnStagingConfigResponse) SetStatusCode(v int32) *RollbackDcdnStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackDcdnStagingConfigResponse) SetBody(v *RollbackDcdnStagingConfigResponseBody) *RollbackDcdnStagingConfigResponse {
	s.Body = v
	return s
}

func (s *RollbackDcdnStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
