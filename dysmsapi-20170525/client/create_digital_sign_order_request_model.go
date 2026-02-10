// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSignOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendMessage(v string) *CreateDigitalSignOrderRequest
	GetExtendMessage() *string
	SetOrderContext(v map[string]interface{}) *CreateDigitalSignOrderRequest
	GetOrderContext() map[string]interface{}
	SetOrderType(v string) *CreateDigitalSignOrderRequest
	GetOrderType() *string
	SetOwnerId(v int64) *CreateDigitalSignOrderRequest
	GetOwnerId() *int64
	SetQualificationId(v int64) *CreateDigitalSignOrderRequest
	GetQualificationId() *int64
	SetQualificationVersion(v int64) *CreateDigitalSignOrderRequest
	GetQualificationVersion() *int64
	SetResourceOwnerAccount(v string) *CreateDigitalSignOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDigitalSignOrderRequest
	GetResourceOwnerId() *int64
	SetSignId(v int64) *CreateDigitalSignOrderRequest
	GetSignId() *int64
	SetSignIndustry(v int64) *CreateDigitalSignOrderRequest
	GetSignIndustry() *int64
	SetSignName(v string) *CreateDigitalSignOrderRequest
	GetSignName() *string
	SetSignSource(v int64) *CreateDigitalSignOrderRequest
	GetSignSource() *int64
	SetSubmitter(v string) *CreateDigitalSignOrderRequest
	GetSubmitter() *string
}

type CreateDigitalSignOrderRequest struct {
	// example:
	//
	// 示例值示例值
	ExtendMessage *string                `json:"ExtendMessage,omitempty" xml:"ExtendMessage,omitempty"`
	OrderContext  map[string]interface{} `json:"OrderContext,omitempty" xml:"OrderContext,omitempty"`
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

func (s CreateDigitalSignOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSignOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDigitalSignOrderRequest) GetExtendMessage() *string {
	return s.ExtendMessage
}

func (s *CreateDigitalSignOrderRequest) GetOrderContext() map[string]interface{} {
	return s.OrderContext
}

func (s *CreateDigitalSignOrderRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateDigitalSignOrderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDigitalSignOrderRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *CreateDigitalSignOrderRequest) GetQualificationVersion() *int64 {
	return s.QualificationVersion
}

func (s *CreateDigitalSignOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDigitalSignOrderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDigitalSignOrderRequest) GetSignId() *int64 {
	return s.SignId
}

func (s *CreateDigitalSignOrderRequest) GetSignIndustry() *int64 {
	return s.SignIndustry
}

func (s *CreateDigitalSignOrderRequest) GetSignName() *string {
	return s.SignName
}

func (s *CreateDigitalSignOrderRequest) GetSignSource() *int64 {
	return s.SignSource
}

func (s *CreateDigitalSignOrderRequest) GetSubmitter() *string {
	return s.Submitter
}

func (s *CreateDigitalSignOrderRequest) SetExtendMessage(v string) *CreateDigitalSignOrderRequest {
	s.ExtendMessage = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetOrderContext(v map[string]interface{}) *CreateDigitalSignOrderRequest {
	s.OrderContext = v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetOrderType(v string) *CreateDigitalSignOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetOwnerId(v int64) *CreateDigitalSignOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetQualificationId(v int64) *CreateDigitalSignOrderRequest {
	s.QualificationId = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetQualificationVersion(v int64) *CreateDigitalSignOrderRequest {
	s.QualificationVersion = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetResourceOwnerAccount(v string) *CreateDigitalSignOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetResourceOwnerId(v int64) *CreateDigitalSignOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSignId(v int64) *CreateDigitalSignOrderRequest {
	s.SignId = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSignIndustry(v int64) *CreateDigitalSignOrderRequest {
	s.SignIndustry = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSignName(v string) *CreateDigitalSignOrderRequest {
	s.SignName = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSignSource(v int64) *CreateDigitalSignOrderRequest {
	s.SignSource = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) SetSubmitter(v string) *CreateDigitalSignOrderRequest {
	s.Submitter = &v
	return s
}

func (s *CreateDigitalSignOrderRequest) Validate() error {
	return dara.Validate(s)
}
