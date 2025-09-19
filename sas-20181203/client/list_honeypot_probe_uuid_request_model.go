// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotProbeUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlNodeId(v string) *ListHoneypotProbeUuidRequest
	GetControlNodeId() *string
	SetLang(v string) *ListHoneypotProbeUuidRequest
	GetLang() *string
	SetProbeType(v string) *ListHoneypotProbeUuidRequest
	GetProbeType() *string
}

type ListHoneypotProbeUuidRequest struct {
	// The ID of the management node.
	//
	// >  You can call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to obtain the ID.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	ControlNodeId *string `json:"ControlNodeId,omitempty" xml:"ControlNodeId,omitempty"`
	// The language of the content within the request and the response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the probe. Valid values:
	//
	// 	- **host_probe**: host probe
	//
	// 	- **vpc_black_hole_probe**: virtual private cloud (VPC) probe
	//
	// example:
	//
	// host_probe
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
}

func (s ListHoneypotProbeUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotProbeUuidRequest) GoString() string {
	return s.String()
}

func (s *ListHoneypotProbeUuidRequest) GetControlNodeId() *string {
	return s.ControlNodeId
}

func (s *ListHoneypotProbeUuidRequest) GetLang() *string {
	return s.Lang
}

func (s *ListHoneypotProbeUuidRequest) GetProbeType() *string {
	return s.ProbeType
}

func (s *ListHoneypotProbeUuidRequest) SetControlNodeId(v string) *ListHoneypotProbeUuidRequest {
	s.ControlNodeId = &v
	return s
}

func (s *ListHoneypotProbeUuidRequest) SetLang(v string) *ListHoneypotProbeUuidRequest {
	s.Lang = &v
	return s
}

func (s *ListHoneypotProbeUuidRequest) SetProbeType(v string) *ListHoneypotProbeUuidRequest {
	s.ProbeType = &v
	return s
}

func (s *ListHoneypotProbeUuidRequest) Validate() error {
	return dara.Validate(s)
}
