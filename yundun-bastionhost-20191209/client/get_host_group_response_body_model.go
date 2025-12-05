// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroup(v *GetHostGroupResponseBodyHostGroup) *GetHostGroupResponseBody
	GetHostGroup() *GetHostGroupResponseBodyHostGroup
	SetRequestId(v string) *GetHostGroupResponseBody
	GetRequestId() *string
}

type GetHostGroupResponseBody struct {
	// The returned detailed information about the asset group.
	HostGroup *GetHostGroupResponseBodyHostGroup `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHostGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBody) GetHostGroup() *GetHostGroupResponseBodyHostGroup {
	return s.HostGroup
}

func (s *GetHostGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHostGroupResponseBody) SetHostGroup(v *GetHostGroupResponseBodyHostGroup) *GetHostGroupResponseBody {
	s.HostGroup = v
	return s
}

func (s *GetHostGroupResponseBody) SetRequestId(v string) *GetHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostGroupResponseBody) Validate() error {
	if s.HostGroup != nil {
		if err := s.HostGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHostGroupResponseBodyHostGroup struct {
	// The remarks of the asset group.
	//
	// example:
	//
	// Description
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The asset group ID.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The name of the asset group.
	//
	// example:
	//
	// Host group 1
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s GetHostGroupResponseBodyHostGroup) String() string {
	return dara.Prettify(s)
}

func (s GetHostGroupResponseBodyHostGroup) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBodyHostGroup) GetComment() *string {
	return s.Comment
}

func (s *GetHostGroupResponseBodyHostGroup) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *GetHostGroupResponseBodyHostGroup) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *GetHostGroupResponseBodyHostGroup) SetComment(v string) *GetHostGroupResponseBodyHostGroup {
	s.Comment = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetHostGroupId(v string) *GetHostGroupResponseBodyHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetHostGroupName(v string) *GetHostGroupResponseBodyHostGroup {
	s.HostGroupName = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) Validate() error {
	return dara.Validate(s)
}
