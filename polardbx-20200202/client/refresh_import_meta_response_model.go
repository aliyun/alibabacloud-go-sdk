// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshImportMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshImportMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshImportMetaResponse
	GetStatusCode() *int32
	SetBody(v *RefreshImportMetaResponseBody) *RefreshImportMetaResponse
	GetBody() *RefreshImportMetaResponseBody
}

type RefreshImportMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshImportMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshImportMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshImportMetaResponse) GoString() string {
	return s.String()
}

func (s *RefreshImportMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshImportMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshImportMetaResponse) GetBody() *RefreshImportMetaResponseBody {
	return s.Body
}

func (s *RefreshImportMetaResponse) SetHeaders(v map[string]*string) *RefreshImportMetaResponse {
	s.Headers = v
	return s
}

func (s *RefreshImportMetaResponse) SetStatusCode(v int32) *RefreshImportMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshImportMetaResponse) SetBody(v *RefreshImportMetaResponseBody) *RefreshImportMetaResponse {
	s.Body = v
	return s
}

func (s *RefreshImportMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
