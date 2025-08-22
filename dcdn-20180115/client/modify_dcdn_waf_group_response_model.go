// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDcdnWafGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDcdnWafGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDcdnWafGroupResponseBody) *ModifyDcdnWafGroupResponse
	GetBody() *ModifyDcdnWafGroupResponseBody
}

type ModifyDcdnWafGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDcdnWafGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDcdnWafGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDcdnWafGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDcdnWafGroupResponse) GetBody() *ModifyDcdnWafGroupResponseBody {
	return s.Body
}

func (s *ModifyDcdnWafGroupResponse) SetHeaders(v map[string]*string) *ModifyDcdnWafGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDcdnWafGroupResponse) SetStatusCode(v int32) *ModifyDcdnWafGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDcdnWafGroupResponse) SetBody(v *ModifyDcdnWafGroupResponseBody) *ModifyDcdnWafGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyDcdnWafGroupResponse) Validate() error {
	return dara.Validate(s)
}
