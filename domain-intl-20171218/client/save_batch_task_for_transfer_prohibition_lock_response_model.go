// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForTransferProhibitionLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForTransferProhibitionLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForTransferProhibitionLockResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForTransferProhibitionLockResponseBody) *SaveBatchTaskForTransferProhibitionLockResponse
	GetBody() *SaveBatchTaskForTransferProhibitionLockResponseBody
}

type SaveBatchTaskForTransferProhibitionLockResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForTransferProhibitionLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForTransferProhibitionLockResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForTransferProhibitionLockResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) GetBody() *SaveBatchTaskForTransferProhibitionLockResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForTransferProhibitionLockResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) SetStatusCode(v int32) *SaveBatchTaskForTransferProhibitionLockResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) SetBody(v *SaveBatchTaskForTransferProhibitionLockResponseBody) *SaveBatchTaskForTransferProhibitionLockResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForTransferProhibitionLockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
