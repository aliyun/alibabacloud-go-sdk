// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSafOrderResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeSafOrderResponseBodyResultObject) *DescribeSafOrderResponseBody
	GetResultObject() *DescribeSafOrderResponseBodyResultObject
}

type DescribeSafOrderResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object.
	ResultObject *DescribeSafOrderResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeSafOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSafOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSafOrderResponseBody) GetResultObject() *DescribeSafOrderResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSafOrderResponseBody) SetRequestId(v string) *DescribeSafOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSafOrderResponseBody) SetResultObject(v *DescribeSafOrderResponseBodyResultObject) *DescribeSafOrderResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSafOrderResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSafOrderResponseBodyResultObject struct {
	// Expiration date (timestamp).
	//
	// example:
	//
	// 1755076800000
	ExpirationDate *int64 `json:"expirationDate,omitempty" xml:"expirationDate,omitempty"`
}

func (s DescribeSafOrderResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafOrderResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSafOrderResponseBodyResultObject) GetExpirationDate() *int64 {
	return s.ExpirationDate
}

func (s *DescribeSafOrderResponseBodyResultObject) SetExpirationDate(v int64) *DescribeSafOrderResponseBodyResultObject {
	s.ExpirationDate = &v
	return s
}

func (s *DescribeSafOrderResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
