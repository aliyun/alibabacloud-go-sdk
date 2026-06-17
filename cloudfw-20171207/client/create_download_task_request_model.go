// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateDownloadTaskRequest
	GetLang() *string
	SetTaskData(v string) *CreateDownloadTaskRequest
	GetTaskData() *string
	SetTaskType(v string) *CreateDownloadTaskRequest
	GetTaskType() *string
	SetTimeZone(v string) *CreateDownloadTaskRequest
	GetTimeZone() *string
}

type CreateDownloadTaskRequest struct {
	// The language of the response message.
	//
	// Valid values:
	//
	// - **zh**: Chinese (default)
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The query conditions for the download task.
	//
	// example:
	//
	// {\\"SearchItem\\":\\"\\",\\"UserType\\":\\"buy\\",\\"IpVersion\\":\\"4\\"}
	TaskData *string `json:"TaskData,omitempty" xml:"TaskData,omitempty"`
	// The type of the task. For a list of valid values, see the API for querying file download task types.
	//
	// example:
	//
	// InternetFirewallAsset
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The time zone of the time information in the downloaded file. The value is an IANA time zone identifier. The default value is Asia/Shanghai, which corresponds to UTC+8.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreateDownloadTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDownloadTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDownloadTaskRequest) GetTaskData() *string {
	return s.TaskData
}

func (s *CreateDownloadTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateDownloadTaskRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateDownloadTaskRequest) SetLang(v string) *CreateDownloadTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateDownloadTaskRequest) SetTaskData(v string) *CreateDownloadTaskRequest {
	s.TaskData = &v
	return s
}

func (s *CreateDownloadTaskRequest) SetTaskType(v string) *CreateDownloadTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateDownloadTaskRequest) SetTimeZone(v string) *CreateDownloadTaskRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateDownloadTaskRequest) Validate() error {
	return dara.Validate(s)
}
