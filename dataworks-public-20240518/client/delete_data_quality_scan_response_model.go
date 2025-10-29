// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataQualityScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataQualityScanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataQualityScanResponseBody) *DeleteDataQualityScanResponse
	GetBody() *DeleteDataQualityScanResponseBody
}

type DeleteDataQualityScanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataQualityScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataQualityScanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityScanResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataQualityScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataQualityScanResponse) GetBody() *DeleteDataQualityScanResponseBody {
	return s.Body
}

func (s *DeleteDataQualityScanResponse) SetHeaders(v map[string]*string) *DeleteDataQualityScanResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataQualityScanResponse) SetStatusCode(v int32) *DeleteDataQualityScanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataQualityScanResponse) SetBody(v *DeleteDataQualityScanResponseBody) *DeleteDataQualityScanResponse {
	s.Body = v
	return s
}

func (s *DeleteDataQualityScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
