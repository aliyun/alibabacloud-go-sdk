// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafDeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSafDeOrderResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeSafDeOrderResponseBodyResultObject) *DescribeSafDeOrderResponseBody
	GetResultObject() *DescribeSafDeOrderResponseBodyResultObject
}

type DescribeSafDeOrderResponseBody struct {
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	ResultObject *DescribeSafDeOrderResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeSafDeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafDeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSafDeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSafDeOrderResponseBody) GetResultObject() *DescribeSafDeOrderResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSafDeOrderResponseBody) SetRequestId(v string) *DescribeSafDeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSafDeOrderResponseBody) SetResultObject(v *DescribeSafDeOrderResponseBodyResultObject) *DescribeSafDeOrderResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSafDeOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSafDeOrderResponseBodyResultObject struct {
	// Expiration time
	//
	// example:
	//
	// 1728008155799
	ExpirationDate *int64 `json:"expirationDate,omitempty" xml:"expirationDate,omitempty"`
	// Based on the product type subscribed by the customer, the console permissions are divided into three categories:
	//
	//      1. New Customer: Has not purchased/subscribed to any service.
	//
	//      2. Old Customer (Subscription): Customers who have purchased the SAF product.
	//
	//      3. Pay-As-You-Go: Customers who have purchased the SAF_BAG product or activated SAF_POS.
	//
	// example:
	//
	// 2
	OpenUserType *int32 `json:"openUserType,omitempty" xml:"openUserType,omitempty"`
}

func (s DescribeSafDeOrderResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafDeOrderResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSafDeOrderResponseBodyResultObject) GetExpirationDate() *int64 {
	return s.ExpirationDate
}

func (s *DescribeSafDeOrderResponseBodyResultObject) GetOpenUserType() *int32 {
	return s.OpenUserType
}

func (s *DescribeSafDeOrderResponseBodyResultObject) SetExpirationDate(v int64) *DescribeSafDeOrderResponseBodyResultObject {
	s.ExpirationDate = &v
	return s
}

func (s *DescribeSafDeOrderResponseBodyResultObject) SetOpenUserType(v int32) *DescribeSafDeOrderResponseBodyResultObject {
	s.OpenUserType = &v
	return s
}

func (s *DescribeSafDeOrderResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
