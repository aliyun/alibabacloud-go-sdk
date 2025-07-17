// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateProjectFromResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateProjectFromResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DissociateProjectFromResourceGroupResponseBody
	GetSuccess() *bool
}

type DissociateProjectFromResourceGroupResponseBody struct {
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

func (s DissociateProjectFromResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateProjectFromResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateProjectFromResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DissociateProjectFromResourceGroupResponseBody) SetRequestId(v string) *DissociateProjectFromResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateProjectFromResourceGroupResponseBody) SetSuccess(v bool) *DissociateProjectFromResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DissociateProjectFromResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
