// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionResourceForFrontResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPermissionResourceForFrontResponseBody
	GetRequestId() *string
	SetResourceList(v []*string) *ListPermissionResourceForFrontResponseBody
	GetResourceList() []*string
}

type ListPermissionResourceForFrontResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1D4A64A4-18AB-52CF-AB79-517AEC7DC63B
	RequestId    *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceList []*string `json:"resourceList,omitempty" xml:"resourceList,omitempty" type:"Repeated"`
}

func (s ListPermissionResourceForFrontResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionResourceForFrontResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionResourceForFrontResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPermissionResourceForFrontResponseBody) GetResourceList() []*string {
	return s.ResourceList
}

func (s *ListPermissionResourceForFrontResponseBody) SetRequestId(v string) *ListPermissionResourceForFrontResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionResourceForFrontResponseBody) SetResourceList(v []*string) *ListPermissionResourceForFrontResponseBody {
	s.ResourceList = v
	return s
}

func (s *ListPermissionResourceForFrontResponseBody) Validate() error {
	return dara.Validate(s)
}
