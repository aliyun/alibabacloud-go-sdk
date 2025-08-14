// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackagingGroup(v *VodPackagingGroup) *GetVodPackagingGroupResponseBody
	GetPackagingGroup() *VodPackagingGroup
	SetRequestId(v string) *GetVodPackagingGroupResponseBody
	GetRequestId() *string
}

type GetVodPackagingGroupResponseBody struct {
	// The information about the packaging group.
	PackagingGroup *VodPackagingGroup `json:"PackagingGroup,omitempty" xml:"PackagingGroup,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ****36-3C1E-4417-BDB2-1E034F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVodPackagingGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetVodPackagingGroupResponseBody) GetPackagingGroup() *VodPackagingGroup {
	return s.PackagingGroup
}

func (s *GetVodPackagingGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVodPackagingGroupResponseBody) SetPackagingGroup(v *VodPackagingGroup) *GetVodPackagingGroupResponseBody {
	s.PackagingGroup = v
	return s
}

func (s *GetVodPackagingGroupResponseBody) SetRequestId(v string) *GetVodPackagingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVodPackagingGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
