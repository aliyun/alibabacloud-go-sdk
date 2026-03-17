// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBindableSmartAccessGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeBindableSmartAccessGatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBindableSmartAccessGatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBindableSmartAccessGatewaysResponseBody
	GetRequestId() *string
	SetSmartAccessGateways(v *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways) *DescribeBindableSmartAccessGatewaysResponseBody
	GetSmartAccessGateways() *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways
	SetTotalCount(v int32) *DescribeBindableSmartAccessGatewaysResponseBody
	GetTotalCount() *int32
}

type DescribeBindableSmartAccessGatewaysResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9731C2F5-B9A4-42FD-AFD2-361A403E6E85
	RequestId           *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmartAccessGateways *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways `json:"SmartAccessGateways,omitempty" xml:"SmartAccessGateways,omitempty" type:"Struct"`
	// The total number of SAG instances.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBindableSmartAccessGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBindableSmartAccessGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) GetSmartAccessGateways() *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways {
	return s.SmartAccessGateways
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) SetPageNumber(v int32) *DescribeBindableSmartAccessGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) SetPageSize(v int32) *DescribeBindableSmartAccessGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) SetRequestId(v string) *DescribeBindableSmartAccessGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) SetSmartAccessGateways(v *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways) *DescribeBindableSmartAccessGatewaysResponseBody {
	s.SmartAccessGateways = v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) SetTotalCount(v int32) *DescribeBindableSmartAccessGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBody) Validate() error {
	if s.SmartAccessGateways != nil {
		if err := s.SmartAccessGateways.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways struct {
	SmartAccessGateway []*DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway `json:"SmartAccessGateway,omitempty" xml:"SmartAccessGateway,omitempty" type:"Repeated"`
}

func (s DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways) GoString() string {
	return s.String()
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways) GetSmartAccessGateway() []*DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	return s.SmartAccessGateway
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways) SetSmartAccessGateway(v []*DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways {
	s.SmartAccessGateway = v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGateways) Validate() error {
	if s.SmartAccessGateway != nil {
		for _, item := range s.SmartAccessGateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SmartAGId  *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	SmartAGUid *int64  `json:"SmartAGUid,omitempty" xml:"SmartAGUid,omitempty"`
}

func (s DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GoString() string {
	return s.String()
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetName() *string {
	return s.Name
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetSmartAGUid() *int64 {
	return s.SmartAGUid
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetName(v string) *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.Name = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetSmartAGId(v string) *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.SmartAGId = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetSmartAGUid(v int64) *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.SmartAGUid = &v
	return s
}

func (s *DescribeBindableSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) Validate() error {
	return dara.Validate(s)
}
