// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropShareResponse
	GetStatusCode() *int32
}

type DropShareResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DropShareResponse) String() string {
	return dara.Prettify(s)
}

func (s DropShareResponse) GoString() string {
	return s.String()
}

func (s *DropShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropShareResponse) SetHeaders(v map[string]*string) *DropShareResponse {
	s.Headers = v
	return s
}

func (s *DropShareResponse) SetStatusCode(v int32) *DropShareResponse {
	s.StatusCode = &v
	return s
}

func (s *DropShareResponse) Validate() error {
	return dara.Validate(s)
}
