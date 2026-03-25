// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportableKMSSecretsForHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImportableKMSSecretsForHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImportableKMSSecretsForHostResponse
	GetStatusCode() *int32
	SetBody(v *ListImportableKMSSecretsForHostResponseBody) *ListImportableKMSSecretsForHostResponse
	GetBody() *ListImportableKMSSecretsForHostResponseBody
}

type ListImportableKMSSecretsForHostResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImportableKMSSecretsForHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImportableKMSSecretsForHostResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImportableKMSSecretsForHostResponse) GoString() string {
	return s.String()
}

func (s *ListImportableKMSSecretsForHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImportableKMSSecretsForHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImportableKMSSecretsForHostResponse) GetBody() *ListImportableKMSSecretsForHostResponseBody {
	return s.Body
}

func (s *ListImportableKMSSecretsForHostResponse) SetHeaders(v map[string]*string) *ListImportableKMSSecretsForHostResponse {
	s.Headers = v
	return s
}

func (s *ListImportableKMSSecretsForHostResponse) SetStatusCode(v int32) *ListImportableKMSSecretsForHostResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImportableKMSSecretsForHostResponse) SetBody(v *ListImportableKMSSecretsForHostResponseBody) *ListImportableKMSSecretsForHostResponse {
	s.Body = v
	return s
}

func (s *ListImportableKMSSecretsForHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
