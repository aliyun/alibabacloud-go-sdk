// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetCleanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAssetCleanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAssetCleanConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAssetCleanConfigResponseBody) *ModifyAssetCleanConfigResponse
	GetBody() *ModifyAssetCleanConfigResponseBody
}

type ModifyAssetCleanConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAssetCleanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAssetCleanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetCleanConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAssetCleanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAssetCleanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAssetCleanConfigResponse) GetBody() *ModifyAssetCleanConfigResponseBody {
	return s.Body
}

func (s *ModifyAssetCleanConfigResponse) SetHeaders(v map[string]*string) *ModifyAssetCleanConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAssetCleanConfigResponse) SetStatusCode(v int32) *ModifyAssetCleanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAssetCleanConfigResponse) SetBody(v *ModifyAssetCleanConfigResponseBody) *ModifyAssetCleanConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAssetCleanConfigResponse) Validate() error {
	return dara.Validate(s)
}
