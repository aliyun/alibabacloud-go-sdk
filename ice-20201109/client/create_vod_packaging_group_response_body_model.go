// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodPackagingGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackagingGroup(v *VodPackagingGroup) *CreateVodPackagingGroupResponseBody
	GetPackagingGroup() *VodPackagingGroup
	SetRequestId(v string) *CreateVodPackagingGroupResponseBody
	GetRequestId() *string
}

type CreateVodPackagingGroupResponseBody struct {
	// The packaging group information.
	PackagingGroup *VodPackagingGroup `json:"PackagingGroup,omitempty" xml:"PackagingGroup,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVodPackagingGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVodPackagingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVodPackagingGroupResponseBody) GetPackagingGroup() *VodPackagingGroup {
	return s.PackagingGroup
}

func (s *CreateVodPackagingGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVodPackagingGroupResponseBody) SetPackagingGroup(v *VodPackagingGroup) *CreateVodPackagingGroupResponseBody {
	s.PackagingGroup = v
	return s
}

func (s *CreateVodPackagingGroupResponseBody) SetRequestId(v string) *CreateVodPackagingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVodPackagingGroupResponseBody) Validate() error {
	if s.PackagingGroup != nil {
		if err := s.PackagingGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}
