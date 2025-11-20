// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDingTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncDingTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncDingTypeResponse
	GetStatusCode() *int32
	SetBody(v *SyncDingTypeResponseBody) *SyncDingTypeResponse
	GetBody() *SyncDingTypeResponseBody
}

type SyncDingTypeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDingTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDingTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncDingTypeResponse) GoString() string {
	return s.String()
}

func (s *SyncDingTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncDingTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDingTypeResponse) GetBody() *SyncDingTypeResponseBody {
	return s.Body
}

func (s *SyncDingTypeResponse) SetHeaders(v map[string]*string) *SyncDingTypeResponse {
	s.Headers = v
	return s
}

func (s *SyncDingTypeResponse) SetStatusCode(v int32) *SyncDingTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDingTypeResponse) SetBody(v *SyncDingTypeResponseBody) *SyncDingTypeResponse {
	s.Body = v
	return s
}

func (s *SyncDingTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
