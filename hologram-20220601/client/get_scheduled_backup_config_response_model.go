// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledBackupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledBackupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledBackupConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledBackupConfigResponseBody) *GetScheduledBackupConfigResponse
	GetBody() *GetScheduledBackupConfigResponseBody
}

type GetScheduledBackupConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledBackupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledBackupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledBackupConfigResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledBackupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledBackupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledBackupConfigResponse) GetBody() *GetScheduledBackupConfigResponseBody {
	return s.Body
}

func (s *GetScheduledBackupConfigResponse) SetHeaders(v map[string]*string) *GetScheduledBackupConfigResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledBackupConfigResponse) SetStatusCode(v int32) *GetScheduledBackupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledBackupConfigResponse) SetBody(v *GetScheduledBackupConfigResponseBody) *GetScheduledBackupConfigResponse {
	s.Body = v
	return s
}

func (s *GetScheduledBackupConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
