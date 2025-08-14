// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageChannelGroup(v *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) *UpdateLivePackageChannelGroupResponseBody
	GetLivePackageChannelGroup() *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup
	SetRequestId(v string) *UpdateLivePackageChannelGroupResponseBody
	GetRequestId() *string
}

type UpdateLivePackageChannelGroupResponseBody struct {
	// The information about the channel group.
	LivePackageChannelGroup *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup `json:"LivePackageChannelGroup,omitempty" xml:"LivePackageChannelGroup,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// request-1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLivePackageChannelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelGroupResponseBody) GetLivePackageChannelGroup() *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	return s.LivePackageChannelGroup
}

func (s *UpdateLivePackageChannelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLivePackageChannelGroupResponseBody) SetLivePackageChannelGroup(v *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) *UpdateLivePackageChannelGroupResponseBody {
	s.LivePackageChannelGroup = v
	return s
}

func (s *UpdateLivePackageChannelGroupResponseBody) SetRequestId(v string) *UpdateLivePackageChannelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivePackageChannelGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup struct {
	// The time when the channel group was created. It is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2023-04-01T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The channel group description.
	//
	// example:
	//
	// Updated description of the channel group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name.
	//
	// example:
	//
	// example-group-name
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
	// example-origin.com
	OriginDomain *string `json:"OriginDomain,omitempty" xml:"OriginDomain,omitempty"`
}

func (s UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetDescription() *string {
	return s.Description
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetLastModified() *string {
	return s.LastModified
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetOriginDomain() *string {
	return s.OriginDomain
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetCreateTime(v string) *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.CreateTime = &v
	return s
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetDescription(v string) *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.Description = &v
	return s
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetGroupName(v string) *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetLastModified(v string) *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.LastModified = &v
	return s
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetOriginDomain(v string) *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.OriginDomain = &v
	return s
}

func (s *UpdateLivePackageChannelGroupResponseBodyLivePackageChannelGroup) Validate() error {
	return dara.Validate(s)
}
