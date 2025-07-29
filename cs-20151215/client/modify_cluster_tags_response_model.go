// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterTagsResponse
	GetStatusCode() *int32
}

type ModifyClusterTagsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ModifyClusterTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterTagsResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterTagsResponse) SetHeaders(v map[string]*string) *ModifyClusterTagsResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterTagsResponse) SetStatusCode(v int32) *ModifyClusterTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterTagsResponse) Validate() error {
	return dara.Validate(s)
}
