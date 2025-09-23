// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackMaxQpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQps(v string) *DescribeDomainAttackMaxQpsResponseBody
	GetQps() *string
	SetRequestId(v string) *DescribeDomainAttackMaxQpsResponseBody
	GetRequestId() *string
}

type DescribeDomainAttackMaxQpsResponseBody struct {
	// example:
	//
	// 613
	Qps *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// example:
	//
	// 62F9BD81-8BCA-5B23-A3CB-3FB7CEB7A4CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainAttackMaxQpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackMaxQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackMaxQpsResponseBody) GetQps() *string {
	return s.Qps
}

func (s *DescribeDomainAttackMaxQpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainAttackMaxQpsResponseBody) SetQps(v string) *DescribeDomainAttackMaxQpsResponseBody {
	s.Qps = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsResponseBody) SetRequestId(v string) *DescribeDomainAttackMaxQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsResponseBody) Validate() error {
	return dara.Validate(s)
}
