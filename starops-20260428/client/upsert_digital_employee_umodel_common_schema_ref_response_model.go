// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertDigitalEmployeeUmodelCommonSchemaRefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse
	GetStatusCode() *int32
	SetBody(v *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody) *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse
	GetBody() *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody
}

type UpsertDigitalEmployeeUmodelCommonSchemaRefResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) GoString() string {
	return s.String()
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) GetBody() *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody {
	return s.Body
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) SetHeaders(v map[string]*string) *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse {
	s.Headers = v
	return s
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) SetStatusCode(v int32) *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) SetBody(v *UpsertDigitalEmployeeUmodelCommonSchemaRefResponseBody) *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse {
	s.Body = v
	return s
}

func (s *UpsertDigitalEmployeeUmodelCommonSchemaRefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
