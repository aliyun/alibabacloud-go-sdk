// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubDomainRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRR(v string) *DeleteSubDomainRecordsResponseBody
	GetRR() *string
	SetRequestId(v string) *DeleteSubDomainRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DeleteSubDomainRecordsResponseBody
	GetTotalCount() *string
}

type DeleteSubDomainRecordsResponseBody struct {
	// The hostname.
	//
	// example:
	//
	// www
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the DNS records to be deleted.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DeleteSubDomainRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubDomainRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubDomainRecordsResponseBody) GetRR() *string {
	return s.RR
}

func (s *DeleteSubDomainRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubDomainRecordsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DeleteSubDomainRecordsResponseBody) SetRR(v string) *DeleteSubDomainRecordsResponseBody {
	s.RR = &v
	return s
}

func (s *DeleteSubDomainRecordsResponseBody) SetRequestId(v string) *DeleteSubDomainRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubDomainRecordsResponseBody) SetTotalCount(v string) *DeleteSubDomainRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DeleteSubDomainRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
