// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFeatureConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFeatureConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFeatureConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFeatureConfigResponseBody) *ModifyFeatureConfigResponse
	GetBody() *ModifyFeatureConfigResponseBody
}

type ModifyFeatureConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFeatureConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFeatureConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFeatureConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyFeatureConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFeatureConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFeatureConfigResponse) GetBody() *ModifyFeatureConfigResponseBody {
	return s.Body
}

func (s *ModifyFeatureConfigResponse) SetHeaders(v map[string]*string) *ModifyFeatureConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyFeatureConfigResponse) SetStatusCode(v int32) *ModifyFeatureConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFeatureConfigResponse) SetBody(v *ModifyFeatureConfigResponseBody) *ModifyFeatureConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyFeatureConfigResponse) Validate() error {
	return dara.Validate(s)
}
