// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageChannelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageChannelGroup(v *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) *GetLivePackageChannelGroupResponseBody
	GetLivePackageChannelGroup() *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup
	SetRequestId(v string) *GetLivePackageChannelGroupResponseBody
	GetRequestId() *string
}

type GetLivePackageChannelGroupResponseBody struct {
	// Details of the channel group.
	LivePackageChannelGroup *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup `json:"LivePackageChannelGroup,omitempty" xml:"LivePackageChannelGroup,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// req-abcdefg123456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLivePackageChannelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelGroupResponseBody) GetLivePackageChannelGroup() *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	return s.LivePackageChannelGroup
}

func (s *GetLivePackageChannelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLivePackageChannelGroupResponseBody) SetLivePackageChannelGroup(v *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) *GetLivePackageChannelGroupResponseBody {
	s.LivePackageChannelGroup = v
	return s
}

func (s *GetLivePackageChannelGroupResponseBody) SetRequestId(v string) *GetLivePackageChannelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLivePackageChannelGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup struct {
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
	// channel-group-1
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

func (s GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetDescription() *string {
	return s.Description
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetLastModified() *string {
	return s.LastModified
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) GetOriginDomain() *string {
	return s.OriginDomain
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetCreateTime(v string) *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.CreateTime = &v
	return s
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetDescription(v string) *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.Description = &v
	return s
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetGroupName(v string) *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.GroupName = &v
	return s
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetLastModified(v string) *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.LastModified = &v
	return s
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) SetOriginDomain(v string) *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup {
	s.OriginDomain = &v
	return s
}

func (s *GetLivePackageChannelGroupResponseBodyLivePackageChannelGroup) Validate() error {
	return dara.Validate(s)
}
