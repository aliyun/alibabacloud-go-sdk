// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSyncCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *RunSyncCommandRequest
	GetCommandContent() *string
	SetContentEncoding(v string) *RunSyncCommandRequest
	GetContentEncoding() *string
	SetInstanceIds(v []*string) *RunSyncCommandRequest
	GetInstanceIds() []*string
	SetWaitTime(v int64) *RunSyncCommandRequest
	GetWaitTime() *int64
}

type RunSyncCommandRequest struct {
	// example:
	//
	// ls
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// example:
	//
	// PlainText
	ContentEncoding *string   `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	WaitTime *int64 `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s RunSyncCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSyncCommandRequest) GoString() string {
	return s.String()
}

func (s *RunSyncCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunSyncCommandRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunSyncCommandRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RunSyncCommandRequest) GetWaitTime() *int64 {
	return s.WaitTime
}

func (s *RunSyncCommandRequest) SetCommandContent(v string) *RunSyncCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunSyncCommandRequest) SetContentEncoding(v string) *RunSyncCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunSyncCommandRequest) SetInstanceIds(v []*string) *RunSyncCommandRequest {
	s.InstanceIds = v
	return s
}

func (s *RunSyncCommandRequest) SetWaitTime(v int64) *RunSyncCommandRequest {
	s.WaitTime = &v
	return s
}

func (s *RunSyncCommandRequest) Validate() error {
	return dara.Validate(s)
}
