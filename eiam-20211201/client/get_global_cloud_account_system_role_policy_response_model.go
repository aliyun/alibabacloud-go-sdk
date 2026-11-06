// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGlobalCloudAccountSystemRolePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGlobalCloudAccountSystemRolePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGlobalCloudAccountSystemRolePolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetGlobalCloudAccountSystemRolePolicyResponseBody) *GetGlobalCloudAccountSystemRolePolicyResponse
	GetBody() *GetGlobalCloudAccountSystemRolePolicyResponseBody
}

type GetGlobalCloudAccountSystemRolePolicyResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGlobalCloudAccountSystemRolePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGlobalCloudAccountSystemRolePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalCloudAccountSystemRolePolicyResponse) GoString() string {
	return s.String()
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponse) GetBody() *GetGlobalCloudAccountSystemRolePolicyResponseBody {
	return s.Body
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponse) SetHeaders(v map[string]*string) *GetGlobalCloudAccountSystemRolePolicyResponse {
	s.Headers = v
	return s
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponse) SetStatusCode(v int32) *GetGlobalCloudAccountSystemRolePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponse) SetBody(v *GetGlobalCloudAccountSystemRolePolicyResponseBody) *GetGlobalCloudAccountSystemRolePolicyResponse {
	s.Body = v
	return s
}

func (s *GetGlobalCloudAccountSystemRolePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
