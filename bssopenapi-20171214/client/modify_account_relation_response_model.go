// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountRelationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountRelationResponseBody) *ModifyAccountRelationResponse
	GetBody() *ModifyAccountRelationResponseBody
}

type ModifyAccountRelationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountRelationResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountRelationResponse) GetBody() *ModifyAccountRelationResponseBody {
	return s.Body
}

func (s *ModifyAccountRelationResponse) SetHeaders(v map[string]*string) *ModifyAccountRelationResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountRelationResponse) SetStatusCode(v int32) *ModifyAccountRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountRelationResponse) SetBody(v *ModifyAccountRelationResponseBody) *ModifyAccountRelationResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountRelationResponse) Validate() error {
	return dara.Validate(s)
}
