// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisconnectDesktopSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPreCheck(v bool) *DisconnectDesktopSessionsRequest
	GetPreCheck() *bool
	SetRegionId(v string) *DisconnectDesktopSessionsRequest
	GetRegionId() *string
	SetSessions(v []*DisconnectDesktopSessionsRequestSessions) *DisconnectDesktopSessionsRequest
	GetSessions() []*DisconnectDesktopSessionsRequestSessions
}

type DisconnectDesktopSessionsRequest struct {
	// Specifies whether to perform precheck. If you perform precheck, the system does not disconnect from desktop sessions. Only the sessions that do not meet specific conditions are returned.
	//
	// example:
	//
	// true
	PreCheck *bool `json:"PreCheck,omitempty" xml:"PreCheck,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session details.
	//
	// This parameter is required.
	Sessions []*DisconnectDesktopSessionsRequestSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
}

func (s DisconnectDesktopSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DisconnectDesktopSessionsRequest) GoString() string {
	return s.String()
}

func (s *DisconnectDesktopSessionsRequest) GetPreCheck() *bool {
	return s.PreCheck
}

func (s *DisconnectDesktopSessionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisconnectDesktopSessionsRequest) GetSessions() []*DisconnectDesktopSessionsRequestSessions {
	return s.Sessions
}

func (s *DisconnectDesktopSessionsRequest) SetPreCheck(v bool) *DisconnectDesktopSessionsRequest {
	s.PreCheck = &v
	return s
}

func (s *DisconnectDesktopSessionsRequest) SetRegionId(v string) *DisconnectDesktopSessionsRequest {
	s.RegionId = &v
	return s
}

func (s *DisconnectDesktopSessionsRequest) SetSessions(v []*DisconnectDesktopSessionsRequestSessions) *DisconnectDesktopSessionsRequest {
	s.Sessions = v
	return s
}

func (s *DisconnectDesktopSessionsRequest) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DisconnectDesktopSessionsRequestSessions struct {
	// The cloud desktop ID.
	//
	// example:
	//
	// ecd-90g15fkhsxxxn0unj
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// wy01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
}

func (s DisconnectDesktopSessionsRequestSessions) String() string {
	return dara.Prettify(s)
}

func (s DisconnectDesktopSessionsRequestSessions) GoString() string {
	return s.String()
}

func (s *DisconnectDesktopSessionsRequestSessions) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DisconnectDesktopSessionsRequestSessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DisconnectDesktopSessionsRequestSessions) SetDesktopId(v string) *DisconnectDesktopSessionsRequestSessions {
	s.DesktopId = &v
	return s
}

func (s *DisconnectDesktopSessionsRequestSessions) SetEndUserId(v string) *DisconnectDesktopSessionsRequestSessions {
	s.EndUserId = &v
	return s
}

func (s *DisconnectDesktopSessionsRequestSessions) Validate() error {
	return dara.Validate(s)
}
