// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDiagnosisCheckItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckItems(v []*GetClusterDiagnosisCheckItemsResponseBodyCheckItems) *GetClusterDiagnosisCheckItemsResponseBody
	GetCheckItems() []*GetClusterDiagnosisCheckItemsResponseBodyCheckItems
	SetCode(v string) *GetClusterDiagnosisCheckItemsResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *GetClusterDiagnosisCheckItemsResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetClusterDiagnosisCheckItemsResponseBody
	GetRequestId() *string
}

type GetClusterDiagnosisCheckItemsResponseBody struct {
	// The check item.
	CheckItems []*GetClusterDiagnosisCheckItemsResponseBodyCheckItems `json:"check_items,omitempty" xml:"check_items,omitempty" type:"Repeated"`
	// The status code.
	//
	// example:
	//
	// success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether the check is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1DFFD8C6-259E-582B-8B40-002C17DC****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s GetClusterDiagnosisCheckItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDiagnosisCheckItemsResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) GetCheckItems() []*GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	return s.CheckItems
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) SetCheckItems(v []*GetClusterDiagnosisCheckItemsResponseBodyCheckItems) *GetClusterDiagnosisCheckItemsResponseBody {
	s.CheckItems = v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) SetCode(v string) *GetClusterDiagnosisCheckItemsResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) SetIsSuccess(v bool) *GetClusterDiagnosisCheckItemsResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) SetRequestId(v string) *GetClusterDiagnosisCheckItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClusterDiagnosisCheckItemsResponseBodyCheckItems struct {
	// The description.
	//
	// example:
	//
	// Check whether the node can access host dns service
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The display name.
	//
	// example:
	//
	// HostDNS
	Display *string `json:"display,omitempty" xml:"display,omitempty"`
	// The name of the group to which the check item belongs.
	//
	// example:
	//
	// Node
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The severity level of the check result.
	//
	// Valid values:
	//
	// 	- normal
	//
	// 	- warning
	//
	// 	- error
	//
	// example:
	//
	// normal
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The check result.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// HostDNS
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The reference value.
	//
	// example:
	//
	// True
	Refer *string `json:"refer,omitempty" xml:"refer,omitempty"`
	// The value of the check item.
	//
	// example:
	//
	// True
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetClusterDiagnosisCheckItemsResponseBodyCheckItems) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GoString() string {
	return s.String()
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GetDesc() *string {
	return s.Desc
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GetDisplay() *string {
	return s.Display
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GetGroup() *string {
	return s.Group
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GetLevel() *string {
	return s.Level
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GetMessage() *string {
	return s.Message
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GetName() *string {
	return s.Name
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GetRefer() *string {
	return s.Refer
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) GetValue() *string {
	return s.Value
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) SetDesc(v string) *GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	s.Desc = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) SetDisplay(v string) *GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	s.Display = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) SetGroup(v string) *GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	s.Group = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) SetLevel(v string) *GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	s.Level = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) SetMessage(v string) *GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	s.Message = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) SetName(v string) *GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	s.Name = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) SetRefer(v string) *GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	s.Refer = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) SetValue(v string) *GetClusterDiagnosisCheckItemsResponseBodyCheckItems {
	s.Value = &v
	return s
}

func (s *GetClusterDiagnosisCheckItemsResponseBodyCheckItems) Validate() error {
	return dara.Validate(s)
}
