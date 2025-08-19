// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProvisionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProvisionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProvisionConfigResponse
	GetStatusCode() *int32
}

type DeleteProvisionConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteProvisionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProvisionConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteProvisionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProvisionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProvisionConfigResponse) SetHeaders(v map[string]*string) *DeleteProvisionConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteProvisionConfigResponse) SetStatusCode(v int32) *DeleteProvisionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProvisionConfigResponse) Validate() error {
	return dara.Validate(s)
}
