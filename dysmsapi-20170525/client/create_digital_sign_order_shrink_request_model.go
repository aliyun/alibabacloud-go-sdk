// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSignOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendMessage(v string) *CreateDigitalSignOrderShrinkRequest
	GetExtendMessage() *string
	SetOrderContextShrink(v string) *CreateDigitalSignOrderShrinkRequest
	GetOrderContextShrink() *string
	SetOrderType(v string) *CreateDigitalSignOrderShrinkRequest
	GetOrderType() *string
	SetOwnerId(v int64) *CreateDigitalSignOrderShrinkRequest
	GetOwnerId() *int64
	SetQualificationId(v int64) *CreateDigitalSignOrderShrinkRequest
	GetQualificationId() *int64
	SetQualificationVersion(v int64) *CreateDigitalSignOrderShrinkRequest
	GetQualificationVersion() *int64
	SetResourceOwnerAccount(v string) *CreateDigitalSignOrderShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDigitalSignOrderShrinkRequest
	GetResourceOwnerId() *int64
	SetSignId(v int64) *CreateDigitalSignOrderShrinkRequest
	GetSignId() *int64
	SetSignIndustry(v int64) *CreateDigitalSignOrderShrinkRequest
	GetSignIndustry() *int64
	SetSignName(v string) *CreateDigitalSignOrderShrinkRequest
	GetSignName() *string
	SetSignSource(v int64) *CreateDigitalSignOrderShrinkRequest
	GetSignSource() *int64
	SetSubmitter(v string) *CreateDigitalSignOrderShrinkRequest
	GetSubmitter() *string
}

type CreateDigitalSignOrderShrinkRequest struct {
	// example:
	//
	// 示例值示例值
	ExtendMessage      *string `json:"ExtendMessage,omitempty" xml:"ExtendMessage,omitempty"`
	OrderContextShrink *string `json:"OrderContext,omitempty" xml:"OrderContext,omitempty"`
	// example:
	//
	// CREATE_DIGITALSMS_SIGN
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 41
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// example:
	//
	// 49
	QualificationVersion *int64  `json:"QualificationVersion,omitempty" xml:"QualificationVersion,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 80
	SignId *int64 `json:"SignId,omitempty" xml:"SignId,omitempty"`
	// example:
	//
	// 0
	SignIndustry *int64 `json:"SignIndustry,omitempty" xml:"SignIndustry,omitempty"`
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// example:
	//
	// 0
	SignSource *int64 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// example:
	//
	// 110000001750080
	Submitter *string `json:"Submitter,omitempty" xml:"Submitter,omitempty"`
}

func (s CreateDigitalSignOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSignOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDigitalSignOrderShrinkRequest) GetExtendMessage() *string {
	return s.ExtendMessage
}

func (s *CreateDigitalSignOrderShrinkRequest) GetOrderContextShrink() *string {
	return s.OrderContextShrink
}

func (s *CreateDigitalSignOrderShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateDigitalSignOrderShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDigitalSignOrderShrinkRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *CreateDigitalSignOrderShrinkRequest) GetQualificationVersion() *int64 {
	return s.QualificationVersion
}

func (s *CreateDigitalSignOrderShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDigitalSignOrderShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSignId() *int64 {
	return s.SignId
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSignIndustry() *int64 {
	return s.SignIndustry
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSignName() *string {
	return s.SignName
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSignSource() *int64 {
	return s.SignSource
}

func (s *CreateDigitalSignOrderShrinkRequest) GetSubmitter() *string {
	return s.Submitter
}

func (s *CreateDigitalSignOrderShrinkRequest) SetExtendMessage(v string) *CreateDigitalSignOrderShrinkRequest {
	s.ExtendMessage = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetOrderContextShrink(v string) *CreateDigitalSignOrderShrinkRequest {
	s.OrderContextShrink = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetOrderType(v string) *CreateDigitalSignOrderShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetOwnerId(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetQualificationId(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.QualificationId = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetQualificationVersion(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.QualificationVersion = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetResourceOwnerAccount(v string) *CreateDigitalSignOrderShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetResourceOwnerId(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSignId(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.SignId = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSignIndustry(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.SignIndustry = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSignName(v string) *CreateDigitalSignOrderShrinkRequest {
	s.SignName = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSignSource(v int64) *CreateDigitalSignOrderShrinkRequest {
	s.SignSource = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) SetSubmitter(v string) *CreateDigitalSignOrderShrinkRequest {
	s.Submitter = &v
	return s
}

func (s *CreateDigitalSignOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
