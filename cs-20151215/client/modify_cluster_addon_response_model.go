// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterAddonResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterAddonResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterAddonResponse
	GetStatusCode() *int32
}

type ModifyClusterAddonResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ModifyClusterAddonResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAddonResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterAddonResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterAddonResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterAddonResponse) SetHeaders(v map[string]*string) *ModifyClusterAddonResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterAddonResponse) SetStatusCode(v int32) *ModifyClusterAddonResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterAddonResponse) Validate() error {
	return dara.Validate(s)
}
