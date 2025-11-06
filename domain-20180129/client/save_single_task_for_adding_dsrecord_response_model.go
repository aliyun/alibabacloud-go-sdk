// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForAddingDSRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForAddingDSRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForAddingDSRecordResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForAddingDSRecordResponseBody) *SaveSingleTaskForAddingDSRecordResponse
	GetBody() *SaveSingleTaskForAddingDSRecordResponseBody
}

type SaveSingleTaskForAddingDSRecordResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForAddingDSRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForAddingDSRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForAddingDSRecordResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAddingDSRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForAddingDSRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForAddingDSRecordResponse) GetBody() *SaveSingleTaskForAddingDSRecordResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForAddingDSRecordResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForAddingDSRecordResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordResponse) SetStatusCode(v int32) *SaveSingleTaskForAddingDSRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordResponse) SetBody(v *SaveSingleTaskForAddingDSRecordResponseBody) *SaveSingleTaskForAddingDSRecordResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForAddingDSRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
