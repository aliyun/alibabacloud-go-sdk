// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorGroupMemberResponse
	GetStatusCode() *int32
}

type DeleteMonitorGroupMemberResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteMonitorGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorGroupMemberResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupMemberResponse) SetStatusCode(v int32) *DeleteMonitorGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorGroupMemberResponse) Validate() error {
	return dara.Validate(s)
}
