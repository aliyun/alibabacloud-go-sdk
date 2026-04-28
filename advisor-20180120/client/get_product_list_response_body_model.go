// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetProductListResponseBodyAccessDeniedDetail) *GetProductListResponseBody
	GetAccessDeniedDetail() *GetProductListResponseBodyAccessDeniedDetail
	SetCode(v string) *GetProductListResponseBody
	GetCode() *string
	SetData(v []*GetProductListResponseBodyData) *GetProductListResponseBody
	GetData() []*GetProductListResponseBodyData
	SetMessage(v string) *GetProductListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProductListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProductListResponseBody
	GetSuccess() *bool
	SetUserMessage(v string) *GetProductListResponseBody
	GetUserMessage() *string
}

type GetProductListResponseBody struct {
	AccessDeniedDetail *GetProductListResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetProductListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// *
	UserMessage *string `json:"UserMessage,omitempty" xml:"UserMessage,omitempty"`
}

func (s GetProductListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBody) GetAccessDeniedDetail() *GetProductListResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetProductListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProductListResponseBody) GetData() []*GetProductListResponseBodyData {
	return s.Data
}

func (s *GetProductListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProductListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProductListResponseBody) GetUserMessage() *string {
	return s.UserMessage
}

func (s *GetProductListResponseBody) SetAccessDeniedDetail(v *GetProductListResponseBodyAccessDeniedDetail) *GetProductListResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetProductListResponseBody) SetCode(v string) *GetProductListResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductListResponseBody) SetData(v []*GetProductListResponseBodyData) *GetProductListResponseBody {
	s.Data = v
	return s
}

func (s *GetProductListResponseBody) SetMessage(v string) *GetProductListResponseBody {
	s.Message = &v
	return s
}

func (s *GetProductListResponseBody) SetRequestId(v string) *GetProductListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductListResponseBody) SetSuccess(v bool) *GetProductListResponseBody {
	s.Success = &v
	return s
}

func (s *GetProductListResponseBody) SetUserMessage(v string) *GetProductListResponseBody {
	s.UserMessage = &v
	return s
}

func (s *GetProductListResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
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

type GetProductListResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// *
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// *
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// *
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// *
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// ****
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// *
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// *
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetProductListResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetProductListResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetProductListResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetProductListResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetProductListResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetProductListResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetProductListResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetProductListResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetProductListResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetProductListResponseBodyData struct {
	Category    *string                                      `json:"Category,omitempty" xml:"Category,omitempty"`
	ProductList []*GetProductListResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
}

func (s GetProductListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *GetProductListResponseBodyData) GetProductList() []*GetProductListResponseBodyDataProductList {
	return s.ProductList
}

func (s *GetProductListResponseBodyData) SetCategory(v string) *GetProductListResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetProductListResponseBodyData) SetProductList(v []*GetProductListResponseBodyDataProductList) *GetProductListResponseBodyData {
	s.ProductList = v
	return s
}

func (s *GetProductListResponseBodyData) Validate() error {
	if s.ProductList != nil {
		for _, item := range s.ProductList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProductListResponseBodyDataProductList struct {
	NewLabel *string `json:"NewLabel,omitempty" xml:"NewLabel,omitempty"`
	// example:
	//
	// hologres
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetProductListResponseBodyDataProductList) String() string {
	return dara.Prettify(s)
}

func (s GetProductListResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *GetProductListResponseBodyDataProductList) GetNewLabel() *string {
	return s.NewLabel
}

func (s *GetProductListResponseBodyDataProductList) GetProduct() *string {
	return s.Product
}

func (s *GetProductListResponseBodyDataProductList) SetNewLabel(v string) *GetProductListResponseBodyDataProductList {
	s.NewLabel = &v
	return s
}

func (s *GetProductListResponseBodyDataProductList) SetProduct(v string) *GetProductListResponseBodyDataProductList {
	s.Product = &v
	return s
}

func (s *GetProductListResponseBodyDataProductList) Validate() error {
	return dara.Validate(s)
}
