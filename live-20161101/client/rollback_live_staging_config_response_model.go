// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackLiveStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackLiveStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackLiveStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *RollbackLiveStagingConfigResponseBody) *RollbackLiveStagingConfigResponse
	GetBody() *RollbackLiveStagingConfigResponseBody
}

type RollbackLiveStagingConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackLiveStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackLiveStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackLiveStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *RollbackLiveStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackLiveStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackLiveStagingConfigResponse) GetBody() *RollbackLiveStagingConfigResponseBody {
	return s.Body
}

func (s *RollbackLiveStagingConfigResponse) SetHeaders(v map[string]*string) *RollbackLiveStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *RollbackLiveStagingConfigResponse) SetStatusCode(v int32) *RollbackLiveStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackLiveStagingConfigResponse) SetBody(v *RollbackLiveStagingConfigResponseBody) *RollbackLiveStagingConfigResponse {
	s.Body = v
	return s
}

func (s *RollbackLiveStagingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
