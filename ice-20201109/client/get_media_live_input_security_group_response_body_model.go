// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveInputSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMediaLiveInputSecurityGroupResponseBody
	GetRequestId() *string
	SetSecurityGroup(v *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) *GetMediaLiveInputSecurityGroupResponseBody
	GetSecurityGroup() *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup
}

type GetMediaLiveInputSecurityGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security group information.
	SecurityGroup *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Struct"`
}

func (s GetMediaLiveInputSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaLiveInputSecurityGroupResponseBody) GetSecurityGroup() *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup {
	return s.SecurityGroup
}

func (s *GetMediaLiveInputSecurityGroupResponseBody) SetRequestId(v string) *GetMediaLiveInputSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponseBody) SetSecurityGroup(v *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) *GetMediaLiveInputSecurityGroupResponseBody {
	s.SecurityGroup = v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponseBody) Validate() error {
	if s.SecurityGroup != nil {
		if err := s.SecurityGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMediaLiveInputSecurityGroupResponseBodySecurityGroup struct {
	// The time when the security group was created. It follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-06-13T08:31:56Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The IDs of the inputs associated with the security group.
	InputIds []*string `json:"InputIds,omitempty" xml:"InputIds,omitempty" type:"Repeated"`
	// The name of the security group.
	//
	// example:
	//
	// mysg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The security group rules.
	WhitelistRules []*string `json:"WhitelistRules,omitempty" xml:"WhitelistRules,omitempty" type:"Repeated"`
}

func (s GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) GetInputIds() []*string {
	return s.InputIds
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) GetName() *string {
	return s.Name
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) GetWhitelistRules() []*string {
	return s.WhitelistRules
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) SetCreateTime(v string) *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup {
	s.CreateTime = &v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) SetInputIds(v []*string) *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup {
	s.InputIds = v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) SetName(v string) *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup {
	s.Name = &v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) SetSecurityGroupId(v string) *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) SetWhitelistRules(v []*string) *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup {
	s.WhitelistRules = v
	return s
}

func (s *GetMediaLiveInputSecurityGroupResponseBodySecurityGroup) Validate() error {
	return dara.Validate(s)
}
