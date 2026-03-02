// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePdpServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePdpServiceGroupResponse
	GetStatusCode() *int32
}

type DeletePdpServiceGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeletePdpServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeletePdpServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePdpServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePdpServiceGroupResponse) SetHeaders(v map[string]*string) *DeletePdpServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeletePdpServiceGroupResponse) SetStatusCode(v int32) *DeletePdpServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePdpServiceGroupResponse) Validate() error {
	return dara.Validate(s)
}
