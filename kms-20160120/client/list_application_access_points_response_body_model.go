// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccessPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccessPoints(v *ListApplicationAccessPointsResponseBodyApplicationAccessPoints) *ListApplicationAccessPointsResponseBody
	GetApplicationAccessPoints() *ListApplicationAccessPointsResponseBodyApplicationAccessPoints
	SetPageNumber(v int32) *ListApplicationAccessPointsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationAccessPointsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApplicationAccessPointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationAccessPointsResponseBody
	GetTotalCount() *int32
}

type ListApplicationAccessPointsResponseBody struct {
	// A list of AAPs.
	ApplicationAccessPoints *ListApplicationAccessPointsResponseBodyApplicationAccessPoints `json:"ApplicationAccessPoints,omitempty" xml:"ApplicationAccessPoints,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationAccessPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsResponseBody) GetApplicationAccessPoints() *ListApplicationAccessPointsResponseBodyApplicationAccessPoints {
	return s.ApplicationAccessPoints
}

func (s *ListApplicationAccessPointsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationAccessPointsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationAccessPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationAccessPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationAccessPointsResponseBody) SetApplicationAccessPoints(v *ListApplicationAccessPointsResponseBodyApplicationAccessPoints) *ListApplicationAccessPointsResponseBody {
	s.ApplicationAccessPoints = v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) SetPageNumber(v int32) *ListApplicationAccessPointsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) SetPageSize(v int32) *ListApplicationAccessPointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) SetRequestId(v string) *ListApplicationAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) SetTotalCount(v int32) *ListApplicationAccessPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationAccessPointsResponseBodyApplicationAccessPoints struct {
	ApplicationAccessPoint []*ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint `json:"ApplicationAccessPoint,omitempty" xml:"ApplicationAccessPoint,omitempty" type:"Repeated"`
}

func (s ListApplicationAccessPointsResponseBodyApplicationAccessPoints) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccessPointsResponseBodyApplicationAccessPoints) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPoints) GetApplicationAccessPoint() []*ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint {
	return s.ApplicationAccessPoint
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPoints) SetApplicationAccessPoint(v []*ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) *ListApplicationAccessPointsResponseBodyApplicationAccessPoints {
	s.ApplicationAccessPoint = v
	return s
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPoints) Validate() error {
	return dara.Validate(s)
}

type ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint struct {
	// The authentication method.
	//
	// example:
	//
	// ClientKey
	AuthenticationMethod *string `json:"AuthenticationMethod,omitempty" xml:"AuthenticationMethod,omitempty"`
	// The name of the AAP.
	//
	// example:
	//
	// aap_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) GetAuthenticationMethod() *string {
	return s.AuthenticationMethod
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) GetName() *string {
	return s.Name
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) SetAuthenticationMethod(v string) *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint {
	s.AuthenticationMethod = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) SetName(v string) *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint {
	s.Name = &v
	return s
}

func (s *ListApplicationAccessPointsResponseBodyApplicationAccessPointsApplicationAccessPoint) Validate() error {
	return dara.Validate(s)
}
