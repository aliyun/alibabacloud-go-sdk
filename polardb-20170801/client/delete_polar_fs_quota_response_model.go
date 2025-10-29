// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarFsQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarFsQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarFsQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarFsQuotaResponseBody) *DeletePolarFsQuotaResponse
	GetBody() *DeletePolarFsQuotaResponseBody
}

type DeletePolarFsQuotaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarFsQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarFsQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarFsQuotaResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarFsQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarFsQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarFsQuotaResponse) GetBody() *DeletePolarFsQuotaResponseBody {
	return s.Body
}

func (s *DeletePolarFsQuotaResponse) SetHeaders(v map[string]*string) *DeletePolarFsQuotaResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarFsQuotaResponse) SetStatusCode(v int32) *DeletePolarFsQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarFsQuotaResponse) SetBody(v *DeletePolarFsQuotaResponseBody) *DeletePolarFsQuotaResponse {
	s.Body = v
	return s
}

func (s *DeletePolarFsQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
