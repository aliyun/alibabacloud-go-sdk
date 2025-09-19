// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBaselineSecurityCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemIds(v []*int64) *StartBaselineSecurityCheckRequest
	GetItemIds() []*int64
	SetLang(v string) *StartBaselineSecurityCheckRequest
	GetLang() *string
	SetResourceOwnerId(v int64) *StartBaselineSecurityCheckRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *StartBaselineSecurityCheckRequest
	GetSourceIp() *string
	SetType(v string) *StartBaselineSecurityCheckRequest
	GetType() *string
}

type StartBaselineSecurityCheckRequest struct {
	// The IDs of the check items.
	//
	// > To perform a check task on cloud service configurations, you must specify the ID of the check item. You can call the [DescribeRiskItemType](~~DescribeRiskItemType~~) operation to query the IDs of check items.
	ItemIds []*int64 `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 106.11.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the check task. Valid values:
	//
	// 	- **check**
	//
	// 	- **verify**
	//
	// This parameter is required.
	//
	// example:
	//
	// verify
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StartBaselineSecurityCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s StartBaselineSecurityCheckRequest) GoString() string {
	return s.String()
}

func (s *StartBaselineSecurityCheckRequest) GetItemIds() []*int64 {
	return s.ItemIds
}

func (s *StartBaselineSecurityCheckRequest) GetLang() *string {
	return s.Lang
}

func (s *StartBaselineSecurityCheckRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartBaselineSecurityCheckRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *StartBaselineSecurityCheckRequest) GetType() *string {
	return s.Type
}

func (s *StartBaselineSecurityCheckRequest) SetItemIds(v []*int64) *StartBaselineSecurityCheckRequest {
	s.ItemIds = v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetLang(v string) *StartBaselineSecurityCheckRequest {
	s.Lang = &v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetResourceOwnerId(v int64) *StartBaselineSecurityCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetSourceIp(v string) *StartBaselineSecurityCheckRequest {
	s.SourceIp = &v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetType(v string) *StartBaselineSecurityCheckRequest {
	s.Type = &v
	return s
}

func (s *StartBaselineSecurityCheckRequest) Validate() error {
	return dara.Validate(s)
}
