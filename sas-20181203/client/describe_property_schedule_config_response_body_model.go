// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScheduleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribePropertyScheduleConfigResponseBody
	GetConfig() *string
	SetRequestId(v string) *DescribePropertyScheduleConfigResponseBody
	GetRequestId() *string
}

type DescribePropertyScheduleConfigResponseBody struct {
	// The configuration time. Unit: hours.
	//
	// >  A value **0*	- indicates that asset fingerprint collection is disabled for this type of asset.
	//
	// example:
	//
	// 3
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyScheduleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyScheduleConfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribePropertyScheduleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyScheduleConfigResponseBody) SetConfig(v string) *DescribePropertyScheduleConfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribePropertyScheduleConfigResponseBody) SetRequestId(v string) *DescribePropertyScheduleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyScheduleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
