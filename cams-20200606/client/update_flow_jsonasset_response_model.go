// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowJSONAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowJSONAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowJSONAssetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFlowJSONAssetResponseBody) *UpdateFlowJSONAssetResponse
	GetBody() *UpdateFlowJSONAssetResponseBody
}

type UpdateFlowJSONAssetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlowJSONAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowJSONAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowJSONAssetResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowJSONAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowJSONAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowJSONAssetResponse) GetBody() *UpdateFlowJSONAssetResponseBody {
	return s.Body
}

func (s *UpdateFlowJSONAssetResponse) SetHeaders(v map[string]*string) *UpdateFlowJSONAssetResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowJSONAssetResponse) SetStatusCode(v int32) *UpdateFlowJSONAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowJSONAssetResponse) SetBody(v *UpdateFlowJSONAssetResponseBody) *UpdateFlowJSONAssetResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowJSONAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
