// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMqGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMqGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMqGroupResponse
	GetStatusCode() *int32
}

type DeleteMqGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteMqGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMqGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMqGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMqGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMqGroupResponse) SetHeaders(v map[string]*string) *DeleteMqGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMqGroupResponse) SetStatusCode(v int32) *DeleteMqGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMqGroupResponse) Validate() error {
	return dara.Validate(s)
}
