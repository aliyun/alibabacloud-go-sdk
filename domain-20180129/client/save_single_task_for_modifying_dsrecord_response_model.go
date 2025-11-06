// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForModifyingDSRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForModifyingDSRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForModifyingDSRecordResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForModifyingDSRecordResponseBody) *SaveSingleTaskForModifyingDSRecordResponse
	GetBody() *SaveSingleTaskForModifyingDSRecordResponseBody
}

type SaveSingleTaskForModifyingDSRecordResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForModifyingDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForModifyingDSRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForModifyingDSRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) GetBody() *SaveSingleTaskForModifyingDSRecordResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForModifyingDSRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) SetStatusCode(v int32) *SaveSingleTaskForModifyingDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) SetBody(v *SaveSingleTaskForModifyingDSRecordResponseBody) *SaveSingleTaskForModifyingDSRecordResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForModifyingDSRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
