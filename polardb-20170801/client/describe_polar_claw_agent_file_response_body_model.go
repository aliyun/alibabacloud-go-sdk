// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DescribePolarClawAgentFileResponseBody
	GetAgentId() *string
	SetApplicationId(v string) *DescribePolarClawAgentFileResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribePolarClawAgentFileResponseBody
	GetCode() *int32
	SetFile(v *DescribePolarClawAgentFileResponseBodyFile) *DescribePolarClawAgentFileResponseBody
	GetFile() *DescribePolarClawAgentFileResponseBodyFile
	SetMessage(v string) *DescribePolarClawAgentFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePolarClawAgentFileResponseBody
	GetRequestId() *string
	SetWorkspace(v string) *DescribePolarClawAgentFileResponseBody
	GetWorkspace() *string
}

type DescribePolarClawAgentFileResponseBody struct {
	// Agent ID
	//
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The file details.
	File *DescribePolarClawAgentFileResponseBodyFile `json:"File,omitempty" xml:"File,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The working directory path.
	//
	// example:
	//
	// /home/node/.openclaw/workspace-main
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DescribePolarClawAgentFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentFileResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribePolarClawAgentFileResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawAgentFileResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawAgentFileResponseBody) GetFile() *DescribePolarClawAgentFileResponseBodyFile {
	return s.File
}

func (s *DescribePolarClawAgentFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawAgentFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawAgentFileResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *DescribePolarClawAgentFileResponseBody) SetAgentId(v string) *DescribePolarClawAgentFileResponseBody {
	s.AgentId = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBody) SetApplicationId(v string) *DescribePolarClawAgentFileResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBody) SetCode(v int32) *DescribePolarClawAgentFileResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBody) SetFile(v *DescribePolarClawAgentFileResponseBodyFile) *DescribePolarClawAgentFileResponseBody {
	s.File = v
	return s
}

func (s *DescribePolarClawAgentFileResponseBody) SetMessage(v string) *DescribePolarClawAgentFileResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBody) SetRequestId(v string) *DescribePolarClawAgentFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBody) SetWorkspace(v string) *DescribePolarClawAgentFileResponseBody {
	s.Workspace = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBody) Validate() error {
	if s.File != nil {
		if err := s.File.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarClawAgentFileResponseBodyFile struct {
	// The file content.
	//
	// example:
	//
	// You are a helpful assistant.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates whether the file is missing.
	//
	// example:
	//
	// false
	Missing *bool `json:"Missing,omitempty" xml:"Missing,omitempty"`
	// The file name.
	//
	// example:
	//
	// SOUL.md
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The file path.
	//
	// example:
	//
	// /home/node/.openclaw/workspace-main/SOUL.md
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The file size, in bytes.
	//
	// example:
	//
	// 1024
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The last updated UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 1716000000000
	UpdatedAtMs *int64 `json:"UpdatedAtMs,omitempty" xml:"UpdatedAtMs,omitempty"`
}

func (s DescribePolarClawAgentFileResponseBodyFile) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentFileResponseBodyFile) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentFileResponseBodyFile) GetContent() *string {
	return s.Content
}

func (s *DescribePolarClawAgentFileResponseBodyFile) GetMissing() *bool {
	return s.Missing
}

func (s *DescribePolarClawAgentFileResponseBodyFile) GetName() *string {
	return s.Name
}

func (s *DescribePolarClawAgentFileResponseBodyFile) GetPath() *string {
	return s.Path
}

func (s *DescribePolarClawAgentFileResponseBodyFile) GetSize() *int64 {
	return s.Size
}

func (s *DescribePolarClawAgentFileResponseBodyFile) GetUpdatedAtMs() *int64 {
	return s.UpdatedAtMs
}

func (s *DescribePolarClawAgentFileResponseBodyFile) SetContent(v string) *DescribePolarClawAgentFileResponseBodyFile {
	s.Content = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBodyFile) SetMissing(v bool) *DescribePolarClawAgentFileResponseBodyFile {
	s.Missing = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBodyFile) SetName(v string) *DescribePolarClawAgentFileResponseBodyFile {
	s.Name = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBodyFile) SetPath(v string) *DescribePolarClawAgentFileResponseBodyFile {
	s.Path = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBodyFile) SetSize(v int64) *DescribePolarClawAgentFileResponseBodyFile {
	s.Size = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBodyFile) SetUpdatedAtMs(v int64) *DescribePolarClawAgentFileResponseBodyFile {
	s.UpdatedAtMs = &v
	return s
}

func (s *DescribePolarClawAgentFileResponseBodyFile) Validate() error {
	return dara.Validate(s)
}
