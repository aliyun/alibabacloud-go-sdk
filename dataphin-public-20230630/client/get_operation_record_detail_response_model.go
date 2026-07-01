// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOperationRecordDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOperationRecordDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetOperationRecordDetailResponseBody) *GetOperationRecordDetailResponse
	GetBody() *GetOperationRecordDetailResponseBody
}

type GetOperationRecordDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperationRecordDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperationRecordDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOperationRecordDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOperationRecordDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOperationRecordDetailResponse) GetBody() *GetOperationRecordDetailResponseBody {
	return s.Body
}

func (s *GetOperationRecordDetailResponse) SetHeaders(v map[string]*string) *GetOperationRecordDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOperationRecordDetailResponse) SetStatusCode(v int32) *GetOperationRecordDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperationRecordDetailResponse) SetBody(v *GetOperationRecordDetailResponseBody) *GetOperationRecordDetailResponse {
	s.Body = v
	return s
}

func (s *GetOperationRecordDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
