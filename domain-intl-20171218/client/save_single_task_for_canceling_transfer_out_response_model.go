// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCancelingTransferOutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForCancelingTransferOutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForCancelingTransferOutResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForCancelingTransferOutResponseBody) *SaveSingleTaskForCancelingTransferOutResponse
	GetBody() *SaveSingleTaskForCancelingTransferOutResponseBody
}

type SaveSingleTaskForCancelingTransferOutResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForCancelingTransferOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForCancelingTransferOutResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferOutResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) GetBody() *SaveSingleTaskForCancelingTransferOutResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCancelingTransferOutResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) SetStatusCode(v int32) *SaveSingleTaskForCancelingTransferOutResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) SetBody(v *SaveSingleTaskForCancelingTransferOutResponseBody) *SaveSingleTaskForCancelingTransferOutResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForCancelingTransferOutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
