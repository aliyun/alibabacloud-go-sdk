// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupResponseBody) *GetResourceGroupResponse
	GetBody() *GetResourceGroupResponseBody
}

type GetResourceGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupResponse) GetBody() *GetResourceGroupResponseBody {
	return s.Body
}

func (s *GetResourceGroupResponse) SetHeaders(v map[string]*string) *GetResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupResponse) SetStatusCode(v int32) *GetResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupResponse) SetBody(v *GetResourceGroupResponseBody) *GetResourceGroupResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
