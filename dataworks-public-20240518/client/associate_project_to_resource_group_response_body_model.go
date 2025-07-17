// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateProjectToResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateProjectToResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssociateProjectToResourceGroupResponseBody
	GetSuccess() *bool
}

type AssociateProjectToResourceGroupResponseBody struct {
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssociateProjectToResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateProjectToResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateProjectToResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateProjectToResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssociateProjectToResourceGroupResponseBody) SetRequestId(v string) *AssociateProjectToResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateProjectToResourceGroupResponseBody) SetSuccess(v bool) *AssociateProjectToResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *AssociateProjectToResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
