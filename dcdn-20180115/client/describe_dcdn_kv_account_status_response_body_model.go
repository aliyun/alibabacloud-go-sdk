// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnKvAccountStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDcdnKvAccountStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDcdnKvAccountStatusResponseBody
	GetStatus() *string
}

type DescribeDcdnKvAccountStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-82C0-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the account.
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnKvAccountStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnKvAccountStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnKvAccountStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnKvAccountStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnKvAccountStatusResponseBody) SetRequestId(v string) *DescribeDcdnKvAccountStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnKvAccountStatusResponseBody) SetStatus(v string) *DescribeDcdnKvAccountStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDcdnKvAccountStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
