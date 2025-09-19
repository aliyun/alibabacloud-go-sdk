// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateImageBaselineWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineItemKeyList(v string) *OperateImageBaselineWhitelistRequest
	GetBaselineItemKeyList() *string
	SetImageUuid(v string) *OperateImageBaselineWhitelistRequest
	GetImageUuid() *string
	SetLang(v string) *OperateImageBaselineWhitelistRequest
	GetLang() *string
	SetOperation(v string) *OperateImageBaselineWhitelistRequest
	GetOperation() *string
	SetScanRange(v []*string) *OperateImageBaselineWhitelistRequest
	GetScanRange() []*string
}

type OperateImageBaselineWhitelistRequest struct {
	// The keys of baseline check items.
	//
	// This parameter is required.
	//
	// example:
	//
	// hc.image.checklist.identify.hc_exploit_es_linux.item
	BaselineItemKeyList *string `json:"BaselineItemKeyList,omitempty" xml:"BaselineItemKeyList,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// a5250ebca765dc9eb1a84b790b0e****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The operation that you want to perform on the check items. Valid values:
	//
	// 	- **add**: adds the check items to the whitelist
	//
	// 	- **del**: removes the check items from the whitelist
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The types of the assets that you want to scan.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
}

func (s OperateImageBaselineWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateImageBaselineWhitelistRequest) GoString() string {
	return s.String()
}

func (s *OperateImageBaselineWhitelistRequest) GetBaselineItemKeyList() *string {
	return s.BaselineItemKeyList
}

func (s *OperateImageBaselineWhitelistRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *OperateImageBaselineWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *OperateImageBaselineWhitelistRequest) GetOperation() *string {
	return s.Operation
}

func (s *OperateImageBaselineWhitelistRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *OperateImageBaselineWhitelistRequest) SetBaselineItemKeyList(v string) *OperateImageBaselineWhitelistRequest {
	s.BaselineItemKeyList = &v
	return s
}

func (s *OperateImageBaselineWhitelistRequest) SetImageUuid(v string) *OperateImageBaselineWhitelistRequest {
	s.ImageUuid = &v
	return s
}

func (s *OperateImageBaselineWhitelistRequest) SetLang(v string) *OperateImageBaselineWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *OperateImageBaselineWhitelistRequest) SetOperation(v string) *OperateImageBaselineWhitelistRequest {
	s.Operation = &v
	return s
}

func (s *OperateImageBaselineWhitelistRequest) SetScanRange(v []*string) *OperateImageBaselineWhitelistRequest {
	s.ScanRange = v
	return s
}

func (s *OperateImageBaselineWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
