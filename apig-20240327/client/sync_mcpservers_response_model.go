// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMCPServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncMCPServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncMCPServersResponse
	GetStatusCode() *int32
	SetBody(v *SyncMCPServersResponseBody) *SyncMCPServersResponse
	GetBody() *SyncMCPServersResponseBody
}

type SyncMCPServersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncMCPServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncMCPServersResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncMCPServersResponse) GoString() string {
	return s.String()
}

func (s *SyncMCPServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncMCPServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncMCPServersResponse) GetBody() *SyncMCPServersResponseBody {
	return s.Body
}

func (s *SyncMCPServersResponse) SetHeaders(v map[string]*string) *SyncMCPServersResponse {
	s.Headers = v
	return s
}

func (s *SyncMCPServersResponse) SetStatusCode(v int32) *SyncMCPServersResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncMCPServersResponse) SetBody(v *SyncMCPServersResponseBody) *SyncMCPServersResponse {
	s.Body = v
	return s
}

func (s *SyncMCPServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
