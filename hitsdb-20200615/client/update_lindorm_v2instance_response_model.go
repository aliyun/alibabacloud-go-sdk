// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2InstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLindormV2InstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLindormV2InstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLindormV2InstanceResponseBody) *UpdateLindormV2InstanceResponse
	GetBody() *UpdateLindormV2InstanceResponseBody
}

type UpdateLindormV2InstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLindormV2InstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLindormV2InstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLindormV2InstanceResponse) GetBody() *UpdateLindormV2InstanceResponseBody {
	return s.Body
}

func (s *UpdateLindormV2InstanceResponse) SetHeaders(v map[string]*string) *UpdateLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateLindormV2InstanceResponse) SetStatusCode(v int32) *UpdateLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLindormV2InstanceResponse) SetBody(v *UpdateLindormV2InstanceResponseBody) *UpdateLindormV2InstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateLindormV2InstanceResponse) Validate() error {
	return dara.Validate(s)
}
