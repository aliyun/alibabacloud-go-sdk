// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExportLogStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExportLogStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExportLogStateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExportLogStateResponseBody) *UpdateExportLogStateResponse
	GetBody() *UpdateExportLogStateResponseBody
}

type UpdateExportLogStateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExportLogStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExportLogStateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExportLogStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateExportLogStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExportLogStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExportLogStateResponse) GetBody() *UpdateExportLogStateResponseBody {
	return s.Body
}

func (s *UpdateExportLogStateResponse) SetHeaders(v map[string]*string) *UpdateExportLogStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateExportLogStateResponse) SetStatusCode(v int32) *UpdateExportLogStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExportLogStateResponse) SetBody(v *UpdateExportLogStateResponseBody) *UpdateExportLogStateResponse {
	s.Body = v
	return s
}

func (s *UpdateExportLogStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
