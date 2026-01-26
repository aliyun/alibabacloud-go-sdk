// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCmsExporterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCmsExporterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCmsExporterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCmsExporterResponseBody) *DeleteCmsExporterResponse
	GetBody() *DeleteCmsExporterResponseBody
}

type DeleteCmsExporterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCmsExporterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCmsExporterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCmsExporterResponse) GoString() string {
	return s.String()
}

func (s *DeleteCmsExporterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCmsExporterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCmsExporterResponse) GetBody() *DeleteCmsExporterResponseBody {
	return s.Body
}

func (s *DeleteCmsExporterResponse) SetHeaders(v map[string]*string) *DeleteCmsExporterResponse {
	s.Headers = v
	return s
}

func (s *DeleteCmsExporterResponse) SetStatusCode(v int32) *DeleteCmsExporterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCmsExporterResponse) SetBody(v *DeleteCmsExporterResponseBody) *DeleteCmsExporterResponse {
	s.Body = v
	return s
}

func (s *DeleteCmsExporterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
