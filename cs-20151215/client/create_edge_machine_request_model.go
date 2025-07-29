// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostname(v string) *CreateEdgeMachineRequest
	GetHostname() *string
	SetModel(v string) *CreateEdgeMachineRequest
	GetModel() *string
	SetSn(v string) *CreateEdgeMachineRequest
	GetSn() *string
}

type CreateEdgeMachineRequest struct {
	// The `hostname` of the cloud-native box.
	//
	// >  After the cloud-native box is activated, the `hostname` is automatically modified. The `hostname` is prefixed with the model and the prefix is followed by a random string.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACK-B-B010-****
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The model of the cloud-native box.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACK-V-B010
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The serial number of the cloud-native box.
	//
	// This parameter is required.
	//
	// example:
	//
	// Q2CB5XZAFBFG****
	Sn *string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (s CreateEdgeMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeMachineRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeMachineRequest) GetHostname() *string {
	return s.Hostname
}

func (s *CreateEdgeMachineRequest) GetModel() *string {
	return s.Model
}

func (s *CreateEdgeMachineRequest) GetSn() *string {
	return s.Sn
}

func (s *CreateEdgeMachineRequest) SetHostname(v string) *CreateEdgeMachineRequest {
	s.Hostname = &v
	return s
}

func (s *CreateEdgeMachineRequest) SetModel(v string) *CreateEdgeMachineRequest {
	s.Model = &v
	return s
}

func (s *CreateEdgeMachineRequest) SetSn(v string) *CreateEdgeMachineRequest {
	s.Sn = &v
	return s
}

func (s *CreateEdgeMachineRequest) Validate() error {
	return dara.Validate(s)
}
