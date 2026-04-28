// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveGroupMemberResponse
	GetStatusCode() *int32
}

type RemoveGroupMemberResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveGroupMemberResponse) SetHeaders(v map[string]*string) *RemoveGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupMemberResponse) SetStatusCode(v int32) *RemoveGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveGroupMemberResponse) Validate() error {
	return dara.Validate(s)
}
