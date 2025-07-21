// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDeviceSeatsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeviceSeatsRequest
	GetPageSize() *int32
	SetSerialNo(v string) *DescribeDeviceSeatsRequest
	GetSerialNo() *string
	SetSerialNoList(v []*string) *DescribeDeviceSeatsRequest
	GetSerialNoList() []*string
	SetSiteId(v string) *DescribeDeviceSeatsRequest
	GetSiteId() *string
	SetTenantId(v string) *DescribeDeviceSeatsRequest
	GetTenantId() *string
}

type DescribeDeviceSeatsRequest struct {
	PageNumber   *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SerialNo     *string   `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SerialNoList []*string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty" type:"Repeated"`
	SiteId       *string   `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	TenantId     *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeDeviceSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceSeatsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeviceSeatsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeviceSeatsRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *DescribeDeviceSeatsRequest) GetSerialNoList() []*string {
	return s.SerialNoList
}

func (s *DescribeDeviceSeatsRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeDeviceSeatsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeDeviceSeatsRequest) SetPageNumber(v int32) *DescribeDeviceSeatsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetPageSize(v int32) *DescribeDeviceSeatsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetSerialNo(v string) *DescribeDeviceSeatsRequest {
	s.SerialNo = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetSerialNoList(v []*string) *DescribeDeviceSeatsRequest {
	s.SerialNoList = v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetSiteId(v string) *DescribeDeviceSeatsRequest {
	s.SiteId = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) SetTenantId(v string) *DescribeDeviceSeatsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeDeviceSeatsRequest) Validate() error {
	return dara.Validate(s)
}
