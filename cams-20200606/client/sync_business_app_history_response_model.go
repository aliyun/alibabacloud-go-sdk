// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncBusinessAppHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncBusinessAppHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncBusinessAppHistoryResponse
	GetStatusCode() *int32
	SetBody(v *SyncBusinessAppHistoryResponseBody) *SyncBusinessAppHistoryResponse
	GetBody() *SyncBusinessAppHistoryResponseBody
}

type SyncBusinessAppHistoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncBusinessAppHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncBusinessAppHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncBusinessAppHistoryResponse) GoString() string {
	return s.String()
}

func (s *SyncBusinessAppHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncBusinessAppHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncBusinessAppHistoryResponse) GetBody() *SyncBusinessAppHistoryResponseBody {
	return s.Body
}

func (s *SyncBusinessAppHistoryResponse) SetHeaders(v map[string]*string) *SyncBusinessAppHistoryResponse {
	s.Headers = v
	return s
}

func (s *SyncBusinessAppHistoryResponse) SetStatusCode(v int32) *SyncBusinessAppHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncBusinessAppHistoryResponse) SetBody(v *SyncBusinessAppHistoryResponseBody) *SyncBusinessAppHistoryResponse {
	s.Body = v
	return s
}

func (s *SyncBusinessAppHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
