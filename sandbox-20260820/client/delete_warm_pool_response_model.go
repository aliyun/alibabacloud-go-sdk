// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarmPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWarmPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWarmPoolResponse
	GetStatusCode() *int32
}

type DeleteWarmPoolResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteWarmPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarmPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarmPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWarmPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWarmPoolResponse) SetHeaders(v map[string]*string) *DeleteWarmPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarmPoolResponse) SetStatusCode(v int32) *DeleteWarmPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarmPoolResponse) Validate() error {
	return dara.Validate(s)
}
