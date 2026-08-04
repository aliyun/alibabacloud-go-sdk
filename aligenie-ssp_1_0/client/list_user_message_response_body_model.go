// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserMessageResponseBody
	GetCode() *string
	SetMessage(v string) *ListUserMessageResponseBody
	GetMessage() *string
	SetResult(v []*ListUserMessageResponseBodyResult) *ListUserMessageResponseBody
	GetResult() []*ListUserMessageResponseBodyResult
}

type ListUserMessageResponseBody struct {
	// Status code returned by the service. SUCCESS indicates success; otherwise, it indicates failure.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	//
	// example:
	//
	// 外部userId映射关系不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// List of user message query results
	Result []*ListUserMessageResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListUserMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserMessageResponseBody) GetResult() []*ListUserMessageResponseBodyResult {
	return s.Result
}

func (s *ListUserMessageResponseBody) SetCode(v string) *ListUserMessageResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserMessageResponseBody) SetMessage(v string) *ListUserMessageResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserMessageResponseBody) SetResult(v []*ListUserMessageResponseBodyResult) *ListUserMessageResponseBody {
	s.Result = v
	return s
}

func (s *ListUserMessageResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserMessageResponseBodyResult struct {
	// Message text
	//
	// example:
	//
	// 哈哈哈
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Device name
	//
	// example:
	//
	// 卧室的小芳
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// Time when the message was sent
	//
	// example:
	//
	// 2022-07-27 14:06:27.000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Message ID
	//
	// example:
	//
	// 123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Device Image
	//
	// example:
	//
	// http://xx
	Pic *string `json:"Pic,omitempty" xml:"Pic,omitempty"`
	// Message source: app or box
	//
	// example:
	//
	// app
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Source Device ID
	//
	// example:
	//
	// AF188**065EE4B**DD68CE**951D84D4
	SourceUuid *string `json:"SourceUuid,omitempty" xml:"SourceUuid,omitempty"`
	// Message status: 0 indicates unread, and 1 indicates read.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Currently only audio is supported.
	//
	// example:
	//
	// audio
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Audio message link
	//
	// example:
	//
	// http://xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListUserMessageResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserMessageResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *ListUserMessageResponseBodyResult) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ListUserMessageResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListUserMessageResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListUserMessageResponseBodyResult) GetPic() *string {
	return s.Pic
}

func (s *ListUserMessageResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListUserMessageResponseBodyResult) GetSourceUuid() *string {
	return s.SourceUuid
}

func (s *ListUserMessageResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *ListUserMessageResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListUserMessageResponseBodyResult) GetUrl() *string {
	return s.Url
}

func (s *ListUserMessageResponseBodyResult) SetContent(v string) *ListUserMessageResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetDeviceName(v string) *ListUserMessageResponseBodyResult {
	s.DeviceName = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetGmtCreate(v string) *ListUserMessageResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetId(v string) *ListUserMessageResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetPic(v string) *ListUserMessageResponseBodyResult {
	s.Pic = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetSource(v string) *ListUserMessageResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetSourceUuid(v string) *ListUserMessageResponseBodyResult {
	s.SourceUuid = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetStatus(v int32) *ListUserMessageResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetType(v string) *ListUserMessageResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) SetUrl(v string) *ListUserMessageResponseBodyResult {
	s.Url = &v
	return s
}

func (s *ListUserMessageResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
