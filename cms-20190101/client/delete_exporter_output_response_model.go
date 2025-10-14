// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExporterOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExporterOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExporterOutputResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExporterOutputResponseBody) *DeleteExporterOutputResponse
	GetBody() *DeleteExporterOutputResponseBody
}

type DeleteExporterOutputResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExporterOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExporterOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExporterOutputResponse) GoString() string {
	return s.String()
}

func (s *DeleteExporterOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExporterOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExporterOutputResponse) GetBody() *DeleteExporterOutputResponseBody {
	return s.Body
}

func (s *DeleteExporterOutputResponse) SetHeaders(v map[string]*string) *DeleteExporterOutputResponse {
	s.Headers = v
	return s
}

func (s *DeleteExporterOutputResponse) SetStatusCode(v int32) *DeleteExporterOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExporterOutputResponse) SetBody(v *DeleteExporterOutputResponseBody) *DeleteExporterOutputResponse {
	s.Body = v
	return s
}

func (s *DeleteExporterOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
