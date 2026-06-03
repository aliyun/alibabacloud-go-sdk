// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdateProhibitionLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForUpdateProhibitionLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForUpdateProhibitionLockResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForUpdateProhibitionLockResponseBody) *SaveBatchTaskForUpdateProhibitionLockResponse
	GetBody() *SaveBatchTaskForUpdateProhibitionLockResponseBody
}

type SaveBatchTaskForUpdateProhibitionLockResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForUpdateProhibitionLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForUpdateProhibitionLockResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdateProhibitionLockResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) GetBody() *SaveBatchTaskForUpdateProhibitionLockResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForUpdateProhibitionLockResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) SetStatusCode(v int32) *SaveBatchTaskForUpdateProhibitionLockResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) SetBody(v *SaveBatchTaskForUpdateProhibitionLockResponseBody) *SaveBatchTaskForUpdateProhibitionLockResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForUpdateProhibitionLockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
