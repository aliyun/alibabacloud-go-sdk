// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDetailStartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *DescribeDetailStartRequest
	GetAppKey() *string
	SetEndDs(v string) *DescribeDetailStartRequest
	GetEndDs() *string
	SetRegId(v string) *DescribeDetailStartRequest
	GetRegId() *string
	SetService(v string) *DescribeDetailStartRequest
	GetService() *string
	SetStartDs(v string) *DescribeDetailStartRequest
	GetStartDs() *string
}

type DescribeDetailStartRequest struct {
	// Application appkey.
	//
	// example:
	//
	// ***
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// End time
	//
	//
	//
	// Format yyyymmdd
	//
	// example:
	//
	// 20250320
	EndDs *string `json:"endDs,omitempty" xml:"endDs,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Called service
	//
	// example:
	//
	// service_tender_cee
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// Start time
	//
	//
	//
	//  Format yyyymmdd
	//
	// example:
	//
	// 20250310
	StartDs *string `json:"startDs,omitempty" xml:"startDs,omitempty"`
}

func (s DescribeDetailStartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetailStartRequest) GoString() string {
	return s.String()
}

func (s *DescribeDetailStartRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeDetailStartRequest) GetEndDs() *string {
	return s.EndDs
}

func (s *DescribeDetailStartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeDetailStartRequest) GetService() *string {
	return s.Service
}

func (s *DescribeDetailStartRequest) GetStartDs() *string {
	return s.StartDs
}

func (s *DescribeDetailStartRequest) SetAppKey(v string) *DescribeDetailStartRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeDetailStartRequest) SetEndDs(v string) *DescribeDetailStartRequest {
	s.EndDs = &v
	return s
}

func (s *DescribeDetailStartRequest) SetRegId(v string) *DescribeDetailStartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeDetailStartRequest) SetService(v string) *DescribeDetailStartRequest {
	s.Service = &v
	return s
}

func (s *DescribeDetailStartRequest) SetStartDs(v string) *DescribeDetailStartRequest {
	s.StartDs = &v
	return s
}

func (s *DescribeDetailStartRequest) Validate() error {
	return dara.Validate(s)
}
