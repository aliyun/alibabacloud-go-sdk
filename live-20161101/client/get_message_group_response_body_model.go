// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMessageGroupResponseBody
	GetRequestId() *string
	SetResult(v *GetMessageGroupResponseBodyResult) *GetMessageGroupResponseBody
	GetResult() *GetMessageGroupResponseBodyResult
}

type GetMessageGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *GetMessageGroupResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageGroupResponseBody) GetResult() *GetMessageGroupResponseBodyResult {
	return s.Result
}

func (s *GetMessageGroupResponseBody) SetRequestId(v string) *GetMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageGroupResponseBody) SetResult(v *GetMessageGroupResponseBodyResult) *GetMessageGroupResponseBody {
	s.Result = v
	return s
}

func (s *GetMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMessageGroupResponseBodyResult struct {
	// The time when the message group was created. The time is displayed in UTC.
	//
	// example:
	//
	// 1502280113
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// as****hs
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The extended field.
	//
	// example:
	//
	// 1
	Extension map[string]interface{} `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The ID of the message group.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether the message group is muted.
	//
	// 	- true: The message group is muted.
	//
	// 	- false: The message group is not muted.
	//
	// example:
	//
	// true
	IsMuteAll *bool `json:"IsMuteAll,omitempty" xml:"IsMuteAll,omitempty"`
	// The status of the message group. The default value is **1**, which indicates that the message group is normal.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMessageGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMessageGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMessageGroupResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMessageGroupResponseBodyResult) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetMessageGroupResponseBodyResult) GetExtension() map[string]interface{} {
	return s.Extension
}

func (s *GetMessageGroupResponseBodyResult) GetGroupId() *string {
	return s.GroupId
}

func (s *GetMessageGroupResponseBodyResult) GetIsMuteAll() *bool {
	return s.IsMuteAll
}

func (s *GetMessageGroupResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *GetMessageGroupResponseBodyResult) SetCreateTime(v int64) *GetMessageGroupResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetMessageGroupResponseBodyResult) SetCreatorId(v string) *GetMessageGroupResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *GetMessageGroupResponseBodyResult) SetExtension(v map[string]interface{}) *GetMessageGroupResponseBodyResult {
	s.Extension = v
	return s
}

func (s *GetMessageGroupResponseBodyResult) SetGroupId(v string) *GetMessageGroupResponseBodyResult {
	s.GroupId = &v
	return s
}

func (s *GetMessageGroupResponseBodyResult) SetIsMuteAll(v bool) *GetMessageGroupResponseBodyResult {
	s.IsMuteAll = &v
	return s
}

func (s *GetMessageGroupResponseBodyResult) SetStatus(v int32) *GetMessageGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetMessageGroupResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
