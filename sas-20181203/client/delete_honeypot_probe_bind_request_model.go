// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotProbeBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindId(v string) *DeleteHoneypotProbeBindRequest
	GetBindId() *string
	SetLang(v string) *DeleteHoneypotProbeBindRequest
	GetLang() *string
	SetProbeId(v string) *DeleteHoneypotProbeBindRequest
	GetProbeId() *string
}

type DeleteHoneypotProbeBindRequest struct {
	// The unique ID of the bound service.
	//
	// example:
	//
	// aa20815f-f0f3-4e3b-8e13-55771742****
	BindId *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The probe ID.
	//
	// >  You can call the [ListHoneypotProbe](~~ListHoneypotProbe~~) operation to query the IDs of probes.
	//
	// example:
	//
	// aa234650-cfcf-4e25-b61f-c58f603f****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
}

func (s DeleteHoneypotProbeBindRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotProbeBindRequest) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotProbeBindRequest) GetBindId() *string {
	return s.BindId
}

func (s *DeleteHoneypotProbeBindRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteHoneypotProbeBindRequest) GetProbeId() *string {
	return s.ProbeId
}

func (s *DeleteHoneypotProbeBindRequest) SetBindId(v string) *DeleteHoneypotProbeBindRequest {
	s.BindId = &v
	return s
}

func (s *DeleteHoneypotProbeBindRequest) SetLang(v string) *DeleteHoneypotProbeBindRequest {
	s.Lang = &v
	return s
}

func (s *DeleteHoneypotProbeBindRequest) SetProbeId(v string) *DeleteHoneypotProbeBindRequest {
	s.ProbeId = &v
	return s
}

func (s *DeleteHoneypotProbeBindRequest) Validate() error {
	return dara.Validate(s)
}
