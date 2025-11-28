// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadIoTDataToBlockchainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadIoTDataToBlockchainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadIoTDataToBlockchainResponse
	GetStatusCode() *int32
	SetBody(v *UploadIoTDataToBlockchainResponseBody) *UploadIoTDataToBlockchainResponse
	GetBody() *UploadIoTDataToBlockchainResponseBody
}

type UploadIoTDataToBlockchainResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadIoTDataToBlockchainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadIoTDataToBlockchainResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadIoTDataToBlockchainResponse) GoString() string {
	return s.String()
}

func (s *UploadIoTDataToBlockchainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadIoTDataToBlockchainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadIoTDataToBlockchainResponse) GetBody() *UploadIoTDataToBlockchainResponseBody {
	return s.Body
}

func (s *UploadIoTDataToBlockchainResponse) SetHeaders(v map[string]*string) *UploadIoTDataToBlockchainResponse {
	s.Headers = v
	return s
}

func (s *UploadIoTDataToBlockchainResponse) SetStatusCode(v int32) *UploadIoTDataToBlockchainResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadIoTDataToBlockchainResponse) SetBody(v *UploadIoTDataToBlockchainResponseBody) *UploadIoTDataToBlockchainResponse {
	s.Body = v
	return s
}

func (s *UploadIoTDataToBlockchainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
