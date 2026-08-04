// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchSetMemberAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterBatchSetMemberAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterBatchSetMemberAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterBatchSetMemberAuthorizationResponseBody) *ModelRouterBatchSetMemberAuthorizationResponse
	GetBody() *ModelRouterBatchSetMemberAuthorizationResponseBody
}

type ModelRouterBatchSetMemberAuthorizationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterBatchSetMemberAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterBatchSetMemberAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchSetMemberAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchSetMemberAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterBatchSetMemberAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterBatchSetMemberAuthorizationResponse) GetBody() *ModelRouterBatchSetMemberAuthorizationResponseBody {
	return s.Body
}

func (s *ModelRouterBatchSetMemberAuthorizationResponse) SetHeaders(v map[string]*string) *ModelRouterBatchSetMemberAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponse) SetStatusCode(v int32) *ModelRouterBatchSetMemberAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponse) SetBody(v *ModelRouterBatchSetMemberAuthorizationResponseBody) *ModelRouterBatchSetMemberAuthorizationResponse {
	s.Body = v
	return s
}

func (s *ModelRouterBatchSetMemberAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
