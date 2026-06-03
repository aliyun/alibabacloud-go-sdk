// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDeletingDSRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForDeletingDSRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForDeletingDSRecordResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForDeletingDSRecordResponseBody) *SaveSingleTaskForDeletingDSRecordResponse
	GetBody() *SaveSingleTaskForDeletingDSRecordResponseBody
}

type SaveSingleTaskForDeletingDSRecordResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForDeletingDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForDeletingDSRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDeletingDSRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) GetBody() *SaveSingleTaskForDeletingDSRecordResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForDeletingDSRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) SetStatusCode(v int32) *SaveSingleTaskForDeletingDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) SetBody(v *SaveSingleTaskForDeletingDSRecordResponseBody) *SaveSingleTaskForDeletingDSRecordResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForDeletingDSRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
