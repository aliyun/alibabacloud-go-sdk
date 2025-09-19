// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCreateVulWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCreateVulWhitelistResponseBody
	GetRequestId() *string
	SetVulWhitelistList(v []*ModifyCreateVulWhitelistResponseBodyVulWhitelistList) *ModifyCreateVulWhitelistResponseBody
	GetVulWhitelistList() []*ModifyCreateVulWhitelistResponseBodyVulWhitelistList
}

type ModifyCreateVulWhitelistResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6B23A612-D997-5176-8C3B-D640DFD65772
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the information about the whitelist.
	VulWhitelistList []*ModifyCreateVulWhitelistResponseBodyVulWhitelistList `json:"VulWhitelistList,omitempty" xml:"VulWhitelistList,omitempty" type:"Repeated"`
}

func (s ModifyCreateVulWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCreateVulWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCreateVulWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCreateVulWhitelistResponseBody) GetVulWhitelistList() []*ModifyCreateVulWhitelistResponseBodyVulWhitelistList {
	return s.VulWhitelistList
}

func (s *ModifyCreateVulWhitelistResponseBody) SetRequestId(v string) *ModifyCreateVulWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCreateVulWhitelistResponseBody) SetVulWhitelistList(v []*ModifyCreateVulWhitelistResponseBodyVulWhitelistList) *ModifyCreateVulWhitelistResponseBody {
	s.VulWhitelistList = v
	return s
}

func (s *ModifyCreateVulWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyCreateVulWhitelistResponseBodyVulWhitelistList struct {
	// The ID of the whitelist.
	//
	// example:
	//
	// 30376
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ModifyCreateVulWhitelistResponseBodyVulWhitelistList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCreateVulWhitelistResponseBodyVulWhitelistList) GoString() string {
	return s.String()
}

func (s *ModifyCreateVulWhitelistResponseBodyVulWhitelistList) GetId() *int64 {
	return s.Id
}

func (s *ModifyCreateVulWhitelistResponseBodyVulWhitelistList) SetId(v int64) *ModifyCreateVulWhitelistResponseBodyVulWhitelistList {
	s.Id = &v
	return s
}

func (s *ModifyCreateVulWhitelistResponseBodyVulWhitelistList) Validate() error {
	return dara.Validate(s)
}
