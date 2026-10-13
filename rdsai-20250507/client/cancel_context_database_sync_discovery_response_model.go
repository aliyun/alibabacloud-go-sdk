// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelContextDatabaseSyncDiscoveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelContextDatabaseSyncDiscoveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelContextDatabaseSyncDiscoveryResponse
	GetStatusCode() *int32
	SetBody(v *CancelContextDatabaseSyncDiscoveryResponseBody) *CancelContextDatabaseSyncDiscoveryResponse
	GetBody() *CancelContextDatabaseSyncDiscoveryResponseBody
}

type CancelContextDatabaseSyncDiscoveryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelContextDatabaseSyncDiscoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelContextDatabaseSyncDiscoveryResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelContextDatabaseSyncDiscoveryResponse) GoString() string {
	return s.String()
}

func (s *CancelContextDatabaseSyncDiscoveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelContextDatabaseSyncDiscoveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelContextDatabaseSyncDiscoveryResponse) GetBody() *CancelContextDatabaseSyncDiscoveryResponseBody {
	return s.Body
}

func (s *CancelContextDatabaseSyncDiscoveryResponse) SetHeaders(v map[string]*string) *CancelContextDatabaseSyncDiscoveryResponse {
	s.Headers = v
	return s
}

func (s *CancelContextDatabaseSyncDiscoveryResponse) SetStatusCode(v int32) *CancelContextDatabaseSyncDiscoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelContextDatabaseSyncDiscoveryResponse) SetBody(v *CancelContextDatabaseSyncDiscoveryResponseBody) *CancelContextDatabaseSyncDiscoveryResponse {
	s.Body = v
	return s
}

func (s *CancelContextDatabaseSyncDiscoveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
