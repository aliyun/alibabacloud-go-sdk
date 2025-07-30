// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairKVCacheCustomInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTairKVCacheCustomInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTairKVCacheCustomInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTairKVCacheCustomInstanceAttributeResponseBody) *ModifyTairKVCacheCustomInstanceAttributeResponse
	GetBody() *ModifyTairKVCacheCustomInstanceAttributeResponseBody
}

type ModifyTairKVCacheCustomInstanceAttributeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTairKVCacheCustomInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTairKVCacheCustomInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairKVCacheCustomInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponse) GetBody() *ModifyTairKVCacheCustomInstanceAttributeResponseBody {
	return s.Body
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyTairKVCacheCustomInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponse) SetStatusCode(v int32) *ModifyTairKVCacheCustomInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponse) SetBody(v *ModifyTairKVCacheCustomInstanceAttributeResponseBody) *ModifyTairKVCacheCustomInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyTairKVCacheCustomInstanceAttributeResponse) Validate() error {
	return dara.Validate(s)
}
