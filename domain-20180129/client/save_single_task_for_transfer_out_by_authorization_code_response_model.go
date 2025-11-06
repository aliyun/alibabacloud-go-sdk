// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForTransferOutByAuthorizationCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForTransferOutByAuthorizationCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForTransferOutByAuthorizationCodeResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) *SaveSingleTaskForTransferOutByAuthorizationCodeResponse
	GetBody() *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody
}

type SaveSingleTaskForTransferOutByAuthorizationCodeResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForTransferOutByAuthorizationCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForTransferOutByAuthorizationCodeResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponse) GetBody() *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForTransferOutByAuthorizationCodeResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponse) SetStatusCode(v int32) *SaveSingleTaskForTransferOutByAuthorizationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponse) SetBody(v *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) *SaveSingleTaskForTransferOutByAuthorizationCodeResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
