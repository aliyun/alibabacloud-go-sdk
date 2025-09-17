// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReplicationJobAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyReplicationJobAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyReplicationJobAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyReplicationJobAttributeResponseBody) *ModifyReplicationJobAttributeResponse
	GetBody() *ModifyReplicationJobAttributeResponseBody
}

type ModifyReplicationJobAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReplicationJobAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReplicationJobAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyReplicationJobAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyReplicationJobAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyReplicationJobAttributeResponse) GetBody() *ModifyReplicationJobAttributeResponseBody {
	return s.Body
}

func (s *ModifyReplicationJobAttributeResponse) SetHeaders(v map[string]*string) *ModifyReplicationJobAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyReplicationJobAttributeResponse) SetStatusCode(v int32) *ModifyReplicationJobAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReplicationJobAttributeResponse) SetBody(v *ModifyReplicationJobAttributeResponseBody) *ModifyReplicationJobAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyReplicationJobAttributeResponse) Validate() error {
	return dara.Validate(s)
}
