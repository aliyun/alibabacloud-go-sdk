// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForTransferOutByAuthorizationCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForTransferOutByAuthorizationCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForTransferOutByAuthorizationCodeResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) *SaveBatchTaskForTransferOutByAuthorizationCodeResponse
	GetBody() *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody
}

type SaveBatchTaskForTransferOutByAuthorizationCodeResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForTransferOutByAuthorizationCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForTransferOutByAuthorizationCodeResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponse) GetBody() *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForTransferOutByAuthorizationCodeResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponse) SetStatusCode(v int32) *SaveBatchTaskForTransferOutByAuthorizationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponse) SetBody(v *SaveBatchTaskForTransferOutByAuthorizationCodeResponseBody) *SaveBatchTaskForTransferOutByAuthorizationCodeResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForTransferOutByAuthorizationCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
