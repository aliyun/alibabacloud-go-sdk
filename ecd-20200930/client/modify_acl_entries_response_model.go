// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAclEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAclEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAclEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAclEntriesResponseBody) *ModifyAclEntriesResponse
	GetBody() *ModifyAclEntriesResponseBody
}

type ModifyAclEntriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAclEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAclEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAclEntriesResponse) GoString() string {
	return s.String()
}

func (s *ModifyAclEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAclEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAclEntriesResponse) GetBody() *ModifyAclEntriesResponseBody {
	return s.Body
}

func (s *ModifyAclEntriesResponse) SetHeaders(v map[string]*string) *ModifyAclEntriesResponse {
	s.Headers = v
	return s
}

func (s *ModifyAclEntriesResponse) SetStatusCode(v int32) *ModifyAclEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAclEntriesResponse) SetBody(v *ModifyAclEntriesResponseBody) *ModifyAclEntriesResponse {
	s.Body = v
	return s
}

func (s *ModifyAclEntriesResponse) Validate() error {
	return dara.Validate(s)
}
