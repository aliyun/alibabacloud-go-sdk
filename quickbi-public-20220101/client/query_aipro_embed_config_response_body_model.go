// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiproEmbedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryAiproEmbedConfigResponseBody
	GetRequestId() *string
	SetResult(v []*QueryAiproEmbedConfigResponseBodyResult) *QueryAiproEmbedConfigResponseBody
	GetResult() []*QueryAiproEmbedConfigResponseBodyResult
	SetSuccess(v bool) *QueryAiproEmbedConfigResponseBody
	GetSuccess() *bool
}

type QueryAiproEmbedConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D8749***C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of AIPro configurations.
	Result []*QueryAiproEmbedConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAiproEmbedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAiproEmbedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiproEmbedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAiproEmbedConfigResponseBody) GetResult() []*QueryAiproEmbedConfigResponseBodyResult {
	return s.Result
}

func (s *QueryAiproEmbedConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAiproEmbedConfigResponseBody) SetRequestId(v string) *QueryAiproEmbedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBody) SetResult(v []*QueryAiproEmbedConfigResponseBodyResult) *QueryAiproEmbedConfigResponseBody {
	s.Result = v
	return s
}

func (s *QueryAiproEmbedConfigResponseBody) SetSuccess(v bool) *QueryAiproEmbedConfigResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBody) Validate() error {
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

type QueryAiproEmbedConfigResponseBodyResult struct {
	// The agent nickname.
	//
	// example:
	//
	// Xiao Q
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The embed configuration ID, which is used for CreateTicket4AiPro.
	//
	// example:
	//
	// 123124
	AiProId *string `json:"AiProId,omitempty" xml:"AiProId,omitempty"`
	// Indicates whether chat history is displayed. Valid values:
	//
	// - 1: Displayed.
	//
	// - 0: Not displayed.
	//
	// example:
	//
	// 1
	ChatHistoryStatus *int32 `json:"ChatHistoryStatus,omitempty" xml:"ChatHistoryStatus,omitempty"`
	// The JSON string of the embed appearance configuration.
	//
	// example:
	//
	// test
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether memory is enabled. Valid values:
	//
	// - 1: Enabled.
	//
	// - 0: Disabled.
	//
	// example:
	//
	// 1
	MemoryEnabled *int32 `json:"MemoryEnabled,omitempty" xml:"MemoryEnabled,omitempty"`
	// The configuration name.
	//
	// example:
	//
	// 0327
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The status. Valid values:
	//
	// - 1: Enabled.
	//
	// - 0: Disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether file upload is allowed. Valid values:
	//
	// - 1: Allowed.
	//
	// - 0: Not allowed.
	//
	// example:
	//
	// 1
	UploadEnabled *int32 `json:"UploadEnabled,omitempty" xml:"UploadEnabled,omitempty"`
}

func (s QueryAiproEmbedConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryAiproEmbedConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAiproEmbedConfigResponseBodyResult) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryAiproEmbedConfigResponseBodyResult) GetAiProId() *string {
	return s.AiProId
}

func (s *QueryAiproEmbedConfigResponseBodyResult) GetChatHistoryStatus() *int32 {
	return s.ChatHistoryStatus
}

func (s *QueryAiproEmbedConfigResponseBodyResult) GetConfig() *string {
	return s.Config
}

func (s *QueryAiproEmbedConfigResponseBodyResult) GetMemoryEnabled() *int32 {
	return s.MemoryEnabled
}

func (s *QueryAiproEmbedConfigResponseBodyResult) GetShowName() *string {
	return s.ShowName
}

func (s *QueryAiproEmbedConfigResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *QueryAiproEmbedConfigResponseBodyResult) GetUploadEnabled() *int32 {
	return s.UploadEnabled
}

func (s *QueryAiproEmbedConfigResponseBodyResult) SetAgentName(v string) *QueryAiproEmbedConfigResponseBodyResult {
	s.AgentName = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBodyResult) SetAiProId(v string) *QueryAiproEmbedConfigResponseBodyResult {
	s.AiProId = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBodyResult) SetChatHistoryStatus(v int32) *QueryAiproEmbedConfigResponseBodyResult {
	s.ChatHistoryStatus = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBodyResult) SetConfig(v string) *QueryAiproEmbedConfigResponseBodyResult {
	s.Config = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBodyResult) SetMemoryEnabled(v int32) *QueryAiproEmbedConfigResponseBodyResult {
	s.MemoryEnabled = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBodyResult) SetShowName(v string) *QueryAiproEmbedConfigResponseBodyResult {
	s.ShowName = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBodyResult) SetStatus(v int32) *QueryAiproEmbedConfigResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBodyResult) SetUploadEnabled(v int32) *QueryAiproEmbedConfigResponseBodyResult {
	s.UploadEnabled = &v
	return s
}

func (s *QueryAiproEmbedConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
