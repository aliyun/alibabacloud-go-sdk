// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartContextDatabaseSyncSourceDiscoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartContextDatabaseSyncSourceDiscoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartContextDatabaseSyncSourceDiscoveryResponse
	GetStatusCode() *int32
	SetBody(v *StartContextDatabaseSyncSourceDiscoveryResponseBody) *StartContextDatabaseSyncSourceDiscoveryResponse
	GetBody() *StartContextDatabaseSyncSourceDiscoveryResponseBody
}

type StartContextDatabaseSyncSourceDiscoveryResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartContextDatabaseSyncSourceDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartContextDatabaseSyncSourceDiscoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s StartContextDatabaseSyncSourceDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponse) GetBody() *StartContextDatabaseSyncSourceDiscoveryResponseBody {
	return s.Body
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponse) SetHeaders(v map[string]*string) *StartContextDatabaseSyncSourceDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponse) SetStatusCode(v int32) *StartContextDatabaseSyncSourceDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponse) SetBody(v *StartContextDatabaseSyncSourceDiscoveryResponseBody) *StartContextDatabaseSyncSourceDiscoveryResponse {
	s.Body = v
	return s
}

func (s *StartContextDatabaseSyncSourceDiscoveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
