// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataExportOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataExportOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataExportOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataExportOrderResponseBody) *CreateDataExportOrderResponse
	GetBody() *CreateDataExportOrderResponseBody
}

type CreateDataExportOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataExportOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataExportOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataExportOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataExportOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataExportOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataExportOrderResponse) GetBody() *CreateDataExportOrderResponseBody {
	return s.Body
}

func (s *CreateDataExportOrderResponse) SetHeaders(v map[string]*string) *CreateDataExportOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataExportOrderResponse) SetStatusCode(v int32) *CreateDataExportOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataExportOrderResponse) SetBody(v *CreateDataExportOrderResponseBody) *CreateDataExportOrderResponse {
	s.Body = v
	return s
}

func (s *CreateDataExportOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
