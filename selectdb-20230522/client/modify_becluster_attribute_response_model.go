// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBEClusterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBEClusterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBEClusterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBEClusterAttributeResponseBody) *ModifyBEClusterAttributeResponse
	GetBody() *ModifyBEClusterAttributeResponseBody
}

type ModifyBEClusterAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBEClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBEClusterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBEClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyBEClusterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBEClusterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBEClusterAttributeResponse) GetBody() *ModifyBEClusterAttributeResponseBody {
	return s.Body
}

func (s *ModifyBEClusterAttributeResponse) SetHeaders(v map[string]*string) *ModifyBEClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyBEClusterAttributeResponse) SetStatusCode(v int32) *ModifyBEClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBEClusterAttributeResponse) SetBody(v *ModifyBEClusterAttributeResponseBody) *ModifyBEClusterAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyBEClusterAttributeResponse) Validate() error {
	return dara.Validate(s)
}
