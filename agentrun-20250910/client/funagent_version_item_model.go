// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunagentVersionItem interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *FunagentVersionItem
	GetDescription() *string
	SetPublishContent(v []*string) *FunagentVersionItem
	GetPublishContent() []*string
	SetPublishTime(v string) *FunagentVersionItem
	GetPublishTime() *string
	SetVersion(v string) *FunagentVersionItem
	GetVersion() *string
}

type FunagentVersionItem struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 多条更新说明
	PublishContent []*string `json:"publishContent" xml:"publishContent" type:"Repeated"`
	// 日期或 ISO 8601 字符串
	PublishTime *string `json:"publishTime,omitempty" xml:"publishTime,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s FunagentVersionItem) String() string {
	return dara.Prettify(s)
}

func (s FunagentVersionItem) GoString() string {
	return s.String()
}

func (s *FunagentVersionItem) GetDescription() *string {
	return s.Description
}

func (s *FunagentVersionItem) GetPublishContent() []*string {
	return s.PublishContent
}

func (s *FunagentVersionItem) GetPublishTime() *string {
	return s.PublishTime
}

func (s *FunagentVersionItem) GetVersion() *string {
	return s.Version
}

func (s *FunagentVersionItem) SetDescription(v string) *FunagentVersionItem {
	s.Description = &v
	return s
}

func (s *FunagentVersionItem) SetPublishContent(v []*string) *FunagentVersionItem {
	s.PublishContent = v
	return s
}

func (s *FunagentVersionItem) SetPublishTime(v string) *FunagentVersionItem {
	s.PublishTime = &v
	return s
}

func (s *FunagentVersionItem) SetVersion(v string) *FunagentVersionItem {
	s.Version = &v
	return s
}

func (s *FunagentVersionItem) Validate() error {
	return dara.Validate(s)
}
