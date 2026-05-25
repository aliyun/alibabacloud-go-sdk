// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogTaskCollectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetTlogTaskCollectionsRequest
	GetAppKey() *int64
	SetOs(v string) *GetTlogTaskCollectionsRequest
	GetOs() *string
	SetTaskId(v string) *GetTlogTaskCollectionsRequest
	GetTaskId() *string
}

type GetTlogTaskCollectionsRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 24780725
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WXAIGC5ZCVUIYY78ABTCOUHGSQ
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTlogTaskCollectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTlogTaskCollectionsRequest) GoString() string {
	return s.String()
}

func (s *GetTlogTaskCollectionsRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetTlogTaskCollectionsRequest) GetOs() *string {
	return s.Os
}

func (s *GetTlogTaskCollectionsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTlogTaskCollectionsRequest) SetAppKey(v int64) *GetTlogTaskCollectionsRequest {
	s.AppKey = &v
	return s
}

func (s *GetTlogTaskCollectionsRequest) SetOs(v string) *GetTlogTaskCollectionsRequest {
	s.Os = &v
	return s
}

func (s *GetTlogTaskCollectionsRequest) SetTaskId(v string) *GetTlogTaskCollectionsRequest {
	s.TaskId = &v
	return s
}

func (s *GetTlogTaskCollectionsRequest) Validate() error {
	return dara.Validate(s)
}
