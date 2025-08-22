// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnL2IpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDcdnL2IpsResponseBody
	GetRequestId() *string
	SetVips(v []*string) *DescribeDcdnL2IpsResponseBody
	GetVips() []*string
}

type DescribeDcdnL2IpsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The CIDR blocks of the POPs.
	Vips []*string `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Repeated"`
}

func (s DescribeDcdnL2IpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnL2IpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnL2IpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnL2IpsResponseBody) GetVips() []*string {
	return s.Vips
}

func (s *DescribeDcdnL2IpsResponseBody) SetRequestId(v string) *DescribeDcdnL2IpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnL2IpsResponseBody) SetVips(v []*string) *DescribeDcdnL2IpsResponseBody {
	s.Vips = v
	return s
}

func (s *DescribeDcdnL2IpsResponseBody) Validate() error {
	return dara.Validate(s)
}
