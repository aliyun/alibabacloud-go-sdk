// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHASwitchConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHAConfig(v string) *DescribeHASwitchConfigResponseBody
	GetHAConfig() *string
	SetManualHATime(v string) *DescribeHASwitchConfigResponseBody
	GetManualHATime() *string
	SetRequestId(v string) *DescribeHASwitchConfigResponseBody
	GetRequestId() *string
}

type DescribeHASwitchConfigResponseBody struct {
	HAConfig     *string `json:"HAConfig,omitempty" xml:"HAConfig,omitempty"`
	ManualHATime *string `json:"ManualHATime,omitempty" xml:"ManualHATime,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHASwitchConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHASwitchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHASwitchConfigResponseBody) GetHAConfig() *string {
	return s.HAConfig
}

func (s *DescribeHASwitchConfigResponseBody) GetManualHATime() *string {
	return s.ManualHATime
}

func (s *DescribeHASwitchConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHASwitchConfigResponseBody) SetHAConfig(v string) *DescribeHASwitchConfigResponseBody {
	s.HAConfig = &v
	return s
}

func (s *DescribeHASwitchConfigResponseBody) SetManualHATime(v string) *DescribeHASwitchConfigResponseBody {
	s.ManualHATime = &v
	return s
}

func (s *DescribeHASwitchConfigResponseBody) SetRequestId(v string) *DescribeHASwitchConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHASwitchConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
