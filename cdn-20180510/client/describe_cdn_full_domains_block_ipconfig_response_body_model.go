// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnFullDomainsBlockIPConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCdnFullDomainsBlockIPConfigResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeCdnFullDomainsBlockIPConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCdnFullDomainsBlockIPConfigResponseBody
	GetRequestId() *string
}

type DescribeCdnFullDomainsBlockIPConfigResponseBody struct {
	// The response code.
	//
	// The value of Code is not 0 in the following scenarios:
	//
	// 	- The format of the IP address is invalid.
	//
	// 	- The number of IP addresses exceeds the upper limit.
	//
	// 	- Other abnormal scenarios.
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
	// The request ID.
	//
	// example:
	//
	// 95994621-8382-464B-8762-C708E73568D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnFullDomainsBlockIPConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnFullDomainsBlockIPConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponseBody) SetCode(v int32) *DescribeCdnFullDomainsBlockIPConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponseBody) SetMessage(v string) *DescribeCdnFullDomainsBlockIPConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponseBody) SetRequestId(v string) *DescribeCdnFullDomainsBlockIPConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
