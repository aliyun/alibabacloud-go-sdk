// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceGroupResponseBody) *CreateResourceGroupResponse
	GetBody() *CreateResourceGroupResponseBody
}

type CreateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceGroupResponse) GetBody() *CreateResourceGroupResponseBody {
	return s.Body
}

func (s *CreateResourceGroupResponse) SetHeaders(v map[string]*string) *CreateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceGroupResponse) SetStatusCode(v int32) *CreateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceGroupResponse) SetBody(v *CreateResourceGroupResponseBody) *CreateResourceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
