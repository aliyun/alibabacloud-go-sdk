// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertUmodelCommonSchemaRefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertUmodelCommonSchemaRefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertUmodelCommonSchemaRefResponse
	GetStatusCode() *int32
	SetBody(v *UpsertUmodelCommonSchemaRefResponseBody) *UpsertUmodelCommonSchemaRefResponse
	GetBody() *UpsertUmodelCommonSchemaRefResponseBody
}

type UpsertUmodelCommonSchemaRefResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertUmodelCommonSchemaRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertUmodelCommonSchemaRefResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertUmodelCommonSchemaRefResponse) GoString() string {
	return s.String()
}

func (s *UpsertUmodelCommonSchemaRefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertUmodelCommonSchemaRefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertUmodelCommonSchemaRefResponse) GetBody() *UpsertUmodelCommonSchemaRefResponseBody {
	return s.Body
}

func (s *UpsertUmodelCommonSchemaRefResponse) SetHeaders(v map[string]*string) *UpsertUmodelCommonSchemaRefResponse {
	s.Headers = v
	return s
}

func (s *UpsertUmodelCommonSchemaRefResponse) SetStatusCode(v int32) *UpsertUmodelCommonSchemaRefResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertUmodelCommonSchemaRefResponse) SetBody(v *UpsertUmodelCommonSchemaRefResponseBody) *UpsertUmodelCommonSchemaRefResponse {
	s.Body = v
	return s
}

func (s *UpsertUmodelCommonSchemaRefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
