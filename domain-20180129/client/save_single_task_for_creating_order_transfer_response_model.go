// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderTransferResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForCreatingOrderTransferResponseBody) *SaveSingleTaskForCreatingOrderTransferResponse
	GetBody() *SaveSingleTaskForCreatingOrderTransferResponseBody
}

type SaveSingleTaskForCreatingOrderTransferResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForCreatingOrderTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderTransferResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) GetBody() *SaveSingleTaskForCreatingOrderTransferResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderTransferResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) SetBody(v *SaveSingleTaskForCreatingOrderTransferResponseBody) *SaveSingleTaskForCreatingOrderTransferResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderTransferResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
