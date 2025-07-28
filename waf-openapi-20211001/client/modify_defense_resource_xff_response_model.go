// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseResourceXffResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseResourceXffResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseResourceXffResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseResourceXffResponseBody) *ModifyDefenseResourceXffResponse
	GetBody() *ModifyDefenseResourceXffResponseBody
}

type ModifyDefenseResourceXffResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseResourceXffResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseResourceXffResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceXffResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceXffResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseResourceXffResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseResourceXffResponse) GetBody() *ModifyDefenseResourceXffResponseBody {
	return s.Body
}

func (s *ModifyDefenseResourceXffResponse) SetHeaders(v map[string]*string) *ModifyDefenseResourceXffResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseResourceXffResponse) SetStatusCode(v int32) *ModifyDefenseResourceXffResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseResourceXffResponse) SetBody(v *ModifyDefenseResourceXffResponseBody) *ModifyDefenseResourceXffResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseResourceXffResponse) Validate() error {
	return dara.Validate(s)
}
