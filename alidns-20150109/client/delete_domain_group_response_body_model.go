// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteDomainGroupResponseBody
	GetGroupName() *string
	SetRequestId(v string) *DeleteDomainGroupResponseBody
	GetRequestId() *string
}

type DeleteDomainGroupResponseBody struct {
	// The name of the domain name group.
	//
	// example:
	//
	// MyGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteDomainGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainGroupResponseBody) SetGroupName(v string) *DeleteDomainGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *DeleteDomainGroupResponseBody) SetRequestId(v string) *DeleteDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
