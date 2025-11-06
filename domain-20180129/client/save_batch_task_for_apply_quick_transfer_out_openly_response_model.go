// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForApplyQuickTransferOutOpenlyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse
	GetBody() *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody
}

type SaveBatchTaskForApplyQuickTransferOutOpenlyResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) GetBody() *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) SetStatusCode(v int32) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) SetBody(v *SaveBatchTaskForApplyQuickTransferOutOpenlyResponseBody) *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
