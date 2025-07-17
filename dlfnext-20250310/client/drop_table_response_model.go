// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropTableResponse
	GetStatusCode() *int32
}

type DropTableResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DropTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DropTableResponse) GoString() string {
	return s.String()
}

func (s *DropTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropTableResponse) SetHeaders(v map[string]*string) *DropTableResponse {
	s.Headers = v
	return s
}

func (s *DropTableResponse) SetStatusCode(v int32) *DropTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DropTableResponse) Validate() error {
	return dara.Validate(s)
}
