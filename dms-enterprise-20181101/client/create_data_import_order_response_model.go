// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataImportOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataImportOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataImportOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataImportOrderResponseBody) *CreateDataImportOrderResponse
	GetBody() *CreateDataImportOrderResponseBody
}

type CreateDataImportOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataImportOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataImportOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataImportOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataImportOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataImportOrderResponse) GetBody() *CreateDataImportOrderResponseBody {
	return s.Body
}

func (s *CreateDataImportOrderResponse) SetHeaders(v map[string]*string) *CreateDataImportOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataImportOrderResponse) SetStatusCode(v int32) *CreateDataImportOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataImportOrderResponse) SetBody(v *CreateDataImportOrderResponseBody) *CreateDataImportOrderResponse {
	s.Body = v
	return s
}

func (s *CreateDataImportOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
