// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SearchFormDataSecondGenerationResponseBodyData) *SearchFormDataSecondGenerationResponseBody
	GetData() []*SearchFormDataSecondGenerationResponseBodyData
	SetPageNumber(v int64) *SearchFormDataSecondGenerationResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *SearchFormDataSecondGenerationResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *SearchFormDataSecondGenerationResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *SearchFormDataSecondGenerationResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SearchFormDataSecondGenerationResponseBody
	GetVendorType() *string
}

type SearchFormDataSecondGenerationResponseBody struct {
	Data []*SearchFormDataSecondGenerationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBody) GetData() []*SearchFormDataSecondGenerationResponseBodyData {
	return s.Data
}

func (s *SearchFormDataSecondGenerationResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchFormDataSecondGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchFormDataSecondGenerationResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchFormDataSecondGenerationResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SearchFormDataSecondGenerationResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SearchFormDataSecondGenerationResponseBody) SetData(v []*SearchFormDataSecondGenerationResponseBodyData) *SearchFormDataSecondGenerationResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetPageNumber(v int64) *SearchFormDataSecondGenerationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetRequestId(v string) *SearchFormDataSecondGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetTotalCount(v int64) *SearchFormDataSecondGenerationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetVendorRequestId(v string) *SearchFormDataSecondGenerationResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetVendorType(v string) *SearchFormDataSecondGenerationResponseBody {
	s.VendorType = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchFormDataSecondGenerationResponseBodyData struct {
	// example:
	//
	// 2021-05-01 10:10:10
	CreateTimeGMT *string `json:"CreateTimeGMT,omitempty" xml:"CreateTimeGMT,omitempty"`
	// example:
	//
	// 012345
	CreatorUserId *string                `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	FormData      map[string]interface{} `json:"FormData,omitempty" xml:"FormData,omitempty"`
	// example:
	//
	// FINST-xxxx
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
	// example:
	//
	// FORM-xxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// 1023
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {}
	InstanceValue *string `json:"InstanceValue,omitempty" xml:"InstanceValue,omitempty"`
	// example:
	//
	// 2021-05-01 10:10:10
	ModifiedTimeGMT *string                                                   `json:"ModifiedTimeGMT,omitempty" xml:"ModifiedTimeGMT,omitempty"`
	Modifier        *string                                                   `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	ModifyUser      *SearchFormDataSecondGenerationResponseBodyDataModifyUser `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty" type:"Struct"`
	Originator      *SearchFormDataSecondGenerationResponseBodyDataOriginator `json:"Originator,omitempty" xml:"Originator,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// sagc1b3090d
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 3
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetFormData() map[string]interface{} {
	return s.FormData
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetInstanceValue() *string {
	return s.InstanceValue
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetModifyUser() *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	return s.ModifyUser
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetOriginator() *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	return s.Originator
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetSequence() *string {
	return s.Sequence
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *SearchFormDataSecondGenerationResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetCreateTimeGMT(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetCreatorUserId(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormInstanceId(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormUuid(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetId(v int64) *SearchFormDataSecondGenerationResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetInstanceValue(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifier(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifyUser(v *SearchFormDataSecondGenerationResponseBodyDataModifyUser) *SearchFormDataSecondGenerationResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetOriginator(v *SearchFormDataSecondGenerationResponseBodyDataOriginator) *SearchFormDataSecondGenerationResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetSequence(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetSerialNumber(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetTitle(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetVersion(v int64) *SearchFormDataSecondGenerationResponseBodyData {
	s.Version = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SearchFormDataSecondGenerationResponseBodyDataModifyUser struct {
	Name *SearchFormDataSecondGenerationResponseBodyDataModifyUserName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUser) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) GetName() *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	return s.Name
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) GetUserId() *string {
	return s.UserId
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) SetName(v *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) Validate() error {
	return dara.Validate(s)
}

type SearchFormDataSecondGenerationResponseBodyDataModifyUserName struct {
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// English
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUserName) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) SetNameInChinese(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) Validate() error {
	return dara.Validate(s)
}

type SearchFormDataSecondGenerationResponseBodyDataOriginator struct {
	Name *SearchFormDataSecondGenerationResponseBodyDataOriginatorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginator) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) GetName() *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	return s.Name
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) GetUserId() *string {
	return s.UserId
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) SetName(v *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) SetUserId(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) Validate() error {
	return dara.Validate(s)
}

type SearchFormDataSecondGenerationResponseBodyDataOriginatorName struct {
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// English
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginatorName) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) SetNameInChinese(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) Validate() error {
	return dara.Validate(s)
}
