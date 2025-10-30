// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretNoRemainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *QuerySecretNoRemainRequest
	GetCity() *string
	SetOwnerId(v int64) *QuerySecretNoRemainRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySecretNoRemainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySecretNoRemainRequest
	GetResourceOwnerId() *int64
	SetSecretNo(v string) *QuerySecretNoRemainRequest
	GetSecretNo() *string
	SetSpecId(v int64) *QuerySecretNoRemainRequest
	GetSpecId() *int64
}

type QuerySecretNoRemainRequest struct {
	// The home location of the phone number.
	//
	// 	- If **SpecId*	- is set to 1 or 2, you can specify the **City*	- parameter to query the quantity of available phone numbers.
	//
	// 1.  You can enter a single city name to perform a query.
	//
	// 2.  You can enter National to query the quantity of remaining phone numbers available in the Chinese mainland for online purchase.
	//
	// 3.  You can enter National List to query the cities with available phone numbers and the quantities of remaining phone numbers in the Chinese mainland. Cities without available phone numbers will not be returned.
	//
	// 	- If **SpecId*	- is set to 3, home locations are not distinguished for phone numbers that start with 95 and only the quantity of all the remaining phone numbers that start with 95 and are available for online purchase can be queried. If SpecId is set to 3, **City*	- must be set to **Nationwide**.
	//
	// >  Home locations can be set to only locations in the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// hangzhou
	City                 *string `json:"City,omitempty" xml:"City,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The prefix of the phone number. When you call the QuerySecretNoRemain operation with **SecretNo*	- specified, the quantity of remaining phone numbers with the specified prefix that are available for online purchase is queried.
	//
	// Up to 18 digits of a phone number prefix can be specified.
	//
	// example:
	//
	// 130
	SecretNo *string `json:"SecretNo,omitempty" xml:"SecretNo,omitempty"`
	// The type of the phone number. Valid values:
	//
	// 	- **1**: a phone number assigned by a virtual network operator, that is, a phone number that belongs to the 170 or 171 number segment.
	//
	// 	- **2**: a phone number provided by a carrier.
	//
	// 	- **3**: a phone number that starts with 95.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SpecId *int64 `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s QuerySecretNoRemainRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretNoRemainRequest) GoString() string {
	return s.String()
}

func (s *QuerySecretNoRemainRequest) GetCity() *string {
	return s.City
}

func (s *QuerySecretNoRemainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySecretNoRemainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySecretNoRemainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySecretNoRemainRequest) GetSecretNo() *string {
	return s.SecretNo
}

func (s *QuerySecretNoRemainRequest) GetSpecId() *int64 {
	return s.SpecId
}

func (s *QuerySecretNoRemainRequest) SetCity(v string) *QuerySecretNoRemainRequest {
	s.City = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetOwnerId(v int64) *QuerySecretNoRemainRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetResourceOwnerAccount(v string) *QuerySecretNoRemainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetResourceOwnerId(v int64) *QuerySecretNoRemainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetSecretNo(v string) *QuerySecretNoRemainRequest {
	s.SecretNo = &v
	return s
}

func (s *QuerySecretNoRemainRequest) SetSpecId(v int64) *QuerySecretNoRemainRequest {
	s.SpecId = &v
	return s
}

func (s *QuerySecretNoRemainRequest) Validate() error {
	return dara.Validate(s)
}
