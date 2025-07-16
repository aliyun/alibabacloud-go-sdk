// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubscribeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubscribeResponse
	GetStatusCode() *int32
}

type SubscribeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s SubscribeResponse) String() string {
	return dara.Prettify(s)
}

func (s SubscribeResponse) GoString() string {
	return s.String()
}

func (s *SubscribeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubscribeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubscribeResponse) SetHeaders(v map[string]*string) *SubscribeResponse {
	s.Headers = v
	return s
}

func (s *SubscribeResponse) SetStatusCode(v int32) *SubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeResponse) Validate() error {
	return dara.Validate(s)
}
