// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLWebSocketDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeSQLWebSocketDomainResponseBody
	GetCode() *int64
	SetDomain(v string) *DescribeSQLWebSocketDomainResponseBody
	GetDomain() *string
	SetMessage(v string) *DescribeSQLWebSocketDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSQLWebSocketDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSQLWebSocketDomainResponseBody
	GetSuccess() *bool
}

type DescribeSQLWebSocketDomainResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The domain name.
	//
	// example:
	//
	// adb-ws-beijing.console.aliyun.com/query
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E03F0806-A67B-5B24-8562-9589F20DEEB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSQLWebSocketDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLWebSocketDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLWebSocketDomainResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeSQLWebSocketDomainResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeSQLWebSocketDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSQLWebSocketDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLWebSocketDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSQLWebSocketDomainResponseBody) SetCode(v int64) *DescribeSQLWebSocketDomainResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSQLWebSocketDomainResponseBody) SetDomain(v string) *DescribeSQLWebSocketDomainResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeSQLWebSocketDomainResponseBody) SetMessage(v string) *DescribeSQLWebSocketDomainResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSQLWebSocketDomainResponseBody) SetRequestId(v string) *DescribeSQLWebSocketDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLWebSocketDomainResponseBody) SetSuccess(v bool) *DescribeSQLWebSocketDomainResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSQLWebSocketDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
