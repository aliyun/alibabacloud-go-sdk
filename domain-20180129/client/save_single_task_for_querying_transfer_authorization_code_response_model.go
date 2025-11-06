// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForQueryingTransferAuthorizationCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse
	GetBody() *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody
}

type SaveSingleTaskForQueryingTransferAuthorizationCodeResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) GetBody() *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) SetStatusCode(v int32) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) SetBody(v *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
