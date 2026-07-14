// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindDmAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListBindDmAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListBindDmAccountResponseBody
	GetCode() *string
	SetData(v []*ListBindDmAccountResponseBodyData) *ListBindDmAccountResponseBody
	GetData() []*ListBindDmAccountResponseBodyData
	SetMessage(v string) *ListBindDmAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBindDmAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBindDmAccountResponseBody
	GetSuccess() *bool
}

type ListBindDmAccountResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListBindDmAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxx-xx**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBindDmAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBindDmAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindDmAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListBindDmAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBindDmAccountResponseBody) GetData() []*ListBindDmAccountResponseBodyData {
	return s.Data
}

func (s *ListBindDmAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBindDmAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBindDmAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBindDmAccountResponseBody) SetAccessDeniedDetail(v string) *ListBindDmAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListBindDmAccountResponseBody) SetCode(v string) *ListBindDmAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ListBindDmAccountResponseBody) SetData(v []*ListBindDmAccountResponseBodyData) *ListBindDmAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListBindDmAccountResponseBody) SetMessage(v string) *ListBindDmAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ListBindDmAccountResponseBody) SetRequestId(v string) *ListBindDmAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindDmAccountResponseBody) SetSuccess(v bool) *ListBindDmAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ListBindDmAccountResponseBody) Validate() error {
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

type ListBindDmAccountResponseBodyData struct {
	// example:
	//
	// xx@xx.com
	AccountCode *string                                      `json:"AccountCode,omitempty" xml:"AccountCode,omitempty"`
	ExtendAttr  *ListBindDmAccountResponseBodyDataExtendAttr `json:"ExtendAttr,omitempty" xml:"ExtendAttr,omitempty" type:"Struct"`
	// example:
	//
	// cams-*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ins
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ListBindDmAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBindDmAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBindDmAccountResponseBodyData) GetAccountCode() *string {
	return s.AccountCode
}

func (s *ListBindDmAccountResponseBodyData) GetExtendAttr() *ListBindDmAccountResponseBodyDataExtendAttr {
	return s.ExtendAttr
}

func (s *ListBindDmAccountResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBindDmAccountResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListBindDmAccountResponseBodyData) SetAccountCode(v string) *ListBindDmAccountResponseBodyData {
	s.AccountCode = &v
	return s
}

func (s *ListBindDmAccountResponseBodyData) SetExtendAttr(v *ListBindDmAccountResponseBodyDataExtendAttr) *ListBindDmAccountResponseBodyData {
	s.ExtendAttr = v
	return s
}

func (s *ListBindDmAccountResponseBodyData) SetInstanceId(v string) *ListBindDmAccountResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListBindDmAccountResponseBodyData) SetInstanceName(v string) *ListBindDmAccountResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListBindDmAccountResponseBodyData) Validate() error {
	if s.ExtendAttr != nil {
		if err := s.ExtendAttr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBindDmAccountResponseBodyDataExtendAttr struct {
	// example:
	//
	// xx@xx.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// batch
	SendType *string `json:"SendType,omitempty" xml:"SendType,omitempty"`
}

func (s ListBindDmAccountResponseBodyDataExtendAttr) String() string {
	return dara.Prettify(s)
}

func (s ListBindDmAccountResponseBodyDataExtendAttr) GoString() string {
	return s.String()
}

func (s *ListBindDmAccountResponseBodyDataExtendAttr) GetAccountName() *string {
	return s.AccountName
}

func (s *ListBindDmAccountResponseBodyDataExtendAttr) GetSendType() *string {
	return s.SendType
}

func (s *ListBindDmAccountResponseBodyDataExtendAttr) SetAccountName(v string) *ListBindDmAccountResponseBodyDataExtendAttr {
	s.AccountName = &v
	return s
}

func (s *ListBindDmAccountResponseBodyDataExtendAttr) SetSendType(v string) *ListBindDmAccountResponseBodyDataExtendAttr {
	s.SendType = &v
	return s
}

func (s *ListBindDmAccountResponseBodyDataExtendAttr) Validate() error {
	return dara.Validate(s)
}
