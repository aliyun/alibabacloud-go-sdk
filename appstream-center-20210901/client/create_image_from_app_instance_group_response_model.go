// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageFromAppInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageFromAppInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageFromAppInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageFromAppInstanceGroupResponseBody) *CreateImageFromAppInstanceGroupResponse
	GetBody() *CreateImageFromAppInstanceGroupResponseBody
}

type CreateImageFromAppInstanceGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageFromAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageFromAppInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageFromAppInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageFromAppInstanceGroupResponse) GetBody() *CreateImageFromAppInstanceGroupResponseBody {
	return s.Body
}

func (s *CreateImageFromAppInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateImageFromAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponse) SetStatusCode(v int32) *CreateImageFromAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponse) SetBody(v *CreateImageFromAppInstanceGroupResponseBody) *CreateImageFromAppInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponse) Validate() error {
	return dara.Validate(s)
}
