// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCancelingTransferInResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForCancelingTransferInResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForCancelingTransferInResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForCancelingTransferInResponseBody) *SaveSingleTaskForCancelingTransferInResponse
	GetBody() *SaveSingleTaskForCancelingTransferInResponseBody
}

type SaveSingleTaskForCancelingTransferInResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForCancelingTransferInResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForCancelingTransferInResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCancelingTransferInResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCancelingTransferInResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForCancelingTransferInResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForCancelingTransferInResponse) GetBody() *SaveSingleTaskForCancelingTransferInResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForCancelingTransferInResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCancelingTransferInResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInResponse) SetStatusCode(v int32) *SaveSingleTaskForCancelingTransferInResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInResponse) SetBody(v *SaveSingleTaskForCancelingTransferInResponseBody) *SaveSingleTaskForCancelingTransferInResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForCancelingTransferInResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
