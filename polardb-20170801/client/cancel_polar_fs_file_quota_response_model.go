// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPolarFsFileQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPolarFsFileQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPolarFsFileQuotaResponse
	GetStatusCode() *int32
	SetBody(v *CancelPolarFsFileQuotaResponseBody) *CancelPolarFsFileQuotaResponse
	GetBody() *CancelPolarFsFileQuotaResponseBody
}

type CancelPolarFsFileQuotaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPolarFsFileQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPolarFsFileQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPolarFsFileQuotaResponse) GoString() string {
	return s.String()
}

func (s *CancelPolarFsFileQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPolarFsFileQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPolarFsFileQuotaResponse) GetBody() *CancelPolarFsFileQuotaResponseBody {
	return s.Body
}

func (s *CancelPolarFsFileQuotaResponse) SetHeaders(v map[string]*string) *CancelPolarFsFileQuotaResponse {
	s.Headers = v
	return s
}

func (s *CancelPolarFsFileQuotaResponse) SetStatusCode(v int32) *CancelPolarFsFileQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPolarFsFileQuotaResponse) SetBody(v *CancelPolarFsFileQuotaResponseBody) *CancelPolarFsFileQuotaResponse {
	s.Body = v
	return s
}

func (s *CancelPolarFsFileQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
