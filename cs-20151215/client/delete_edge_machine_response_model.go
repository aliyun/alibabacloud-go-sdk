// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEdgeMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEdgeMachineResponse
	GetStatusCode() *int32
}

type DeleteEdgeMachineResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteEdgeMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeMachineResponse) GoString() string {
	return s.String()
}

func (s *DeleteEdgeMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEdgeMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEdgeMachineResponse) SetHeaders(v map[string]*string) *DeleteEdgeMachineResponse {
	s.Headers = v
	return s
}

func (s *DeleteEdgeMachineResponse) SetStatusCode(v int32) *DeleteEdgeMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEdgeMachineResponse) Validate() error {
	return dara.Validate(s)
}
