// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchRenewMemberApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterBatchRenewMemberApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterBatchRenewMemberApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterBatchRenewMemberApiKeysResponseBody) *ModelRouterBatchRenewMemberApiKeysResponse
	GetBody() *ModelRouterBatchRenewMemberApiKeysResponseBody
}

type ModelRouterBatchRenewMemberApiKeysResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterBatchRenewMemberApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterBatchRenewMemberApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchRenewMemberApiKeysResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchRenewMemberApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterBatchRenewMemberApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterBatchRenewMemberApiKeysResponse) GetBody() *ModelRouterBatchRenewMemberApiKeysResponseBody {
	return s.Body
}

func (s *ModelRouterBatchRenewMemberApiKeysResponse) SetHeaders(v map[string]*string) *ModelRouterBatchRenewMemberApiKeysResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponse) SetStatusCode(v int32) *ModelRouterBatchRenewMemberApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponse) SetBody(v *ModelRouterBatchRenewMemberApiKeysResponseBody) *ModelRouterBatchRenewMemberApiKeysResponse {
	s.Body = v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
