// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBizCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBizCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBizCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBizCategoryResponseBody) *ModifyBizCategoryResponse
	GetBody() *ModifyBizCategoryResponseBody
}

type ModifyBizCategoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBizCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBizCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBizCategoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyBizCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBizCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBizCategoryResponse) GetBody() *ModifyBizCategoryResponseBody {
	return s.Body
}

func (s *ModifyBizCategoryResponse) SetHeaders(v map[string]*string) *ModifyBizCategoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyBizCategoryResponse) SetStatusCode(v int32) *ModifyBizCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBizCategoryResponse) SetBody(v *ModifyBizCategoryResponseBody) *ModifyBizCategoryResponse {
	s.Body = v
	return s
}

func (s *ModifyBizCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
