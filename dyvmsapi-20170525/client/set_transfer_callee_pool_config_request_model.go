// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTransferCalleePoolConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledRouteMode(v string) *SetTransferCalleePoolConfigRequest
	GetCalledRouteMode() *string
	SetDetails(v []*SetTransferCalleePoolConfigRequestDetails) *SetTransferCalleePoolConfigRequest
	GetDetails() []*SetTransferCalleePoolConfigRequestDetails
	SetOwnerId(v int64) *SetTransferCalleePoolConfigRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *SetTransferCalleePoolConfigRequest
	GetPhoneNumber() *string
	SetQualificationId(v string) *SetTransferCalleePoolConfigRequest
	GetQualificationId() *string
	SetResourceOwnerAccount(v string) *SetTransferCalleePoolConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetTransferCalleePoolConfigRequest
	GetResourceOwnerId() *int64
}

type SetTransferCalleePoolConfigRequest struct {
	// The call mode. Valid values:
	//
	// 	- **roundRobin**
	//
	// 	- **random**
	//
	// This parameter is required.
	//
	// example:
	//
	// roundRobin
	CalledRouteMode *string `json:"CalledRouteMode,omitempty" xml:"CalledRouteMode,omitempty"`
	// The information about the phone numbers for transferring the call.
	//
	// This parameter is required.
	Details []*SetTransferCalleePoolConfigRequestDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	OwnerId *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number used for transferring the call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The qualification ID. You can call the [GetHotlineQualificationByOrder](https://help.aliyun.com/document_detail/393548.html) operation to obtain the qualification ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 190***
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetTransferCalleePoolConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetTransferCalleePoolConfigRequest) GoString() string {
	return s.String()
}

func (s *SetTransferCalleePoolConfigRequest) GetCalledRouteMode() *string {
	return s.CalledRouteMode
}

func (s *SetTransferCalleePoolConfigRequest) GetDetails() []*SetTransferCalleePoolConfigRequestDetails {
	return s.Details
}

func (s *SetTransferCalleePoolConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetTransferCalleePoolConfigRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *SetTransferCalleePoolConfigRequest) GetQualificationId() *string {
	return s.QualificationId
}

func (s *SetTransferCalleePoolConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetTransferCalleePoolConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetTransferCalleePoolConfigRequest) SetCalledRouteMode(v string) *SetTransferCalleePoolConfigRequest {
	s.CalledRouteMode = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetDetails(v []*SetTransferCalleePoolConfigRequestDetails) *SetTransferCalleePoolConfigRequest {
	s.Details = v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetOwnerId(v int64) *SetTransferCalleePoolConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetPhoneNumber(v string) *SetTransferCalleePoolConfigRequest {
	s.PhoneNumber = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetQualificationId(v string) *SetTransferCalleePoolConfigRequest {
	s.QualificationId = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetResourceOwnerAccount(v string) *SetTransferCalleePoolConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetResourceOwnerId(v int64) *SetTransferCalleePoolConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) Validate() error {
	return dara.Validate(s)
}

type SetTransferCalleePoolConfigRequestDetails struct {
	// The called number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 151****0000
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 150****0000
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s SetTransferCalleePoolConfigRequestDetails) String() string {
	return dara.Prettify(s)
}

func (s SetTransferCalleePoolConfigRequestDetails) GoString() string {
	return s.String()
}

func (s *SetTransferCalleePoolConfigRequestDetails) GetCalled() *string {
	return s.Called
}

func (s *SetTransferCalleePoolConfigRequestDetails) GetCaller() *string {
	return s.Caller
}

func (s *SetTransferCalleePoolConfigRequestDetails) SetCalled(v string) *SetTransferCalleePoolConfigRequestDetails {
	s.Called = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequestDetails) SetCaller(v string) *SetTransferCalleePoolConfigRequestDetails {
	s.Caller = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequestDetails) Validate() error {
	return dara.Validate(s)
}
