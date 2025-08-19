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
	SetResult(v map[string]interface{}) *UpdateSnapshotSettingResponseBody
	GetResult() map[string]interface{}
}

type UpdateSnapshotSettingResponseBody struct {
	// example:
	//
	// A7B03723-AA73-5A5F-B71C-270792614DD8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {
	//
	//     "quartzRegex": "0 0 01 ? 	- 	- *",
	//
	//     "enable": true
	//
	//   }
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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

func (s *UpdateSnapshotSettingResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *UpdateSnapshotSettingResponseBody) SetRequestId(v string) *UpdateSnapshotSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSnapshotSettingResponseBody) SetResult(v map[string]interface{}) *UpdateSnapshotSettingResponseBody {
	s.Result = v
	return s
}

func (s *UpdateSnapshotSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
