// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLindormV2InstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLindormV2InstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLindormV2InstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLindormV2InstanceResponseBody) *ModifyLindormV2InstanceResponse
	GetBody() *ModifyLindormV2InstanceResponseBody
}

type ModifyLindormV2InstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLindormV2InstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLindormV2InstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLindormV2InstanceResponse) GetBody() *ModifyLindormV2InstanceResponseBody {
	return s.Body
}

func (s *ModifyLindormV2InstanceResponse) SetHeaders(v map[string]*string) *ModifyLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyLindormV2InstanceResponse) SetStatusCode(v int32) *ModifyLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLindormV2InstanceResponse) SetBody(v *ModifyLindormV2InstanceResponseBody) *ModifyLindormV2InstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyLindormV2InstanceResponse) Validate() error {
	return dara.Validate(s)
}
