// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConcernNecessityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConcernNecessity(v []*string) *DescribeConcernNecessityResponseBody
	GetConcernNecessity() []*string
	SetRequestId(v string) *DescribeConcernNecessityResponseBody
	GetRequestId() *string
}

type DescribeConcernNecessityResponseBody struct {
	// The priorities to fix the vulnerabilities. Valid values:
	//
	// 	- asap: high
	//
	// 	- later: medium
	//
	// 	- nntf: low
	ConcernNecessity []*string `json:"ConcernNecessity,omitempty" xml:"ConcernNecessity,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// ECC6B3E3-D496-512D-B46D-E6996A6B63EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConcernNecessityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConcernNecessityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConcernNecessityResponseBody) GetConcernNecessity() []*string {
	return s.ConcernNecessity
}

func (s *DescribeConcernNecessityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConcernNecessityResponseBody) SetConcernNecessity(v []*string) *DescribeConcernNecessityResponseBody {
	s.ConcernNecessity = v
	return s
}

func (s *DescribeConcernNecessityResponseBody) SetRequestId(v string) *DescribeConcernNecessityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConcernNecessityResponseBody) Validate() error {
	return dara.Validate(s)
}
