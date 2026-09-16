// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeUmodelCommonSchemaRefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody) *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse
	GetBody() *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody
}

type DeleteDigitalEmployeeUmodelCommonSchemaRefResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) GetBody() *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody {
	return s.Body
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) SetHeaders(v map[string]*string) *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse {
	s.Headers = v
	return s
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) SetStatusCode(v int32) *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) SetBody(v *DeleteDigitalEmployeeUmodelCommonSchemaRefResponseBody) *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse {
	s.Body = v
	return s
}

func (s *DeleteDigitalEmployeeUmodelCommonSchemaRefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
