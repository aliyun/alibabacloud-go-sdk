// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchDisableMemberApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterBatchDisableMemberApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterBatchDisableMemberApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterBatchDisableMemberApiKeysResponseBody) *ModelRouterBatchDisableMemberApiKeysResponse
	GetBody() *ModelRouterBatchDisableMemberApiKeysResponseBody
}

type ModelRouterBatchDisableMemberApiKeysResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterBatchDisableMemberApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterBatchDisableMemberApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchDisableMemberApiKeysResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchDisableMemberApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterBatchDisableMemberApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterBatchDisableMemberApiKeysResponse) GetBody() *ModelRouterBatchDisableMemberApiKeysResponseBody {
	return s.Body
}

func (s *ModelRouterBatchDisableMemberApiKeysResponse) SetHeaders(v map[string]*string) *ModelRouterBatchDisableMemberApiKeysResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponse) SetStatusCode(v int32) *ModelRouterBatchDisableMemberApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponse) SetBody(v *ModelRouterBatchDisableMemberApiKeysResponseBody) *ModelRouterBatchDisableMemberApiKeysResponse {
	s.Body = v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
