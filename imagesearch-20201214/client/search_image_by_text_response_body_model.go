// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SearchImageByTextResponseBodyAccessDeniedDetail) *SearchImageByTextResponseBody
	GetAccessDeniedDetail() *SearchImageByTextResponseBodyAccessDeniedDetail
	SetAuctions(v []*SearchImageByTextResponseBodyAuctions) *SearchImageByTextResponseBody
	GetAuctions() []*SearchImageByTextResponseBodyAuctions
	SetCode(v int32) *SearchImageByTextResponseBody
	GetCode() *int32
	SetHead(v *SearchImageByTextResponseBodyHead) *SearchImageByTextResponseBody
	GetHead() *SearchImageByTextResponseBodyHead
	SetMsg(v string) *SearchImageByTextResponseBody
	GetMsg() *string
	SetPicInfo(v *SearchImageByTextResponseBodyPicInfo) *SearchImageByTextResponseBody
	GetPicInfo() *SearchImageByTextResponseBodyPicInfo
	SetRequestId(v string) *SearchImageByTextResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchImageByTextResponseBody
	GetSuccess() *bool
}

type SearchImageByTextResponseBody struct {
	AccessDeniedDetail *SearchImageByTextResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Auctions           []*SearchImageByTextResponseBodyAuctions         `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Head *SearchImageByTextResponseBodyHead `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	// example:
	//
	// success
	Msg     *string                               `json:"Msg,omitempty" xml:"Msg,omitempty"`
	PicInfo *SearchImageByTextResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByTextResponseBody) GetAccessDeniedDetail() *SearchImageByTextResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SearchImageByTextResponseBody) GetAuctions() []*SearchImageByTextResponseBodyAuctions {
	return s.Auctions
}

func (s *SearchImageByTextResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SearchImageByTextResponseBody) GetHead() *SearchImageByTextResponseBodyHead {
	return s.Head
}

func (s *SearchImageByTextResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SearchImageByTextResponseBody) GetPicInfo() *SearchImageByTextResponseBodyPicInfo {
	return s.PicInfo
}

func (s *SearchImageByTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchImageByTextResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchImageByTextResponseBody) SetAccessDeniedDetail(v *SearchImageByTextResponseBodyAccessDeniedDetail) *SearchImageByTextResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SearchImageByTextResponseBody) SetAuctions(v []*SearchImageByTextResponseBodyAuctions) *SearchImageByTextResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByTextResponseBody) SetCode(v int32) *SearchImageByTextResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByTextResponseBody) SetHead(v *SearchImageByTextResponseBodyHead) *SearchImageByTextResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByTextResponseBody) SetMsg(v string) *SearchImageByTextResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByTextResponseBody) SetPicInfo(v *SearchImageByTextResponseBodyPicInfo) *SearchImageByTextResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchImageByTextResponseBody) SetRequestId(v string) *SearchImageByTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByTextResponseBody) SetSuccess(v bool) *SearchImageByTextResponseBody {
	s.Success = &v
	return s
}

func (s *SearchImageByTextResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Auctions != nil {
		for _, item := range s.Auctions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Head != nil {
		if err := s.Head.Validate(); err != nil {
			return err
		}
	}
	if s.PicInfo != nil {
		if err := s.PicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchImageByTextResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 111
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 222
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// xxxxxx
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// PauseNotify
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SearchImageByTextResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SearchImageByTextResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SearchImageByTextResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SearchImageByTextResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SearchImageByTextResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SearchImageByTextResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SearchImageByTextResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SearchImageByTextResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SearchImageByTextResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type SearchImageByTextResponseBodyAuctions struct {
	// example:
	//
	// 8888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// example:
	//
	// 2
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	// example:
	//
	// 2
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	// example:
	//
	// 2
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// xxxx
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// example:
	//
	// xxxx
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	// example:
	//
	// xxxx
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	// example:
	//
	// xxxx
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s SearchImageByTextResponseBodyAuctions) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByTextResponseBodyAuctions) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *SearchImageByTextResponseBodyAuctions) GetCustomContent() *string {
	return s.CustomContent
}

func (s *SearchImageByTextResponseBodyAuctions) GetIntAttr() *int32 {
	return s.IntAttr
}

func (s *SearchImageByTextResponseBodyAuctions) GetIntAttr2() *int32 {
	return s.IntAttr2
}

func (s *SearchImageByTextResponseBodyAuctions) GetIntAttr3() *int32 {
	return s.IntAttr3
}

func (s *SearchImageByTextResponseBodyAuctions) GetIntAttr4() *int32 {
	return s.IntAttr4
}

func (s *SearchImageByTextResponseBodyAuctions) GetPicName() *string {
	return s.PicName
}

func (s *SearchImageByTextResponseBodyAuctions) GetProductId() *string {
	return s.ProductId
}

func (s *SearchImageByTextResponseBodyAuctions) GetScore() *float32 {
	return s.Score
}

func (s *SearchImageByTextResponseBodyAuctions) GetStrAttr() *string {
	return s.StrAttr
}

func (s *SearchImageByTextResponseBodyAuctions) GetStrAttr2() *string {
	return s.StrAttr2
}

func (s *SearchImageByTextResponseBodyAuctions) GetStrAttr3() *string {
	return s.StrAttr3
}

func (s *SearchImageByTextResponseBodyAuctions) GetStrAttr4() *string {
	return s.StrAttr4
}

func (s *SearchImageByTextResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByTextResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetCustomContent(v string) *SearchImageByTextResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByTextResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetIntAttr2(v int32) *SearchImageByTextResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetIntAttr3(v int32) *SearchImageByTextResponseBodyAuctions {
	s.IntAttr3 = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetIntAttr4(v int32) *SearchImageByTextResponseBodyAuctions {
	s.IntAttr4 = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetPicName(v string) *SearchImageByTextResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetProductId(v string) *SearchImageByTextResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetScore(v float32) *SearchImageByTextResponseBodyAuctions {
	s.Score = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetStrAttr(v string) *SearchImageByTextResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetStrAttr2(v string) *SearchImageByTextResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetStrAttr3(v string) *SearchImageByTextResponseBodyAuctions {
	s.StrAttr3 = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) SetStrAttr4(v string) *SearchImageByTextResponseBodyAuctions {
	s.StrAttr4 = &v
	return s
}

func (s *SearchImageByTextResponseBodyAuctions) Validate() error {
	return dara.Validate(s)
}

type SearchImageByTextResponseBodyHead struct {
	// example:
	//
	// 10
	DocsFound *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	// example:
	//
	// 10000
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	// example:
	//
	// 95
	SearchTime *int32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
}

func (s SearchImageByTextResponseBodyHead) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchImageByTextResponseBodyHead) GetDocsFound() *int32 {
	return s.DocsFound
}

func (s *SearchImageByTextResponseBodyHead) GetDocsReturn() *int32 {
	return s.DocsReturn
}

func (s *SearchImageByTextResponseBodyHead) GetSearchTime() *int32 {
	return s.SearchTime
}

func (s *SearchImageByTextResponseBodyHead) SetDocsFound(v int32) *SearchImageByTextResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchImageByTextResponseBodyHead) SetDocsReturn(v int32) *SearchImageByTextResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchImageByTextResponseBodyHead) SetSearchTime(v int32) *SearchImageByTextResponseBodyHead {
	s.SearchTime = &v
	return s
}

func (s *SearchImageByTextResponseBodyHead) Validate() error {
	return dara.Validate(s)
}

type SearchImageByTextResponseBodyPicInfo struct {
	AllCategories []*SearchImageByTextResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
}

func (s SearchImageByTextResponseBodyPicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByTextResponseBodyPicInfo) GetAllCategories() []*SearchImageByTextResponseBodyPicInfoAllCategories {
	return s.AllCategories
}

func (s *SearchImageByTextResponseBodyPicInfo) SetAllCategories(v []*SearchImageByTextResponseBodyPicInfoAllCategories) *SearchImageByTextResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *SearchImageByTextResponseBodyPicInfo) Validate() error {
	if s.AllCategories != nil {
		for _, item := range s.AllCategories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchImageByTextResponseBodyPicInfoAllCategories struct {
	// example:
	//
	// 88888888
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// other
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchImageByTextResponseBodyPicInfoAllCategories) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByTextResponseBodyPicInfoAllCategories) GetId() *int32 {
	return s.Id
}

func (s *SearchImageByTextResponseBodyPicInfoAllCategories) GetName() *string {
	return s.Name
}

func (s *SearchImageByTextResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByTextResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *SearchImageByTextResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByTextResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

func (s *SearchImageByTextResponseBodyPicInfoAllCategories) Validate() error {
	return dara.Validate(s)
}
