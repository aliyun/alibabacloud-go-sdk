// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataQualityScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataQualityScanResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataQualityScanResponseBody) *CreateDataQualityScanResponse
	GetBody() *CreateDataQualityScanResponseBody
}

type CreateDataQualityScanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataQualityScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataQualityScanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanResponse) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataQualityScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataQualityScanResponse) GetBody() *CreateDataQualityScanResponseBody {
	return s.Body
}

func (s *CreateDataQualityScanResponse) SetHeaders(v map[string]*string) *CreateDataQualityScanResponse {
	s.Headers = v
	return s
}

func (s *CreateDataQualityScanResponse) SetStatusCode(v int32) *CreateDataQualityScanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataQualityScanResponse) SetBody(v *CreateDataQualityScanResponseBody) *CreateDataQualityScanResponse {
	s.Body = v
	return s
}

func (s *CreateDataQualityScanResponse) Validate() error {
	return dara.Validate(s)
}
