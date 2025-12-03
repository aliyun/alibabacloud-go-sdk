// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHBaseSlbServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateHBaseSlbServerRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateHBaseSlbServerRequest
	GetClusterId() *string
	SetSlbServer(v string) *CreateHBaseSlbServerRequest
	GetSlbServer() *string
}

type CreateHBaseSlbServerRequest struct {
	// example:
	//
	// xxxxx-xxxxx-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s CreateHBaseSlbServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHBaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateHBaseSlbServerRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateHBaseSlbServerRequest) GetSlbServer() *string {
	return s.SlbServer
}

func (s *CreateHBaseSlbServerRequest) SetClientToken(v string) *CreateHBaseSlbServerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHBaseSlbServerRequest) SetClusterId(v string) *CreateHBaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHBaseSlbServerRequest) SetSlbServer(v string) *CreateHBaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

func (s *CreateHBaseSlbServerRequest) Validate() error {
	return dara.Validate(s)
}
