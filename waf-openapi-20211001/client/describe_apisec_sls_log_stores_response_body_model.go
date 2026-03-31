// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSlsLogStoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogStores(v []*string) *DescribeApisecSlsLogStoresResponseBody
	GetLogStores() []*string
	SetRequestId(v string) *DescribeApisecSlsLogStoresResponseBody
	GetRequestId() *string
}

type DescribeApisecSlsLogStoresResponseBody struct {
	// The names of the Logstores in Simple Log Service.
	LogStores []*string `json:"LogStores,omitempty" xml:"LogStores,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApisecSlsLogStoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSlsLogStoresResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecSlsLogStoresResponseBody) GetLogStores() []*string {
	return s.LogStores
}

func (s *DescribeApisecSlsLogStoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecSlsLogStoresResponseBody) SetLogStores(v []*string) *DescribeApisecSlsLogStoresResponseBody {
	s.LogStores = v
	return s
}

func (s *DescribeApisecSlsLogStoresResponseBody) SetRequestId(v string) *DescribeApisecSlsLogStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecSlsLogStoresResponseBody) Validate() error {
	return dara.Validate(s)
}
