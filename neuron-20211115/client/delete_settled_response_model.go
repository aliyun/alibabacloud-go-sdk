// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSettledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSettledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSettledResponse
	GetStatusCode() *int32
	SetBody(v *CommonResult) *DeleteSettledResponse
	GetBody() *CommonResult
}

type DeleteSettledResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CommonResult      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSettledResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSettledResponse) GoString() string {
	return s.String()
}

func (s *DeleteSettledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSettledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSettledResponse) GetBody() *CommonResult {
	return s.Body
}

func (s *DeleteSettledResponse) SetHeaders(v map[string]*string) *DeleteSettledResponse {
	s.Headers = v
	return s
}

func (s *DeleteSettledResponse) SetStatusCode(v int32) *DeleteSettledResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSettledResponse) SetBody(v *CommonResult) *DeleteSettledResponse {
	s.Body = v
	return s
}

func (s *DeleteSettledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
