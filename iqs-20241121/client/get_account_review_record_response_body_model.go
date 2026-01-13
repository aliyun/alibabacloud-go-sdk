// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountReviewRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *GetAccountReviewRecordResponseBody
	GetAddress() *string
	SetAliUid(v string) *GetAccountReviewRecordResponseBody
	GetAliUid() *string
	SetApplyType(v string) *GetAccountReviewRecordResponseBody
	GetApplyType() *string
	SetContactName(v string) *GetAccountReviewRecordResponseBody
	GetContactName() *string
	SetInstanceId(v string) *GetAccountReviewRecordResponseBody
	GetInstanceId() *string
	SetPhone(v string) *GetAccountReviewRecordResponseBody
	GetPhone() *string
	SetProductName(v string) *GetAccountReviewRecordResponseBody
	GetProductName() *string
	SetQps(v int32) *GetAccountReviewRecordResponseBody
	GetQps() *int32
	SetRequestId(v string) *GetAccountReviewRecordResponseBody
	GetRequestId() *string
	SetSceneDesc(v string) *GetAccountReviewRecordResponseBody
	GetSceneDesc() *string
	SetScopes(v []*string) *GetAccountReviewRecordResponseBody
	GetScopes() []*string
	SetServiceType(v string) *GetAccountReviewRecordResponseBody
	GetServiceType() *string
}

type GetAccountReviewRecordResponseBody struct {
	// example:
	//
	// 杭州
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 123123213123
	AliUid *string `json:"aliUid,omitempty" xml:"aliUid,omitempty"`
	// example:
	//
	// 申请服务
	ApplyType *string `json:"applyType,omitempty" xml:"applyType,omitempty"`
	// example:
	//
	// 18987236721
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	// example:
	//
	// 测试商品
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 123214889322
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// 30
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 张三
	Qps *int32 `json:"qps,omitempty" xml:"qps,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6788a2c2-157d4ebe-ad979cd4f296
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 测试
	SceneDesc *string   `json:"sceneDesc,omitempty" xml:"sceneDesc,omitempty"`
	Scopes    []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	// example:
	//
	// 测试
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s GetAccountReviewRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountReviewRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountReviewRecordResponseBody) GetAddress() *string {
	return s.Address
}

func (s *GetAccountReviewRecordResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *GetAccountReviewRecordResponseBody) GetApplyType() *string {
	return s.ApplyType
}

func (s *GetAccountReviewRecordResponseBody) GetContactName() *string {
	return s.ContactName
}

func (s *GetAccountReviewRecordResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAccountReviewRecordResponseBody) GetPhone() *string {
	return s.Phone
}

func (s *GetAccountReviewRecordResponseBody) GetProductName() *string {
	return s.ProductName
}

func (s *GetAccountReviewRecordResponseBody) GetQps() *int32 {
	return s.Qps
}

func (s *GetAccountReviewRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountReviewRecordResponseBody) GetSceneDesc() *string {
	return s.SceneDesc
}

func (s *GetAccountReviewRecordResponseBody) GetScopes() []*string {
	return s.Scopes
}

func (s *GetAccountReviewRecordResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetAccountReviewRecordResponseBody) SetAddress(v string) *GetAccountReviewRecordResponseBody {
	s.Address = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetAliUid(v string) *GetAccountReviewRecordResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetApplyType(v string) *GetAccountReviewRecordResponseBody {
	s.ApplyType = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetContactName(v string) *GetAccountReviewRecordResponseBody {
	s.ContactName = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetInstanceId(v string) *GetAccountReviewRecordResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetPhone(v string) *GetAccountReviewRecordResponseBody {
	s.Phone = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetProductName(v string) *GetAccountReviewRecordResponseBody {
	s.ProductName = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetQps(v int32) *GetAccountReviewRecordResponseBody {
	s.Qps = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetRequestId(v string) *GetAccountReviewRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetSceneDesc(v string) *GetAccountReviewRecordResponseBody {
	s.SceneDesc = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetScopes(v []*string) *GetAccountReviewRecordResponseBody {
	s.Scopes = v
	return s
}

func (s *GetAccountReviewRecordResponseBody) SetServiceType(v string) *GetAccountReviewRecordResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetAccountReviewRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
