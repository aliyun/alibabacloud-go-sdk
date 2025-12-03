// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHBaseSlbServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteHBaseSlbServerRequest
	GetClusterId() *string
	SetSlbServer(v string) *DeleteHBaseSlbServerRequest
	GetSlbServer() *string
}

type DeleteHBaseSlbServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thrift
	SlbServer *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
}

func (s DeleteHBaseSlbServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHBaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteHBaseSlbServerRequest) GetSlbServer() *string {
	return s.SlbServer
}

func (s *DeleteHBaseSlbServerRequest) SetClusterId(v string) *DeleteHBaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHBaseSlbServerRequest) SetSlbServer(v string) *DeleteHBaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

func (s *DeleteHBaseSlbServerRequest) Validate() error {
	return dara.Validate(s)
}
