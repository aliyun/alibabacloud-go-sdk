// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWaitingRoomConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowPct(v int32) *SetWaitingRoomConfigRequest
	GetAllowPct() *int32
	SetDomainName(v string) *SetWaitingRoomConfigRequest
	GetDomainName() *string
	SetGapTime(v int32) *SetWaitingRoomConfigRequest
	GetGapTime() *int32
	SetMaxTimeWait(v int32) *SetWaitingRoomConfigRequest
	GetMaxTimeWait() *int32
	SetWaitUri(v string) *SetWaitingRoomConfigRequest
	GetWaitUri() *string
	SetWaitUrl(v string) *SetWaitingRoomConfigRequest
	GetWaitUrl() *string
}

type SetWaitingRoomConfigRequest struct {
	// The percentage of requests that are allowed to be redirected to the origin server. Valid values: **0*	- to **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	AllowPct *int32 `json:"AllowPct,omitempty" xml:"AllowPct,omitempty"`
	// The accelerated domain name. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The length of waiting time to skip after an exit from the queue. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	GapTime *int32 `json:"GapTime,omitempty" xml:"GapTime,omitempty"`
	// The maximum length of waiting time in the queue. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	MaxTimeWait *int32 `json:"MaxTimeWait,omitempty" xml:"MaxTimeWait,omitempty"`
	// The regular expression that is used to match URI strings for which the virtual waiting room feature is enabled.
	//
	// This parameter is required.
	WaitUri *string `json:"WaitUri,omitempty" xml:"WaitUri,omitempty"`
	// The URL of the waiting page.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/waitingroom.html
	WaitUrl *string `json:"WaitUrl,omitempty" xml:"WaitUrl,omitempty"`
}

func (s SetWaitingRoomConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWaitingRoomConfigRequest) GoString() string {
	return s.String()
}

func (s *SetWaitingRoomConfigRequest) GetAllowPct() *int32 {
	return s.AllowPct
}

func (s *SetWaitingRoomConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetWaitingRoomConfigRequest) GetGapTime() *int32 {
	return s.GapTime
}

func (s *SetWaitingRoomConfigRequest) GetMaxTimeWait() *int32 {
	return s.MaxTimeWait
}

func (s *SetWaitingRoomConfigRequest) GetWaitUri() *string {
	return s.WaitUri
}

func (s *SetWaitingRoomConfigRequest) GetWaitUrl() *string {
	return s.WaitUrl
}

func (s *SetWaitingRoomConfigRequest) SetAllowPct(v int32) *SetWaitingRoomConfigRequest {
	s.AllowPct = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetDomainName(v string) *SetWaitingRoomConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetGapTime(v int32) *SetWaitingRoomConfigRequest {
	s.GapTime = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetMaxTimeWait(v int32) *SetWaitingRoomConfigRequest {
	s.MaxTimeWait = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetWaitUri(v string) *SetWaitingRoomConfigRequest {
	s.WaitUri = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) SetWaitUrl(v string) *SetWaitingRoomConfigRequest {
	s.WaitUrl = &v
	return s
}

func (s *SetWaitingRoomConfigRequest) Validate() error {
	return dara.Validate(s)
}
