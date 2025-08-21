// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSnapshotSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSnapshotSettingResponseBody
	GetRequestId() *string
	SetResult(v *UpdateSnapshotSettingResponseBodyResult) *UpdateSnapshotSettingResponseBody
	GetResult() *UpdateSnapshotSettingResponseBodyResult
}

type UpdateSnapshotSettingResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result *UpdateSnapshotSettingResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateSnapshotSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSnapshotSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSnapshotSettingResponseBody) GetResult() *UpdateSnapshotSettingResponseBodyResult {
	return s.Result
}

func (s *UpdateSnapshotSettingResponseBody) SetRequestId(v string) *UpdateSnapshotSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSnapshotSettingResponseBody) SetResult(v *UpdateSnapshotSettingResponseBodyResult) *UpdateSnapshotSettingResponseBody {
	s.Result = v
	return s
}

func (s *UpdateSnapshotSettingResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateSnapshotSettingResponseBodyResult struct {
	// Specifies whether to enable automatic backup.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The start time of automatic backup.
	//
	// example:
	//
	// 0 0 01 ? 	- 	- *
	QuartzRegex *string `json:"quartzRegex,omitempty" xml:"quartzRegex,omitempty"`
}

func (s UpdateSnapshotSettingResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateSnapshotSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingResponseBodyResult) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateSnapshotSettingResponseBodyResult) GetQuartzRegex() *string {
	return s.QuartzRegex
}

func (s *UpdateSnapshotSettingResponseBodyResult) SetEnable(v bool) *UpdateSnapshotSettingResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *UpdateSnapshotSettingResponseBodyResult) SetQuartzRegex(v string) *UpdateSnapshotSettingResponseBodyResult {
	s.QuartzRegex = &v
	return s
}

func (s *UpdateSnapshotSettingResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
