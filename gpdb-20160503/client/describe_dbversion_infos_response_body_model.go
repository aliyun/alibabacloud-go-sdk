// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBVersionInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDBVersionInfosResponseBody
	GetRequestId() *string
	SetVersionDetails(v *DescribeDBVersionInfosResponseBodyVersionDetails) *DescribeDBVersionInfosResponseBody
	GetVersionDetails() *DescribeDBVersionInfosResponseBodyVersionDetails
}

type DescribeDBVersionInfosResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried minor versions.
	VersionDetails *DescribeDBVersionInfosResponseBodyVersionDetails `json:"VersionDetails,omitempty" xml:"VersionDetails,omitempty" type:"Struct"`
}

func (s DescribeDBVersionInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBVersionInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBVersionInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBVersionInfosResponseBody) GetVersionDetails() *DescribeDBVersionInfosResponseBodyVersionDetails {
	return s.VersionDetails
}

func (s *DescribeDBVersionInfosResponseBody) SetRequestId(v string) *DescribeDBVersionInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBVersionInfosResponseBody) SetVersionDetails(v *DescribeDBVersionInfosResponseBodyVersionDetails) *DescribeDBVersionInfosResponseBody {
	s.VersionDetails = v
	return s
}

func (s *DescribeDBVersionInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBVersionInfosResponseBodyVersionDetails struct {
	// The queried minor version information about the instance in Serverless mode.
	//
	// example:
	//
	// "Serverless": [
	//
	//                 {
	//
	//                     "engineVersion": "6.0",
	//
	//                     "versionInfos": [
	//
	//                         {
	//
	//                             "kernelVersion": "v2.0.0.5",
	//
	//                             "releaseDate": "2023-05-28T07:48Z",
	//
	//                             "expirationDate": "2026-05-28T07:48Z"
	//
	//                         },
	//
	//                         {
	//
	//                             "kernelVersion": "v2.0.0.1",
	//
	//                             "releaseDate": "2023-03-27T12:44Z",
	//
	//                             "expirationDate": "2026-03-27T12:44Z"
	//
	//                         },
	//
	//                         {
	//
	//                             "kernelVersion": "v1.0.5.1",
	//
	//                             "releaseDate": "2023-02-22T11:39Z",
	//
	//                             "expirationDate": "2026-02-22T11:39Z"
	//
	//                         }
	//
	//                     ]
	//
	//                 }
	//
	// ]
	Serverless interface{} `json:"Serverless,omitempty" xml:"Serverless,omitempty"`
	// The queried minor version information about the instance in elastic storage mode.
	//
	// example:
	//
	// "StorageElasic": [
	//
	//                 {
	//
	//                     "engineVersion": "6.0",
	//
	//                     "versionInfos": [
	//
	//                         {
	//
	//                             "kernelVersion": "v6.3.11.2",
	//
	//                             "releaseDate": "2023-08-17T09:14Z",
	//
	//                             "expirationDate": "2026-08-17T09:14Z"
	//
	//                         }
	//
	//           },
	//
	//                 {
	//
	//                     "engineVersion": "7.0",
	//
	//                     "versionInfos": [
	//
	//                         {
	//
	//                             "kernelVersion": "v7.0.2.0",
	//
	//                             "releaseDate": "2023-08-09T06:47Z",
	//
	//                             "expirationDate": "2026-08-09T06:47Z"
	//
	//                         },
	//
	//                         {
	//
	//                             "kernelVersion": "v7.0.1.8",
	//
	//                             "releaseDate": "2023-05-25T06:56Z",
	//
	//                             "expirationDate": "2026-05-25T06:56Z"
	//
	//                         }
	//
	//                     ]
	//
	//                 }
	//
	// ]
	StorageElastic interface{} `json:"StorageElastic,omitempty" xml:"StorageElastic,omitempty"`
}

func (s DescribeDBVersionInfosResponseBodyVersionDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBVersionInfosResponseBodyVersionDetails) GoString() string {
	return s.String()
}

func (s *DescribeDBVersionInfosResponseBodyVersionDetails) GetServerless() interface{} {
	return s.Serverless
}

func (s *DescribeDBVersionInfosResponseBodyVersionDetails) GetStorageElastic() interface{} {
	return s.StorageElastic
}

func (s *DescribeDBVersionInfosResponseBodyVersionDetails) SetServerless(v interface{}) *DescribeDBVersionInfosResponseBodyVersionDetails {
	s.Serverless = v
	return s
}

func (s *DescribeDBVersionInfosResponseBodyVersionDetails) SetStorageElastic(v interface{}) *DescribeDBVersionInfosResponseBodyVersionDetails {
	s.StorageElastic = v
	return s
}

func (s *DescribeDBVersionInfosResponseBodyVersionDetails) Validate() error {
	return dara.Validate(s)
}
