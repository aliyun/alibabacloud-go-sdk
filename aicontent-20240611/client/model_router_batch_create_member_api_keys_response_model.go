// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchCreateMemberApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterBatchCreateMemberApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterBatchCreateMemberApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterBatchCreateMemberApiKeysResponseBody) *ModelRouterBatchCreateMemberApiKeysResponse
	GetBody() *ModelRouterBatchCreateMemberApiKeysResponseBody
}

type ModelRouterBatchCreateMemberApiKeysResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterBatchCreateMemberApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterBatchCreateMemberApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchCreateMemberApiKeysResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchCreateMemberApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterBatchCreateMemberApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterBatchCreateMemberApiKeysResponse) GetBody() *ModelRouterBatchCreateMemberApiKeysResponseBody {
	return s.Body
}

func (s *ModelRouterBatchCreateMemberApiKeysResponse) SetHeaders(v map[string]*string) *ModelRouterBatchCreateMemberApiKeysResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponse) SetStatusCode(v int32) *ModelRouterBatchCreateMemberApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponse) SetBody(v *ModelRouterBatchCreateMemberApiKeysResponseBody) *ModelRouterBatchCreateMemberApiKeysResponse {
	s.Body = v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
