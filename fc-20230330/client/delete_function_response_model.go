// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFunctionResponse
	GetStatusCode() *int32
}

type DeleteFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResponse) SetStatusCode(v int32) *DeleteFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionResponse) Validate() error {
	return dara.Validate(s)
}
