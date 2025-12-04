// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProducts(v []*GetAccessKeyLastUsedProductsResponseBodyProducts) *GetAccessKeyLastUsedProductsResponseBody
	GetProducts() []*GetAccessKeyLastUsedProductsResponseBodyProducts
	SetRequestId(v string) *GetAccessKeyLastUsedProductsResponseBody
	GetRequestId() *string
}

type GetAccessKeyLastUsedProductsResponseBody struct {
	// The list of returned Alibaba Cloud services.
	//
	// This parameter is required.
	Products []*GetAccessKeyLastUsedProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 145318BE-DEE1-4C57-AA7C-5BE7D34A6AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedProductsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedProductsResponseBody) GetProducts() []*GetAccessKeyLastUsedProductsResponseBodyProducts {
	return s.Products
}

func (s *GetAccessKeyLastUsedProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyLastUsedProductsResponseBody) SetProducts(v []*GetAccessKeyLastUsedProductsResponseBodyProducts) *GetAccessKeyLastUsedProductsResponseBody {
	s.Products = v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBody) Validate() error {
	if s.Products != nil {
		for _, item := range s.Products {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAccessKeyLastUsedProductsResponseBodyProducts struct {
	// The event details.
	//
	// example:
	//
	// {
	//
	//   "eventId": "239EB588-CD24-522E-B0B5-174A1A58****",
	//
	//   "eventVersion": 1,
	//
	//   "eventSource": "ecs.cn-hangzhou.aliyuncs.com",
	//
	//   "sourceIpAddress": "``10.10.**.**``",
	//
	//   "eventType": "ApiCall",
	//
	//   "userIdentity": {
	//
	//     "accountId": "104758519118****",
	//
	//     "principalId": "24549429003625****",
	//
	//     "type": "ram-user",
	//
	//     "userName": "alice"
	//
	//   },
	//
	//   "serviceName": "Ecs",
	//
	//   "apiVersion": "2016-01-20",
	//
	//   "requestId": "239EB588-CD24-522E-B0B5-174A1A588BE0",
	//
	//   "eventTime": "2021-08-05T09:21:32Z",
	//
	//   "isGlobal": false,
	//
	//   "acsRegion": "cn-hangzhou",
	//
	//   "eventName": "DescribeInstances"
	//
	// }
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The Alibaba Cloud service.
	//
	// example:
	//
	// Ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The Chinese name of the Alibaba Cloud service.
	//
	// example:
	//
	// Elastic Compute Service (ECS)
	ServiceNameCn *string `json:"ServiceNameCn,omitempty" xml:"ServiceNameCn,omitempty"`
	// The English name of the Alibaba Cloud service.
	//
	// example:
	//
	// Elastic Compute Service
	ServiceNameEn *string `json:"ServiceNameEn,omitempty" xml:"ServiceNameEn,omitempty"`
	// The event source.
	//
	// Valid values:
	//
	// 	- Internal
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     other events
	//
	//     <!-- -->
	//
	// 	- ManagementEvent
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     management events
	//
	//     <!-- -->
	//
	// 	- DataEvent
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     data events
	//
	//     <!-- -->
	//
	// example:
	//
	// ManagementEvent
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Unit: millisecond.
	//
	// example:
	//
	// 1657247532000
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
}

func (s GetAccessKeyLastUsedProductsResponseBodyProducts) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) GetDetail() *string {
	return s.Detail
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) GetServiceNameCn() *string {
	return s.ServiceNameCn
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) GetServiceNameEn() *string {
	return s.ServiceNameEn
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) GetSource() *string {
	return s.Source
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) GetUsedTimestamp() *int64 {
	return s.UsedTimestamp
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetDetail(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetServiceName(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.ServiceName = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetServiceNameCn(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.ServiceNameCn = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetServiceNameEn(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.ServiceNameEn = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetSource(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.UsedTimestamp = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) Validate() error {
	return dara.Validate(s)
}
