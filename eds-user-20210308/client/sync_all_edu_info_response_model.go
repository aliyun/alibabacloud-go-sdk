// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncAllEduInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncAllEduInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncAllEduInfoResponse
	GetStatusCode() *int32
	SetBody(v *SyncAllEduInfoResponseBody) *SyncAllEduInfoResponse
	GetBody() *SyncAllEduInfoResponseBody
}

type SyncAllEduInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncAllEduInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncAllEduInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncAllEduInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncAllEduInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncAllEduInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncAllEduInfoResponse) GetBody() *SyncAllEduInfoResponseBody {
	return s.Body
}

func (s *SyncAllEduInfoResponse) SetHeaders(v map[string]*string) *SyncAllEduInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncAllEduInfoResponse) SetStatusCode(v int32) *SyncAllEduInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncAllEduInfoResponse) SetBody(v *SyncAllEduInfoResponseBody) *SyncAllEduInfoResponse {
	s.Body = v
	return s
}

func (s *SyncAllEduInfoResponse) Validate() error {
	return dara.Validate(s)
}
