// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForApprovingTransferOutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForApprovingTransferOutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForApprovingTransferOutResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForApprovingTransferOutResponseBody) *SaveSingleTaskForApprovingTransferOutResponse
	GetBody() *SaveSingleTaskForApprovingTransferOutResponseBody
}

type SaveSingleTaskForApprovingTransferOutResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForApprovingTransferOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForApprovingTransferOutResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForApprovingTransferOutResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) GetBody() *SaveSingleTaskForApprovingTransferOutResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForApprovingTransferOutResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) SetStatusCode(v int32) *SaveSingleTaskForApprovingTransferOutResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) SetBody(v *SaveSingleTaskForApprovingTransferOutResponseBody) *SaveSingleTaskForApprovingTransferOutResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForApprovingTransferOutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
