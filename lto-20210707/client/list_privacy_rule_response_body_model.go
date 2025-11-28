// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivacyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPrivacyRuleResponseBody
	GetCode() *string
	SetData(v *ListPrivacyRuleResponseBodyData) *ListPrivacyRuleResponseBody
	GetData() *ListPrivacyRuleResponseBodyData
	SetHttpStatusCode(v int32) *ListPrivacyRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPrivacyRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrivacyRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPrivacyRuleResponseBody
	GetSuccess() *bool
}

type ListPrivacyRuleResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListPrivacyRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPrivacyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPrivacyRuleResponseBody) GetData() *ListPrivacyRuleResponseBodyData {
	return s.Data
}

func (s *ListPrivacyRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPrivacyRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrivacyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivacyRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPrivacyRuleResponseBody) SetCode(v string) *ListPrivacyRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrivacyRuleResponseBody) SetData(v *ListPrivacyRuleResponseBodyData) *ListPrivacyRuleResponseBody {
	s.Data = v
	return s
}

func (s *ListPrivacyRuleResponseBody) SetHttpStatusCode(v int32) *ListPrivacyRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPrivacyRuleResponseBody) SetMessage(v string) *ListPrivacyRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrivacyRuleResponseBody) SetRequestId(v string) *ListPrivacyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivacyRuleResponseBody) SetSuccess(v bool) *ListPrivacyRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ListPrivacyRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPrivacyRuleResponseBodyData struct {
	Num      *int32                                     `json:"Num,omitempty" xml:"Num,omitempty"`
	PageData []*ListPrivacyRuleResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	Size     *int32                                     `json:"Size,omitempty" xml:"Size,omitempty"`
	Total    *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPrivacyRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleResponseBodyData) GetNum() *int32 {
	return s.Num
}

func (s *ListPrivacyRuleResponseBodyData) GetPageData() []*ListPrivacyRuleResponseBodyDataPageData {
	return s.PageData
}

func (s *ListPrivacyRuleResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListPrivacyRuleResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListPrivacyRuleResponseBodyData) SetNum(v int32) *ListPrivacyRuleResponseBodyData {
	s.Num = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyData) SetPageData(v []*ListPrivacyRuleResponseBodyDataPageData) *ListPrivacyRuleResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListPrivacyRuleResponseBodyData) SetSize(v int32) *ListPrivacyRuleResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyData) SetTotal(v int32) *ListPrivacyRuleResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivacyRuleResponseBodyDataPageData struct {
	AlgImpl       *string `json:"AlgImpl,omitempty" xml:"AlgImpl,omitempty"`
	AlgType       *string `json:"AlgType,omitempty" xml:"AlgType,omitempty"`
	CurrentUser   *bool   `json:"CurrentUser,omitempty" xml:"CurrentUser,omitempty"`
	MemberName    *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PrivacyRuleId *string `json:"PrivacyRuleId,omitempty" xml:"PrivacyRuleId,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPrivacyRuleResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleResponseBodyDataPageData) GetAlgImpl() *string {
	return s.AlgImpl
}

func (s *ListPrivacyRuleResponseBodyDataPageData) GetAlgType() *string {
	return s.AlgType
}

func (s *ListPrivacyRuleResponseBodyDataPageData) GetCurrentUser() *bool {
	return s.CurrentUser
}

func (s *ListPrivacyRuleResponseBodyDataPageData) GetMemberName() *string {
	return s.MemberName
}

func (s *ListPrivacyRuleResponseBodyDataPageData) GetName() *string {
	return s.Name
}

func (s *ListPrivacyRuleResponseBodyDataPageData) GetPrivacyRuleId() *string {
	return s.PrivacyRuleId
}

func (s *ListPrivacyRuleResponseBodyDataPageData) GetRemark() *string {
	return s.Remark
}

func (s *ListPrivacyRuleResponseBodyDataPageData) GetStatus() *string {
	return s.Status
}

func (s *ListPrivacyRuleResponseBodyDataPageData) SetAlgImpl(v string) *ListPrivacyRuleResponseBodyDataPageData {
	s.AlgImpl = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyDataPageData) SetAlgType(v string) *ListPrivacyRuleResponseBodyDataPageData {
	s.AlgType = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyDataPageData) SetCurrentUser(v bool) *ListPrivacyRuleResponseBodyDataPageData {
	s.CurrentUser = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyDataPageData) SetMemberName(v string) *ListPrivacyRuleResponseBodyDataPageData {
	s.MemberName = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyDataPageData) SetName(v string) *ListPrivacyRuleResponseBodyDataPageData {
	s.Name = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyDataPageData) SetPrivacyRuleId(v string) *ListPrivacyRuleResponseBodyDataPageData {
	s.PrivacyRuleId = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyDataPageData) SetRemark(v string) *ListPrivacyRuleResponseBodyDataPageData {
	s.Remark = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyDataPageData) SetStatus(v string) *ListPrivacyRuleResponseBodyDataPageData {
	s.Status = &v
	return s
}

func (s *ListPrivacyRuleResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}
