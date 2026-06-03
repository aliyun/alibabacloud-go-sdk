// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSynchronizingDSRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForSynchronizingDSRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForSynchronizingDSRecordResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForSynchronizingDSRecordResponseBody) *SaveSingleTaskForSynchronizingDSRecordResponse
	GetBody() *SaveSingleTaskForSynchronizingDSRecordResponseBody
}

type SaveSingleTaskForSynchronizingDSRecordResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForSynchronizingDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDSRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDSRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) GetBody() *SaveSingleTaskForSynchronizingDSRecordResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForSynchronizingDSRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) SetStatusCode(v int32) *SaveSingleTaskForSynchronizingDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) SetBody(v *SaveSingleTaskForSynchronizingDSRecordResponseBody) *SaveSingleTaskForSynchronizingDSRecordResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForSynchronizingDSRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
