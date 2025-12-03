// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PatchGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PatchGroupResponse
	GetStatusCode() *int32
}

type PatchGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PatchGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s PatchGroupResponse) GoString() string {
	return s.String()
}

func (s *PatchGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PatchGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PatchGroupResponse) SetHeaders(v map[string]*string) *PatchGroupResponse {
	s.Headers = v
	return s
}

func (s *PatchGroupResponse) SetStatusCode(v int32) *PatchGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *PatchGroupResponse) Validate() error {
	return dara.Validate(s)
}
