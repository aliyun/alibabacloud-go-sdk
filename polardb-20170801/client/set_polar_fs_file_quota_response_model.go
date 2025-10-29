// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolarFsFileQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPolarFsFileQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPolarFsFileQuotaResponse
	GetStatusCode() *int32
	SetBody(v *SetPolarFsFileQuotaResponseBody) *SetPolarFsFileQuotaResponse
	GetBody() *SetPolarFsFileQuotaResponseBody
}

type SetPolarFsFileQuotaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPolarFsFileQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPolarFsFileQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPolarFsFileQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetPolarFsFileQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPolarFsFileQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPolarFsFileQuotaResponse) GetBody() *SetPolarFsFileQuotaResponseBody {
	return s.Body
}

func (s *SetPolarFsFileQuotaResponse) SetHeaders(v map[string]*string) *SetPolarFsFileQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetPolarFsFileQuotaResponse) SetStatusCode(v int32) *SetPolarFsFileQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPolarFsFileQuotaResponse) SetBody(v *SetPolarFsFileQuotaResponseBody) *SetPolarFsFileQuotaResponse {
	s.Body = v
	return s
}

func (s *SetPolarFsFileQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
