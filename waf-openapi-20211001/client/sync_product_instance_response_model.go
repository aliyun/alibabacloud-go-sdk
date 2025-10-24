// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncProductInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncProductInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncProductInstanceResponse
	GetStatusCode() *int32
	SetBody(v *SyncProductInstanceResponseBody) *SyncProductInstanceResponse
	GetBody() *SyncProductInstanceResponseBody
}

type SyncProductInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncProductInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncProductInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncProductInstanceResponse) GoString() string {
	return s.String()
}

func (s *SyncProductInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncProductInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncProductInstanceResponse) GetBody() *SyncProductInstanceResponseBody {
	return s.Body
}

func (s *SyncProductInstanceResponse) SetHeaders(v map[string]*string) *SyncProductInstanceResponse {
	s.Headers = v
	return s
}

func (s *SyncProductInstanceResponse) SetStatusCode(v int32) *SyncProductInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncProductInstanceResponse) SetBody(v *SyncProductInstanceResponseBody) *SyncProductInstanceResponse {
	s.Body = v
	return s
}

func (s *SyncProductInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
