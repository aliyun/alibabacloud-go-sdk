// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutExporterOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutExporterOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutExporterOutputResponse
	GetStatusCode() *int32
	SetBody(v *PutExporterOutputResponseBody) *PutExporterOutputResponse
	GetBody() *PutExporterOutputResponseBody
}

type PutExporterOutputResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutExporterOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutExporterOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s PutExporterOutputResponse) GoString() string {
	return s.String()
}

func (s *PutExporterOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutExporterOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutExporterOutputResponse) GetBody() *PutExporterOutputResponseBody {
	return s.Body
}

func (s *PutExporterOutputResponse) SetHeaders(v map[string]*string) *PutExporterOutputResponse {
	s.Headers = v
	return s
}

func (s *PutExporterOutputResponse) SetStatusCode(v int32) *PutExporterOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *PutExporterOutputResponse) SetBody(v *PutExporterOutputResponseBody) *PutExporterOutputResponse {
	s.Body = v
	return s
}

func (s *PutExporterOutputResponse) Validate() error {
	return dara.Validate(s)
}
