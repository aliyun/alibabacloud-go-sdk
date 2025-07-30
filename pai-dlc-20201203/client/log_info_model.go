// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogInfo interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *LogInfo
	GetContent() *string
	SetId(v string) *LogInfo
	GetId() *string
	SetIsTruncated(v bool) *LogInfo
	GetIsTruncated() *bool
	SetPodId(v string) *LogInfo
	GetPodId() *string
	SetPodUid(v string) *LogInfo
	GetPodUid() *string
	SetSource(v string) *LogInfo
	GetSource() *string
	SetTime(v string) *LogInfo
	GetTime() *string
}

type LogInfo struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsTruncated *bool   `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	PodId       *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid      *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// example:
	//
	// stderr, stdout
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Time   *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s LogInfo) String() string {
	return dara.Prettify(s)
}

func (s LogInfo) GoString() string {
	return s.String()
}

func (s *LogInfo) GetContent() *string {
	return s.Content
}

func (s *LogInfo) GetId() *string {
	return s.Id
}

func (s *LogInfo) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *LogInfo) GetPodId() *string {
	return s.PodId
}

func (s *LogInfo) GetPodUid() *string {
	return s.PodUid
}

func (s *LogInfo) GetSource() *string {
	return s.Source
}

func (s *LogInfo) GetTime() *string {
	return s.Time
}

func (s *LogInfo) SetContent(v string) *LogInfo {
	s.Content = &v
	return s
}

func (s *LogInfo) SetId(v string) *LogInfo {
	s.Id = &v
	return s
}

func (s *LogInfo) SetIsTruncated(v bool) *LogInfo {
	s.IsTruncated = &v
	return s
}

func (s *LogInfo) SetPodId(v string) *LogInfo {
	s.PodId = &v
	return s
}

func (s *LogInfo) SetPodUid(v string) *LogInfo {
	s.PodUid = &v
	return s
}

func (s *LogInfo) SetSource(v string) *LogInfo {
	s.Source = &v
	return s
}

func (s *LogInfo) SetTime(v string) *LogInfo {
	s.Time = &v
	return s
}

func (s *LogInfo) Validate() error {
	return dara.Validate(s)
}
