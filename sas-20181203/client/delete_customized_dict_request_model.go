// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizedDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *DeleteCustomizedDictRequest
	GetSourceIp() *string
}

type DeleteCustomizedDictRequest struct {
	// The source IP address.
	//
	// example:
	//
	// 123.103.9.***
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteCustomizedDictRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizedDictRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedDictRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteCustomizedDictRequest) SetSourceIp(v string) *DeleteCustomizedDictRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteCustomizedDictRequest) Validate() error {
	return dara.Validate(s)
}
