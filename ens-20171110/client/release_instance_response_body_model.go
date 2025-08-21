// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseInstanceResponseBody
	GetRequestId() *string
	SetResourceType(v string) *ReleaseInstanceResponseBody
	GetResourceType() *string
}

type ReleaseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the resource.
	//
	// Valid values:
	//
	// 	- instance
	//
	// 	- eip
	//
	// 	- disk
	//
	// 	- network
	//
	// 	- natgateway
	//
	// 	- vswitch
	//
	// example:
	//
	// eip
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseInstanceResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetResourceType(v string) *ReleaseInstanceResponseBody {
	s.ResourceType = &v
	return s
}

func (s *ReleaseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
