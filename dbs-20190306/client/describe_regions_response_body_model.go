// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeRegionsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeRegionsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeRegionsResponseBody
	GetHttpStatusCode() *int32
	SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() *DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRegionsResponseBody
	GetSuccess() *bool
}

type DescribeRegionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// This backupPlan can\\"t support this action
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The regions that DBS supports.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EB4DFD5E-3618-498D-BE35-4DBEA0072122
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeRegionsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeRegionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeRegionsResponseBody) GetRegions() *DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRegionsResponseBody) SetErrCode(v string) *DescribeRegionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrMessage(v string) *DescribeRegionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetHttpStatusCode(v int32) *DescribeRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegions struct {
	RegionCode []*string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetRegionCode() []*string {
	return s.RegionCode
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionCode(v []*string) *DescribeRegionsResponseBodyRegions {
	s.RegionCode = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
