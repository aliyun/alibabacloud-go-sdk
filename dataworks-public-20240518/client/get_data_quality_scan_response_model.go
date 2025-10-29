// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityScanResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityScanResponseBody) *GetDataQualityScanResponse
	GetBody() *GetDataQualityScanResponseBody
}

type GetDataQualityScanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityScanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityScanResponse) GetBody() *GetDataQualityScanResponseBody {
	return s.Body
}

func (s *GetDataQualityScanResponse) SetHeaders(v map[string]*string) *GetDataQualityScanResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityScanResponse) SetStatusCode(v int32) *GetDataQualityScanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityScanResponse) SetBody(v *GetDataQualityScanResponseBody) *GetDataQualityScanResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
