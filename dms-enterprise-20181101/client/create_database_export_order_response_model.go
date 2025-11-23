// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseExportOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatabaseExportOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatabaseExportOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatabaseExportOrderResponseBody) *CreateDatabaseExportOrderResponse
	GetBody() *CreateDatabaseExportOrderResponseBody
}

type CreateDatabaseExportOrderResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatabaseExportOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatabaseExportOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseExportOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseExportOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatabaseExportOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatabaseExportOrderResponse) GetBody() *CreateDatabaseExportOrderResponseBody {
	return s.Body
}

func (s *CreateDatabaseExportOrderResponse) SetHeaders(v map[string]*string) *CreateDatabaseExportOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDatabaseExportOrderResponse) SetStatusCode(v int32) *CreateDatabaseExportOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatabaseExportOrderResponse) SetBody(v *CreateDatabaseExportOrderResponseBody) *CreateDatabaseExportOrderResponse {
	s.Body = v
	return s
}

func (s *CreateDatabaseExportOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
