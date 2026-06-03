// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderRenewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderRenewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderRenewResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForCreatingOrderRenewResponseBody) *SaveBatchTaskForCreatingOrderRenewResponse
	GetBody() *SaveBatchTaskForCreatingOrderRenewResponseBody
}

type SaveBatchTaskForCreatingOrderRenewResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForCreatingOrderRenewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRenewResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRenewResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) GetBody() *SaveBatchTaskForCreatingOrderRenewResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderRenewResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) SetBody(v *SaveBatchTaskForCreatingOrderRenewResponseBody) *SaveBatchTaskForCreatingOrderRenewResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRenewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
