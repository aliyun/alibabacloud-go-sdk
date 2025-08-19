// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSnapshotSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSnapshotSettingResponseBody
	GetRequestId() *string
	SetResult(v *GetSnapshotSettingResponseBodyResult) *GetSnapshotSettingResponseBody
	GetResult() *GetSnapshotSettingResponseBodyResult
}

type GetSnapshotSettingResponseBody struct {
	// example:
	//
	// 7B6CE6E1-5BA0-56DA-BFFD-8D90692F1EFC
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetSnapshotSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSnapshotSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSnapshotSettingResponseBody) GetResult() *GetSnapshotSettingResponseBodyResult {
	return s.Result
}

func (s *GetSnapshotSettingResponseBody) SetRequestId(v string) *GetSnapshotSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSnapshotSettingResponseBody) SetResult(v *GetSnapshotSettingResponseBodyResult) *GetSnapshotSettingResponseBody {
	s.Result = v
	return s
}

func (s *GetSnapshotSettingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSnapshotSettingResponseBodyResult struct {
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 0 0 01 ? 	- 	- *
	QuartzRegex *string `json:"quartzRegex,omitempty" xml:"quartzRegex,omitempty"`
}

func (s GetSnapshotSettingResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSnapshotSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSnapshotSettingResponseBodyResult) GetEnable() *bool {
	return s.Enable
}

func (s *GetSnapshotSettingResponseBodyResult) GetQuartzRegex() *string {
	return s.QuartzRegex
}

func (s *GetSnapshotSettingResponseBodyResult) SetEnable(v bool) *GetSnapshotSettingResponseBodyResult {
	s.Enable = &v
	return s
}

func (s *GetSnapshotSettingResponseBodyResult) SetQuartzRegex(v string) *GetSnapshotSettingResponseBodyResult {
	s.QuartzRegex = &v
	return s
}

func (s *GetSnapshotSettingResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
