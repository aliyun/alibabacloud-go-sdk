// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSaveArtExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForSaveArtExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForSaveArtExtensionResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForSaveArtExtensionResponseBody) *SaveSingleTaskForSaveArtExtensionResponse
	GetBody() *SaveSingleTaskForSaveArtExtensionResponseBody
}

type SaveSingleTaskForSaveArtExtensionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForSaveArtExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForSaveArtExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSaveArtExtensionResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) GetBody() *SaveSingleTaskForSaveArtExtensionResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForSaveArtExtensionResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) SetStatusCode(v int32) *SaveSingleTaskForSaveArtExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) SetBody(v *SaveSingleTaskForSaveArtExtensionResponseBody) *SaveSingleTaskForSaveArtExtensionResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
