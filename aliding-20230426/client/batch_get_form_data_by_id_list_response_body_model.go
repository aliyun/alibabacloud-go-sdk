// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetFormDataByIdListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchGetFormDataByIdListResponseBody
	GetRequestId() *string
	SetResult(v []*BatchGetFormDataByIdListResponseBodyResult) *BatchGetFormDataByIdListResponseBody
	GetResult() []*BatchGetFormDataByIdListResponseBodyResult
	SetVendorRequestId(v string) *BatchGetFormDataByIdListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *BatchGetFormDataByIdListResponseBody
	GetVendorType() *string
}

type BatchGetFormDataByIdListResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// [ "FINST-SASNOO39NSIFF780" ]
	Result []*BatchGetFormDataByIdListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetFormDataByIdListResponseBody) GetResult() []*BatchGetFormDataByIdListResponseBodyResult {
	return s.Result
}

func (s *BatchGetFormDataByIdListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *BatchGetFormDataByIdListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *BatchGetFormDataByIdListResponseBody) SetRequestId(v string) *BatchGetFormDataByIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBody) SetResult(v []*BatchGetFormDataByIdListResponseBodyResult) *BatchGetFormDataByIdListResponseBody {
	s.Result = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBody) SetVendorRequestId(v string) *BatchGetFormDataByIdListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBody) SetVendorType(v string) *BatchGetFormDataByIdListResponseBody {
	s.VendorType = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetFormDataByIdListResponseBodyResult struct {
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
	// 21044829126
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {}
	InstanceValue *string `json:"InstanceValue,omitempty" xml:"InstanceValue,omitempty"`
	// example:
	//
	// 2021-05-01 10:10:10
	ModifiedTimeGMT *string                                               `json:"ModifiedTimeGMT,omitempty" xml:"ModifiedTimeGMT,omitempty"`
	Modifier        *string                                               `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	ModifyUser      *BatchGetFormDataByIdListResponseBodyResultModifyUser `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty" type:"Struct"`
	Originator      *BatchGetFormDataByIdListResponseBodyResultOriginator `json:"Originator,omitempty" xml:"Originator,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// example:
	//
	// SA65776
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetFormData() map[string]interface{} {
	return s.FormData
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetFormUuid() *string {
	return s.FormUuid
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetInstanceValue() *string {
	return s.InstanceValue
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetModifiedTimeGMT() *string {
	return s.ModifiedTimeGMT
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetModifier() *string {
	return s.Modifier
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetModifyUser() *BatchGetFormDataByIdListResponseBodyResultModifyUser {
	return s.ModifyUser
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetOriginator() *BatchGetFormDataByIdListResponseBodyResultOriginator {
	return s.Originator
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetSequence() *string {
	return s.Sequence
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *BatchGetFormDataByIdListResponseBodyResult) GetVersion() *int64 {
	return s.Version
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetCreateTimeGMT(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.CreateTimeGMT = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetCreatorUserId(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetFormData(v map[string]interface{}) *BatchGetFormDataByIdListResponseBodyResult {
	s.FormData = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetFormInstanceId(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.FormInstanceId = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetFormUuid(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetId(v int64) *BatchGetFormDataByIdListResponseBodyResult {
	s.Id = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetInstanceValue(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.InstanceValue = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetModifiedTimeGMT(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetModifier(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.Modifier = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetModifyUser(v *BatchGetFormDataByIdListResponseBodyResultModifyUser) *BatchGetFormDataByIdListResponseBodyResult {
	s.ModifyUser = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetOriginator(v *BatchGetFormDataByIdListResponseBodyResultOriginator) *BatchGetFormDataByIdListResponseBodyResult {
	s.Originator = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetSequence(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.Sequence = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetSerialNumber(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetTitle(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetVersion(v int64) *BatchGetFormDataByIdListResponseBodyResult {
	s.Version = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) Validate() error {
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

type BatchGetFormDataByIdListResponseBodyResultModifyUser struct {
	Name *BatchGetFormDataByIdListResponseBodyResultModifyUserName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResultModifyUser) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResultModifyUser) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUser) GetName() *BatchGetFormDataByIdListResponseBodyResultModifyUserName {
	return s.Name
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUser) GetUserId() *string {
	return s.UserId
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUser) SetName(v *BatchGetFormDataByIdListResponseBodyResultModifyUserName) *BatchGetFormDataByIdListResponseBodyResultModifyUser {
	s.Name = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUser) SetUserId(v string) *BatchGetFormDataByIdListResponseBodyResultModifyUser {
	s.UserId = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUser) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetFormDataByIdListResponseBodyResultModifyUserName struct {
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// EnglishName
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResultModifyUserName) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResultModifyUserName) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUserName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUserName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUserName) SetNameInChinese(v string) *BatchGetFormDataByIdListResponseBodyResultModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUserName) SetNameInEnglish(v string) *BatchGetFormDataByIdListResponseBodyResultModifyUserName {
	s.NameInEnglish = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUserName) Validate() error {
	return dara.Validate(s)
}

type BatchGetFormDataByIdListResponseBodyResultOriginator struct {
	Name *BatchGetFormDataByIdListResponseBodyResultOriginatorName `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResultOriginator) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResultOriginator) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginator) GetName() *BatchGetFormDataByIdListResponseBodyResultOriginatorName {
	return s.Name
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginator) GetUserId() *string {
	return s.UserId
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginator) SetName(v *BatchGetFormDataByIdListResponseBodyResultOriginatorName) *BatchGetFormDataByIdListResponseBodyResultOriginator {
	s.Name = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginator) SetUserId(v string) *BatchGetFormDataByIdListResponseBodyResultOriginator {
	s.UserId = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginator) Validate() error {
	if s.Name != nil {
		if err := s.Name.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetFormDataByIdListResponseBodyResultOriginatorName struct {
	NameInChinese *string `json:"NameInChinese,omitempty" xml:"NameInChinese,omitempty"`
	// example:
	//
	// EnglishName
	NameInEnglish *string `json:"NameInEnglish,omitempty" xml:"NameInEnglish,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResultOriginatorName) String() string {
	return dara.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResultOriginatorName) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginatorName) GetNameInChinese() *string {
	return s.NameInChinese
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginatorName) GetNameInEnglish() *string {
	return s.NameInEnglish
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginatorName) SetNameInChinese(v string) *BatchGetFormDataByIdListResponseBodyResultOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginatorName) SetNameInEnglish(v string) *BatchGetFormDataByIdListResponseBodyResultOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginatorName) Validate() error {
	return dara.Validate(s)
}
