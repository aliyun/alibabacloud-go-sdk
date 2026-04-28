// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPunishFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PunishFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PunishFileResponse
	GetStatusCode() *int32
}

type PunishFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PunishFileResponse) String() string {
	return dara.Prettify(s)
}

func (s PunishFileResponse) GoString() string {
	return s.String()
}

func (s *PunishFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PunishFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PunishFileResponse) SetHeaders(v map[string]*string) *PunishFileResponse {
	s.Headers = v
	return s
}

func (s *PunishFileResponse) SetStatusCode(v int32) *PunishFileResponse {
	s.StatusCode = &v
	return s
}

func (s *PunishFileResponse) Validate() error {
	return dara.Validate(s)
}
