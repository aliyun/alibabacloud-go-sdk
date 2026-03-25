// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKMSSecretsForHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportKMSSecretsForHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportKMSSecretsForHostResponse
	GetStatusCode() *int32
	SetBody(v *ImportKMSSecretsForHostResponseBody) *ImportKMSSecretsForHostResponse
	GetBody() *ImportKMSSecretsForHostResponseBody
}

type ImportKMSSecretsForHostResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportKMSSecretsForHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportKMSSecretsForHostResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportKMSSecretsForHostResponse) GoString() string {
	return s.String()
}

func (s *ImportKMSSecretsForHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportKMSSecretsForHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportKMSSecretsForHostResponse) GetBody() *ImportKMSSecretsForHostResponseBody {
	return s.Body
}

func (s *ImportKMSSecretsForHostResponse) SetHeaders(v map[string]*string) *ImportKMSSecretsForHostResponse {
	s.Headers = v
	return s
}

func (s *ImportKMSSecretsForHostResponse) SetStatusCode(v int32) *ImportKMSSecretsForHostResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKMSSecretsForHostResponse) SetBody(v *ImportKMSSecretsForHostResponseBody) *ImportKMSSecretsForHostResponse {
	s.Body = v
	return s
}

func (s *ImportKMSSecretsForHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
