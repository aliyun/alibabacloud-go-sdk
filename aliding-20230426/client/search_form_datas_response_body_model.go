// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDatasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *SearchFormDatasResponseBody
	GetCurrentPage() *int32
	SetData(v []*SearchFormDatasResponseBodyData) *SearchFormDatasResponseBody
	GetData() []*SearchFormDatasResponseBodyData
	SetRequestId(v string) *SearchFormDatasResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *SearchFormDatasResponseBody
	GetTotalCount() *int32
	SetVendorRequestId(v string) *SearchFormDatasResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SearchFormDatasResponseBody
	GetVendorType() *string
}

type SearchFormDatasResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                             `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Data        []*SearchFormDatasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SearchFormDatasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *SearchFormDatasResponseBody) GetData() []*SearchFormDatasResponseBodyData {
	return s.Data
}

func (s *SearchFormDatasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchFormDatasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchFormDatasResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SearchFormDatasResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SearchFormDatasResponseBody) SetCurrentPage(v int32) *SearchFormDatasResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetData(v []*SearchFormDatasResponseBodyData) *SearchFormDatasResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDatasResponseBody) SetRequestId(v string) *SearchFormDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetTotalCount(v int32) *SearchFormDatasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetVendorRequestId(v string) *SearchFormDatasResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetVendorType(v string) *SearchFormDatasResponseBody {
	s.VendorType = &v
	return s
}

func (s *SearchFormDatasResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchFormDatasResponseBodyData struct {
	// example:
	//
	// 2018-01-24 11:22:01
	CreatedTimeGMT *string `json:"CreatedTimeGMT,omitempty" xml:"CreatedTimeGMT,omitempty"`
	// example:
	//
	// 012345
	CreatorUserId *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// example:
	//
	// 1002
	DataId   *int64                 `json:"DataId,omitempty" xml:"DataId,omitempty"`
	FormData map[string]interface{} `json:"FormData,omitempty" xml:"FormData,omitempty"`
	// example:
	//
	// FINST-BNKJDRF
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
	// example:
	//
	// FORM-EF6Y93URN24F1SCX15VA2P918LPEIJ2H3UFORCJ1
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// {}
	InstanceValue *string `json:"InstanceValue,omitempty" xml:"InstanceValue,omitempty"`
	// example:
	//
	// FORM-EF6Y93URN24F1SCX15VA2P918LPEIJ2H3UFORCJ1
	ModelUuid *string `json:"ModelUuid,omitempty" xml:"ModelUuid,omitempty"`
	// example:
	//
	// 2018-01-24 11:22:01
	ModifiedTimeGMT *string `json:"ModifiedTimeGMT,omitempty" xml:"ModifiedTimeGMT,omitempty"`
	// example:
	//
	// 012345
	ModifierUserId *string                                    `json:"ModifierUserId,omitempty" xml:"ModifierUserId,omitempty"`
	ModifyUser     *SearchFormDatasResponseBodyDataModifyUser `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty" type:"Struct"`
	Originator     *SearchFormDatasResponseBodyDataOriginator `json:"Originator,omitempty" xml:"Originator,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// 231008101012015353
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 3
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchFormDatasResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyData) GetCreatedTimeGMT() *string {
	return s.CreatedTimeGMT
}

func (s *SearchFormDatasResponseBodyData) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *SearchFormDatasResponseBodyData) GetDataId() *int64 {
	return s.DataId
}

func (s *SearchFormDatasResponseBodyData) GetFormData() map[string]interface{} {
	return s.FormData
}

func (s *SearchFormDatasResponseBodyData) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *SearchFormDatasResponseBodyData) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SearchFormDatasResponseBodyData) GetInstanceValue() *string {
	return s.InstanceValue
}

func (s *SearchFormDatasResponseBodyData) GetModelUuid() *string {
	return s.ModelUuid
}

func (s *SearchFormDatasResponseBodyData) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *SearchFormDatasResponseBodyData) GetModifierUserId() *string {
	return s.ModifierUserId
}

func (s *SearchFormDatasResponseBodyData) GetModifyUser() *SearchFormDatasResponseBodyDataModifyUser {
	return s.ModifyUser
}

func (s *SearchFormDatasResponseBodyData) GetOriginator() *SearchFormDatasResponseBodyDataOriginator {
	return s.Originator
}

func (s *SearchFormDatasResponseBodyData) GetSequence() *string {
	return s.Sequence
}

func (s *SearchFormDatasResponseBodyData) GetSerialNo() *string {
	return s.SerialNo
}

func (s *SearchFormDatasResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *SearchFormDatasResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *SearchFormDatasResponseBodyData) SetCreatedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.CreatedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetCreatorUserId(v string) *SearchFormDatasResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetDataId(v int64) *SearchFormDatasResponseBodyData {
	s.DataId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDatasResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormInstanceId(v string) *SearchFormDatasResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormUuid(v string) *SearchFormDatasResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetInstanceValue(v string) *SearchFormDatasResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModelUuid(v string) *SearchFormDatasResponseBodyData {
	s.ModelUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifierUserId(v string) *SearchFormDatasResponseBodyData {
	s.ModifierUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifyUser(v *SearchFormDatasResponseBodyDataModifyUser) *SearchFormDatasResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetOriginator(v *SearchFormDatasResponseBodyDataOriginator) *SearchFormDatasResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSequence(v string) *SearchFormDatasResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSerialNo(v string) *SearchFormDatasResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetTitle(v string) *SearchFormDatasResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetVersion(v int64) *SearchFormDatasResponseBodyData {
	s.Version = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SearchFormDatasResponseBodyDataModifyUser struct {
	// example:
	//
	// 012345
	UserId   *string                                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *SearchFormDatasResponseBodyDataModifyUserUserName `json:"UserName,omitempty" xml:"UserName,omitempty" type:"Struct"`
}

func (s SearchFormDatasResponseBodyDataModifyUser) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUser) GetUserId() *string {
	return s.UserId
}

func (s *SearchFormDatasResponseBodyDataModifyUser) GetUserName() *SearchFormDatasResponseBodyDataModifyUserUserName {
	return s.UserName
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserName(v *SearchFormDatasResponseBodyDataModifyUserUserName) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserName = v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) Validate() error {
	return dara.Validate(s)
}

type SearchFormDatasResponseBodyDataModifyUserUserName struct {
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// EngilishName
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
	// example:
	//
	// ZH-CN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataModifyUserUserName) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUserUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) GetType() *string {
	return s.Type
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetType(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.Type = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) Validate() error {
	return dara.Validate(s)
}

type SearchFormDatasResponseBodyDataOriginator struct {
	// example:
	//
	// 012345
	UserId   *string                                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *SearchFormDatasResponseBodyDataOriginatorUserName `json:"UserName,omitempty" xml:"UserName,omitempty" type:"Struct"`
}

func (s SearchFormDatasResponseBodyDataOriginator) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginator) GetUserId() *string {
	return s.UserId
}

func (s *SearchFormDatasResponseBodyDataOriginator) GetUserName() *SearchFormDatasResponseBodyDataOriginatorUserName {
	return s.UserName
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserId(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserName(v *SearchFormDatasResponseBodyDataOriginatorUserName) *SearchFormDatasResponseBodyDataOriginator {
	s.UserName = v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) Validate() error {
	return dara.Validate(s)
}

type SearchFormDatasResponseBodyDataOriginatorUserName struct {
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// EngilishName
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
	// example:
	//
	// ZH-CN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataOriginatorUserName) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginatorUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) GetType() *string {
	return s.Type
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetType(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.Type = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) Validate() error {
	return dara.Validate(s)
}
