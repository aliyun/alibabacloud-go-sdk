// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelCommonSchemaRefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUmodelCommonSchemaRefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUmodelCommonSchemaRefResponse
	GetStatusCode() *int32
	SetBody(v *GetUmodelCommonSchemaRefResponseBody) *GetUmodelCommonSchemaRefResponse
	GetBody() *GetUmodelCommonSchemaRefResponseBody
}

type GetUmodelCommonSchemaRefResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUmodelCommonSchemaRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUmodelCommonSchemaRefResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelCommonSchemaRefResponse) GoString() string {
	return s.String()
}

func (s *GetUmodelCommonSchemaRefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUmodelCommonSchemaRefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUmodelCommonSchemaRefResponse) GetBody() *GetUmodelCommonSchemaRefResponseBody {
	return s.Body
}

func (s *GetUmodelCommonSchemaRefResponse) SetHeaders(v map[string]*string) *GetUmodelCommonSchemaRefResponse {
	s.Headers = v
	return s
}

func (s *GetUmodelCommonSchemaRefResponse) SetStatusCode(v int32) *GetUmodelCommonSchemaRefResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUmodelCommonSchemaRefResponse) SetBody(v *GetUmodelCommonSchemaRefResponseBody) *GetUmodelCommonSchemaRefResponse {
	s.Body = v
	return s
}

func (s *GetUmodelCommonSchemaRefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
