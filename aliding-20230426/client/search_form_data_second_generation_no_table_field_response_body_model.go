// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationNoTableFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SearchFormDataSecondGenerationNoTableFieldResponseBodyData) *SearchFormDataSecondGenerationNoTableFieldResponseBody
	GetData() []*SearchFormDataSecondGenerationNoTableFieldResponseBodyData
	SetPageNumber(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBody
	GetVendorType() *string
}

type SearchFormDataSecondGenerationNoTableFieldResponseBody struct {
	Data []*SearchFormDataSecondGenerationNoTableFieldResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s SearchFormDataSecondGenerationNoTableFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) GetData() []*SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	return s.Data
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetData(v []*SearchFormDataSecondGenerationNoTableFieldResponseBodyData) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetPageNumber(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetRequestId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetTotalCount(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetVendorRequestId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetVendorType(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.VendorType = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyData struct {
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
	// 54114
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {}
	InstanceValue *string `json:"InstanceValue,omitempty" xml:"InstanceValue,omitempty"`
	// example:
	//
	// 2021-05-01 10:10:10
	ModifiedTimeGMT *string                                                               `json:"ModifiedTimeGMT,omitempty" xml:"ModifiedTimeGMT,omitempty"`
	Modifier        *string                                                               `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	ModifyUser      *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty" type:"Struct"`
	Originator      *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator `json:"Originator,omitempty" xml:"Originator,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// 7CE737P1SS
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 3
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetFormData() map[string]interface{} {
	return s.FormData
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetInstanceValue() *string {
	return s.InstanceValue
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetModifyUser() *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser {
	return s.ModifyUser
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetOriginator() *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator {
	return s.Originator
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetSequence() *string {
	return s.Sequence
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetCreateTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetCreatorUserId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetFormInstanceId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetFormUuid(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetId(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetInstanceValue(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetModifier(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetModifyUser(v *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetOriginator(v *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetSequence(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetSerialNumber(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetTitle(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetVersion(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Version = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) Validate() error {
	if s.ModifyUser != nil {
		if err := s.ModifyUser.Validate(); err != nil {
			return err
		}
	}
	if s.Originator != nil {
		if err := s.Originator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser struct {
	Name *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) GetName() *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName {
	return s.Name
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) GetUserId() *string {
	return s.UserId
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) SetName(v *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName struct {
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// English
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) SetNameInChinese(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) Validate() error {
	return dara.Validate(s)
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator struct {
	Name *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) GetName() *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName {
	return s.Name
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) GetUserId() *string {
	return s.UserId
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) SetName(v *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) SetUserId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName struct {
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// English
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) SetNameInChinese(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) Validate() error {
	return dara.Validate(s)
}
