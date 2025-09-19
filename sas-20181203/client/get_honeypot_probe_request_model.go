// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetHoneypotProbeRequest
	GetLang() *string
	SetProbeId(v string) *GetHoneypotProbeRequest
	GetProbeId() *string
}

type GetHoneypotProbeRequest struct {
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
	// The probe ID.
	//
	// > You can call the [ListHoneypotProbe](~~ListHoneypotProbe~~) operation to query the IDs of probes.
	//
	// This parameter is required.
	//
	// example:
	//
	// d6c1ebc9-a90d-4c9e-9490-328814d1ca00
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
}

func (s GetHoneypotProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotProbeRequest) GoString() string {
	return s.String()
}

func (s *GetHoneypotProbeRequest) GetLang() *string {
	return s.Lang
}

func (s *GetHoneypotProbeRequest) GetProbeId() *string {
	return s.ProbeId
}

func (s *GetHoneypotProbeRequest) SetLang(v string) *GetHoneypotProbeRequest {
	s.Lang = &v
	return s
}

func (s *GetHoneypotProbeRequest) SetProbeId(v string) *GetHoneypotProbeRequest {
	s.ProbeId = &v
	return s
}

func (s *GetHoneypotProbeRequest) Validate() error {
	return dara.Validate(s)
}
