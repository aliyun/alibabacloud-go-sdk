// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnFullDomainsBlockIPConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeDcdnFullDomainsBlockIPConfigResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeDcdnFullDomainsBlockIPConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDcdnFullDomainsBlockIPConfigResponseBody
	GetRequestId() *string
}

type DescribeDcdnFullDomainsBlockIPConfigResponseBody struct {
	// The response code.
	//
	// The value of Code is not 0 in the following scenarios:
	//
	// 	- The format of the IP address is invalid.
	//
	// 	- The number of IP addresses exceeds the limit.
	//
	// 	- Other abnormal scenarios
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned results. If the operation is successful, URLs of OSS objects are returned. If the operation fails, an error message is returned.
	//
	// example:
	//
	// http://xxxx-api.oss-cn-hangzhou.aliyuncs.com/blocklist%2Fxxxxxxxxxxxx.txt?Expires=1682663947&OSSAccessKeyId=xxxxxxxxxx&Signature=xxxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C58632F-BA12-1A1E-986D-09285752B42C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnFullDomainsBlockIPConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnFullDomainsBlockIPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponseBody) SetCode(v int32) *DescribeDcdnFullDomainsBlockIPConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponseBody) SetMessage(v string) *DescribeDcdnFullDomainsBlockIPConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponseBody) SetRequestId(v string) *DescribeDcdnFullDomainsBlockIPConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
