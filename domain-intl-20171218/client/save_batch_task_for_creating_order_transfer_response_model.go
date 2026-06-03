// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderTransferResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForCreatingOrderTransferResponseBody) *SaveBatchTaskForCreatingOrderTransferResponse
	GetBody() *SaveBatchTaskForCreatingOrderTransferResponseBody
}

type SaveBatchTaskForCreatingOrderTransferResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForCreatingOrderTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderTransferResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) GetBody() *SaveBatchTaskForCreatingOrderTransferResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderTransferResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) SetBody(v *SaveBatchTaskForCreatingOrderTransferResponseBody) *SaveBatchTaskForCreatingOrderTransferResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderTransferResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
