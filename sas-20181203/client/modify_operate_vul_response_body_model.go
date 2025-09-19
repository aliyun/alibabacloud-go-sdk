// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOperateVulResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOperateVulResponseBody
	GetRequestId() *string
}

type ModifyOperateVulResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DFE4F166-1AC9-4FAC-A4E4-F0608AD705A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOperateVulResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOperateVulResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOperateVulResponseBody) SetRequestId(v string) *ModifyOperateVulResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOperateVulResponseBody) Validate() error {
	return dara.Validate(s)
}
