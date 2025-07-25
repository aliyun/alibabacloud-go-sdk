// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceSystemCnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeGtmInstanceSystemCnameResponseBody
	GetRequestId() *string
	SetSystemCname(v string) *DescribeGtmInstanceSystemCnameResponseBody
	GetSystemCname() *string
}

type DescribeGtmInstanceSystemCnameResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The CNAME record assigned by the system.
	//
	// example:
	//
	// gtm-cn-mp91004xxxx.gtm-a2b4.com
	SystemCname *string `json:"SystemCname,omitempty" xml:"SystemCname,omitempty"`
}

func (s DescribeGtmInstanceSystemCnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceSystemCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceSystemCnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmInstanceSystemCnameResponseBody) GetSystemCname() *string {
	return s.SystemCname
}

func (s *DescribeGtmInstanceSystemCnameResponseBody) SetRequestId(v string) *DescribeGtmInstanceSystemCnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceSystemCnameResponseBody) SetSystemCname(v string) *DescribeGtmInstanceSystemCnameResponseBody {
	s.SystemCname = &v
	return s
}

func (s *DescribeGtmInstanceSystemCnameResponseBody) Validate() error {
	return dara.Validate(s)
}
