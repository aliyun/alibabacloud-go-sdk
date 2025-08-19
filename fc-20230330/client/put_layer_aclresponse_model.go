// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLayerACLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutLayerACLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutLayerACLResponse
	GetStatusCode() *int32
}

type PutLayerACLResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutLayerACLResponse) String() string {
	return dara.Prettify(s)
}

func (s PutLayerACLResponse) GoString() string {
	return s.String()
}

func (s *PutLayerACLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutLayerACLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutLayerACLResponse) SetHeaders(v map[string]*string) *PutLayerACLResponse {
	s.Headers = v
	return s
}

func (s *PutLayerACLResponse) SetStatusCode(v int32) *PutLayerACLResponse {
	s.StatusCode = &v
	return s
}

func (s *PutLayerACLResponse) Validate() error {
	return dara.Validate(s)
}
