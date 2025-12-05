// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKmsInstanceQuotaInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKmsInstanceQuotaInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKmsInstanceQuotaInfosResponse
	GetStatusCode() *int32
	SetBody(v *GetKmsInstanceQuotaInfosResponseBody) *GetKmsInstanceQuotaInfosResponse
	GetBody() *GetKmsInstanceQuotaInfosResponseBody
}

type GetKmsInstanceQuotaInfosResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKmsInstanceQuotaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKmsInstanceQuotaInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKmsInstanceQuotaInfosResponse) GoString() string {
	return s.String()
}

func (s *GetKmsInstanceQuotaInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKmsInstanceQuotaInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKmsInstanceQuotaInfosResponse) GetBody() *GetKmsInstanceQuotaInfosResponseBody {
	return s.Body
}

func (s *GetKmsInstanceQuotaInfosResponse) SetHeaders(v map[string]*string) *GetKmsInstanceQuotaInfosResponse {
	s.Headers = v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponse) SetStatusCode(v int32) *GetKmsInstanceQuotaInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponse) SetBody(v *GetKmsInstanceQuotaInfosResponseBody) *GetKmsInstanceQuotaInfosResponse {
	s.Body = v
	return s
}

func (s *GetKmsInstanceQuotaInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
