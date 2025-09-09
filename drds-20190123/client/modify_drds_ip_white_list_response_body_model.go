// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDrdsIpWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDrdsIpWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDrdsIpWhiteListResponseBody
	GetSuccess() *bool
}

type ModifyDrdsIpWhiteListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 933A1EC2-8260-4D4F-A56A-73BA27******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDrdsIpWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDrdsIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsIpWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDrdsIpWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDrdsIpWhiteListResponseBody) SetRequestId(v string) *ModifyDrdsIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDrdsIpWhiteListResponseBody) SetSuccess(v bool) *ModifyDrdsIpWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDrdsIpWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
