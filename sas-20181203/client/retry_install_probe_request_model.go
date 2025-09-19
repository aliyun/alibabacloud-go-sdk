// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryInstallProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProbeId(v string) *RetryInstallProbeRequest
	GetProbeId() *string
}

type RetryInstallProbeRequest struct {
	// The probe ID.
	//
	// >  You can call the [ListHoneypotProbe](~~ListHoneypotProbe~~) operation to query the IDs of probes.
	//
	// example:
	//
	// c4c47cc1-f60a-4b2f-bcdb-9aed6644****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
}

func (s RetryInstallProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryInstallProbeRequest) GoString() string {
	return s.String()
}

func (s *RetryInstallProbeRequest) GetProbeId() *string {
	return s.ProbeId
}

func (s *RetryInstallProbeRequest) SetProbeId(v string) *RetryInstallProbeRequest {
	s.ProbeId = &v
	return s
}

func (s *RetryInstallProbeRequest) Validate() error {
	return dara.Validate(s)
}
