// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearchConditionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySearchConditionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySearchConditionResponse
	GetStatusCode() *int32
	SetBody(v *ModifySearchConditionResponseBody) *ModifySearchConditionResponse
	GetBody() *ModifySearchConditionResponseBody
}

type ModifySearchConditionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySearchConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySearchConditionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySearchConditionResponse) GoString() string {
	return s.String()
}

func (s *ModifySearchConditionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySearchConditionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySearchConditionResponse) GetBody() *ModifySearchConditionResponseBody {
	return s.Body
}

func (s *ModifySearchConditionResponse) SetHeaders(v map[string]*string) *ModifySearchConditionResponse {
	s.Headers = v
	return s
}

func (s *ModifySearchConditionResponse) SetStatusCode(v int32) *ModifySearchConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySearchConditionResponse) SetBody(v *ModifySearchConditionResponseBody) *ModifySearchConditionResponse {
	s.Body = v
	return s
}

func (s *ModifySearchConditionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
