// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBasicStartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *DescribeBasicStartRequest
	GetAppKey() *string
	SetEndDs(v string) *DescribeBasicStartRequest
	GetEndDs() *string
	SetRegId(v string) *DescribeBasicStartRequest
	GetRegId() *string
	SetService(v string) *DescribeBasicStartRequest
	GetService() *string
	SetStartDs(v string) *DescribeBasicStartRequest
	GetStartDs() *string
}

type DescribeBasicStartRequest struct {
	// Application appkey.
	//
	// example:
	//
	// ***
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// End time, accurate to milliseconds (ms).
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
	// Service to call
	//
	// example:
	//
	// service_tender_cee
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// Start time, accurate to milliseconds (ms).
	//
	// example:
	//
	// 20250310
	StartDs *string `json:"startDs,omitempty" xml:"startDs,omitempty"`
}

func (s DescribeBasicStartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBasicStartRequest) GoString() string {
	return s.String()
}

func (s *DescribeBasicStartRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeBasicStartRequest) GetEndDs() *string {
	return s.EndDs
}

func (s *DescribeBasicStartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeBasicStartRequest) GetService() *string {
	return s.Service
}

func (s *DescribeBasicStartRequest) GetStartDs() *string {
	return s.StartDs
}

func (s *DescribeBasicStartRequest) SetAppKey(v string) *DescribeBasicStartRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeBasicStartRequest) SetEndDs(v string) *DescribeBasicStartRequest {
	s.EndDs = &v
	return s
}

func (s *DescribeBasicStartRequest) SetRegId(v string) *DescribeBasicStartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeBasicStartRequest) SetService(v string) *DescribeBasicStartRequest {
	s.Service = &v
	return s
}

func (s *DescribeBasicStartRequest) SetStartDs(v string) *DescribeBasicStartRequest {
	s.StartDs = &v
	return s
}

func (s *DescribeBasicStartRequest) Validate() error {
	return dara.Validate(s)
}
