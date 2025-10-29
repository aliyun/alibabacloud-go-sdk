// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSnapshotSettingResponseBody
	GetRequestId() *string
	SetResult(v *DescribeSnapshotSettingResponseBodyResult) *DescribeSnapshotSettingResponseBody
	GetResult() *DescribeSnapshotSettingResponseBodyResult
}

type DescribeSnapshotSettingResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result *DescribeSnapshotSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeSnapshotSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotSettingResponseBody) GetResult() *DescribeSnapshotSettingResponseBodyResult {
	return s.Result
}

func (s *DescribeSnapshotSettingResponseBody) SetRequestId(v string) *DescribeSnapshotSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotSettingResponseBody) SetResult(v *DescribeSnapshotSettingResponseBodyResult) *DescribeSnapshotSettingResponseBody {
	s.Result = v
	return s
}

func (s *DescribeSnapshotSettingResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnapshotSettingResponseBodyResult struct {
	// Whether to enable automatic backup.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Automatic backup time configuration, using Quartz Cron expression.
	//
	// example:
	//
	// 0 0 01 ? 	- 	- *
	QuartzRegex *string `json:"QuartzRegex,omitempty" xml:"QuartzRegex,omitempty"`
}

func (s DescribeSnapshotSettingResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotSettingResponseBodyResult) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeSnapshotSettingResponseBodyResult) GetQuartzRegex() *string {
	return s.QuartzRegex
}

func (s *DescribeSnapshotSettingResponseBodyResult) SetEnable(v bool) *DescribeSnapshotSettingResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *DescribeSnapshotSettingResponseBodyResult) SetQuartzRegex(v string) *DescribeSnapshotSettingResponseBodyResult {
	s.QuartzRegex = &v
	return s
}

func (s *DescribeSnapshotSettingResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
