// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataMaskingEncryptionAlgorithmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataMaskingEncryptionAlgorithmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataMaskingEncryptionAlgorithmsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataMaskingEncryptionAlgorithmsResponseBody) *ListDataMaskingEncryptionAlgorithmsResponse
	GetBody() *ListDataMaskingEncryptionAlgorithmsResponseBody
}

type ListDataMaskingEncryptionAlgorithmsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataMaskingEncryptionAlgorithmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataMaskingEncryptionAlgorithmsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataMaskingEncryptionAlgorithmsResponse) GoString() string {
	return s.String()
}

func (s *ListDataMaskingEncryptionAlgorithmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataMaskingEncryptionAlgorithmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataMaskingEncryptionAlgorithmsResponse) GetBody() *ListDataMaskingEncryptionAlgorithmsResponseBody {
	return s.Body
}

func (s *ListDataMaskingEncryptionAlgorithmsResponse) SetHeaders(v map[string]*string) *ListDataMaskingEncryptionAlgorithmsResponse {
	s.Headers = v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsResponse) SetStatusCode(v int32) *ListDataMaskingEncryptionAlgorithmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsResponse) SetBody(v *ListDataMaskingEncryptionAlgorithmsResponseBody) *ListDataMaskingEncryptionAlgorithmsResponse {
	s.Body = v
	return s
}

func (s *ListDataMaskingEncryptionAlgorithmsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
