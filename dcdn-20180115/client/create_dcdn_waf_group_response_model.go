// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnWafGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDcdnWafGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDcdnWafGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateDcdnWafGroupResponseBody) *CreateDcdnWafGroupResponse
	GetBody() *CreateDcdnWafGroupResponseBody
}

type CreateDcdnWafGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDcdnWafGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDcdnWafGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnWafGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnWafGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDcdnWafGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDcdnWafGroupResponse) GetBody() *CreateDcdnWafGroupResponseBody {
	return s.Body
}

func (s *CreateDcdnWafGroupResponse) SetHeaders(v map[string]*string) *CreateDcdnWafGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnWafGroupResponse) SetStatusCode(v int32) *CreateDcdnWafGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDcdnWafGroupResponse) SetBody(v *CreateDcdnWafGroupResponseBody) *CreateDcdnWafGroupResponse {
	s.Body = v
	return s
}

func (s *CreateDcdnWafGroupResponse) Validate() error {
	return dara.Validate(s)
}
