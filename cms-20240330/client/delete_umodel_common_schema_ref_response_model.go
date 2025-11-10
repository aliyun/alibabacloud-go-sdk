// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelCommonSchemaRefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUmodelCommonSchemaRefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUmodelCommonSchemaRefResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUmodelCommonSchemaRefResponseBody) *DeleteUmodelCommonSchemaRefResponse
	GetBody() *DeleteUmodelCommonSchemaRefResponseBody
}

type DeleteUmodelCommonSchemaRefResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUmodelCommonSchemaRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUmodelCommonSchemaRefResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelCommonSchemaRefResponse) GoString() string {
	return s.String()
}

func (s *DeleteUmodelCommonSchemaRefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUmodelCommonSchemaRefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUmodelCommonSchemaRefResponse) GetBody() *DeleteUmodelCommonSchemaRefResponseBody {
	return s.Body
}

func (s *DeleteUmodelCommonSchemaRefResponse) SetHeaders(v map[string]*string) *DeleteUmodelCommonSchemaRefResponse {
	s.Headers = v
	return s
}

func (s *DeleteUmodelCommonSchemaRefResponse) SetStatusCode(v int32) *DeleteUmodelCommonSchemaRefResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUmodelCommonSchemaRefResponse) SetBody(v *DeleteUmodelCommonSchemaRefResponseBody) *DeleteUmodelCommonSchemaRefResponse {
	s.Body = v
	return s
}

func (s *DeleteUmodelCommonSchemaRefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
