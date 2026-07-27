// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *ModifyInstanceConfigResponseBody
	GetBranchName() *string
	SetInstanceName(v string) *ModifyInstanceConfigResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *ModifyInstanceConfigResponseBody
	GetRequestId() *string
}

type ModifyInstanceConfigResponseBody struct {
	BranchName   *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 32DEFB4A-861F-5D1D-ADD5-918E4FD7AB8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *ModifyInstanceConfigResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceConfigResponseBody) SetBranchName(v string) *ModifyInstanceConfigResponseBody {
	s.BranchName = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetInstanceName(v string) *ModifyInstanceConfigResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) SetRequestId(v string) *ModifyInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
