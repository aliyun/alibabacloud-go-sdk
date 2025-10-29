// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataQualityScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataQualityScanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataQualityScanResponseBody) *UpdateDataQualityScanResponse
	GetBody() *UpdateDataQualityScanResponseBody
}

type UpdateDataQualityScanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataQualityScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataQualityScanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataQualityScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataQualityScanResponse) GetBody() *UpdateDataQualityScanResponseBody {
	return s.Body
}

func (s *UpdateDataQualityScanResponse) SetHeaders(v map[string]*string) *UpdateDataQualityScanResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataQualityScanResponse) SetStatusCode(v int32) *UpdateDataQualityScanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataQualityScanResponse) SetBody(v *UpdateDataQualityScanResponseBody) *UpdateDataQualityScanResponse {
	s.Body = v
	return s
}

func (s *UpdateDataQualityScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
