// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDeviceStatusWithAkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncDeviceStatusWithAkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncDeviceStatusWithAkResponse
	GetStatusCode() *int32
	SetBody(v *SyncDeviceStatusWithAkResponseBody) *SyncDeviceStatusWithAkResponse
	GetBody() *SyncDeviceStatusWithAkResponseBody
}

type SyncDeviceStatusWithAkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDeviceStatusWithAkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDeviceStatusWithAkResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncDeviceStatusWithAkResponse) GoString() string {
	return s.String()
}

func (s *SyncDeviceStatusWithAkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncDeviceStatusWithAkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDeviceStatusWithAkResponse) GetBody() *SyncDeviceStatusWithAkResponseBody {
	return s.Body
}

func (s *SyncDeviceStatusWithAkResponse) SetHeaders(v map[string]*string) *SyncDeviceStatusWithAkResponse {
	s.Headers = v
	return s
}

func (s *SyncDeviceStatusWithAkResponse) SetStatusCode(v int32) *SyncDeviceStatusWithAkResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDeviceStatusWithAkResponse) SetBody(v *SyncDeviceStatusWithAkResponseBody) *SyncDeviceStatusWithAkResponse {
	s.Body = v
	return s
}

func (s *SyncDeviceStatusWithAkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
