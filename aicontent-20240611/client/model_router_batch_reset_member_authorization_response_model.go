// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchResetMemberAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterBatchResetMemberAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterBatchResetMemberAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterBatchResetMemberAuthorizationResponseBody) *ModelRouterBatchResetMemberAuthorizationResponse
	GetBody() *ModelRouterBatchResetMemberAuthorizationResponseBody
}

type ModelRouterBatchResetMemberAuthorizationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterBatchResetMemberAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterBatchResetMemberAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchResetMemberAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchResetMemberAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterBatchResetMemberAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterBatchResetMemberAuthorizationResponse) GetBody() *ModelRouterBatchResetMemberAuthorizationResponseBody {
	return s.Body
}

func (s *ModelRouterBatchResetMemberAuthorizationResponse) SetHeaders(v map[string]*string) *ModelRouterBatchResetMemberAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponse) SetStatusCode(v int32) *ModelRouterBatchResetMemberAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponse) SetBody(v *ModelRouterBatchResetMemberAuthorizationResponseBody) *ModelRouterBatchResetMemberAuthorizationResponse {
	s.Body = v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
