// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnWafGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnWafGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnWafGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnWafGroupResponseBody) *DeleteDcdnWafGroupResponse
	GetBody() *DeleteDcdnWafGroupResponseBody
}

type DeleteDcdnWafGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnWafGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnWafGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnWafGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnWafGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnWafGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnWafGroupResponse) GetBody() *DeleteDcdnWafGroupResponseBody {
	return s.Body
}

func (s *DeleteDcdnWafGroupResponse) SetHeaders(v map[string]*string) *DeleteDcdnWafGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnWafGroupResponse) SetStatusCode(v int32) *DeleteDcdnWafGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnWafGroupResponse) SetBody(v *DeleteDcdnWafGroupResponseBody) *DeleteDcdnWafGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnWafGroupResponse) Validate() error {
	return dara.Validate(s)
}
