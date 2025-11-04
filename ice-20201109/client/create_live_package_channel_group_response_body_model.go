// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageChannelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageChannelGroup(v *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) *CreateLivePackageChannelGroupResponseBody
	GetLivePackageChannelGroup() *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup
	SetRequestId(v string) *CreateLivePackageChannelGroupResponseBody
	GetRequestId() *string
}

type CreateLivePackageChannelGroupResponseBody struct {
	// The information about the channel group.
	LivePackageChannelGroup *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup `json:"LivePackageChannelGroup,omitempty" xml:"LivePackageChannelGroup,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLivePackageChannelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelGroupResponseBody) GetLivePackageChannelGroup() *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	return s.LivePackageChannelGroup
}

func (s *CreateLivePackageChannelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLivePackageChannelGroupResponseBody) SetLivePackageChannelGroup(v *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) *CreateLivePackageChannelGroupResponseBody {
	s.LivePackageChannelGroup = v
	return s
}

func (s *CreateLivePackageChannelGroupResponseBody) SetRequestId(v string) *CreateLivePackageChannelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLivePackageChannelGroupResponseBody) Validate() error {
	if s.LivePackageChannelGroup != nil {
		if err := s.LivePackageChannelGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup struct {
	// The time when the channel group was created. It is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2023-04-01T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The channel group description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name.
	//
	// example:
	//
	// example-group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the channel group was last modified. It is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2023-04-01T12:00:00Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The origin domain.
	//
	// example:
	//
	// example.com
	OriginDomain *string `json:"OriginDomain,omitempty" xml:"OriginDomain,omitempty"`
}

func (s CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetDescription() *string {
	return s.Description
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetLastModified() *string {
	return s.LastModified
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetOriginDomain() *string {
	return s.OriginDomain
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetCreateTime(v string) *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.CreateTime = &v
	return s
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetDescription(v string) *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.Description = &v
	return s
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetGroupName(v string) *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.GroupName = &v
	return s
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetLastModified(v string) *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.LastModified = &v
	return s
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetOriginDomain(v string) *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.OriginDomain = &v
	return s
}

func (s *CreateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) Validate() error {
	return dara.Validate(s)
}
