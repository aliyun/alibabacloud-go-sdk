// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteHoneypotProbeRequest
	GetLang() *string
	SetProbeId(v string) *DeleteHoneypotProbeRequest
	GetProbeId() *string
}

type DeleteHoneypotProbeRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the probe.
	//
	// > You can call the [ListHoneypotProbe](~~ListHoneypotProbe~~) operation to query the IDs of probes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95f0f79c-f7e9-4b09-a6e3-95a4cb6d****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
}

func (s DeleteHoneypotProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotProbeRequest) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotProbeRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteHoneypotProbeRequest) GetProbeId() *string {
	return s.ProbeId
}

func (s *DeleteHoneypotProbeRequest) SetLang(v string) *DeleteHoneypotProbeRequest {
	s.Lang = &v
	return s
}

func (s *DeleteHoneypotProbeRequest) SetProbeId(v string) *DeleteHoneypotProbeRequest {
	s.ProbeId = &v
	return s
}

func (s *DeleteHoneypotProbeRequest) Validate() error {
	return dara.Validate(s)
}
