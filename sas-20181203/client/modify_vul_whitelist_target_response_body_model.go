// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulWhitelistTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVulWhitelistTargetResponseBody
	GetRequestId() *string
}

type ModifyVulWhitelistTargetResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 896AC4F0-C881-502B-BFC7-4751C5E3DEAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVulWhitelistTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulWhitelistTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVulWhitelistTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVulWhitelistTargetResponseBody) SetRequestId(v string) *ModifyVulWhitelistTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVulWhitelistTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
