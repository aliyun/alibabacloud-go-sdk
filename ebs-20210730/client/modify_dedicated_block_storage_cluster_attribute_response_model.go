// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedBlockStorageClusterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDedicatedBlockStorageClusterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDedicatedBlockStorageClusterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDedicatedBlockStorageClusterAttributeResponseBody) *ModifyDedicatedBlockStorageClusterAttributeResponse
	GetBody() *ModifyDedicatedBlockStorageClusterAttributeResponseBody
}

type ModifyDedicatedBlockStorageClusterAttributeResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDedicatedBlockStorageClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) GetBody() *ModifyDedicatedBlockStorageClusterAttributeResponseBody {
	return s.Body
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetStatusCode(v int32) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetBody(v *ModifyDedicatedBlockStorageClusterAttributeResponseBody) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
