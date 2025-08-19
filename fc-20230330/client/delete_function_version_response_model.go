// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFunctionVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFunctionVersionResponse
	GetStatusCode() *int32
}

type DeleteFunctionVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteFunctionVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFunctionVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFunctionVersionResponse) SetHeaders(v map[string]*string) *DeleteFunctionVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionVersionResponse) SetStatusCode(v int32) *DeleteFunctionVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionVersionResponse) Validate() error {
	return dara.Validate(s)
}
