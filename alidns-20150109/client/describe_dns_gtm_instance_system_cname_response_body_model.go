// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceSystemCnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDnsGtmInstanceSystemCnameResponseBody
	GetRequestId() *string
	SetSystemCname(v string) *DescribeDnsGtmInstanceSystemCnameResponseBody
	GetSystemCname() *string
}

type DescribeDnsGtmInstanceSystemCnameResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The CNAME domain name assigned by the system.
	//
	// example:
	//
	// gtminstance.com
	SystemCname *string `json:"SystemCname,omitempty" xml:"SystemCname,omitempty"`
}

func (s DescribeDnsGtmInstanceSystemCnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceSystemCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceSystemCnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmInstanceSystemCnameResponseBody) GetSystemCname() *string {
	return s.SystemCname
}

func (s *DescribeDnsGtmInstanceSystemCnameResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceSystemCnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameResponseBody) SetSystemCname(v string) *DescribeDnsGtmInstanceSystemCnameResponseBody {
	s.SystemCname = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameResponseBody) Validate() error {
	return dara.Validate(s)
}
