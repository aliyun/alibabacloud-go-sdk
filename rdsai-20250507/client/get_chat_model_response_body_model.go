// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetChatModelResponseBodyData) *GetChatModelResponseBody
	GetData() []*GetChatModelResponseBodyData
	SetRequestId(v string) *GetChatModelResponseBody
	GetRequestId() *string
}

type GetChatModelResponseBody struct {
	Data []*GetChatModelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatModelResponseBody) GetData() []*GetChatModelResponseBodyData {
	return s.Data
}

func (s *GetChatModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatModelResponseBody) SetData(v []*GetChatModelResponseBodyData) *GetChatModelResponseBody {
	s.Data = v
	return s
}

func (s *GetChatModelResponseBody) SetRequestId(v string) *GetChatModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatModelResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetChatModelResponseBodyData struct {
	// example:
	//
	// 1000000
	ContextWindow *int64 `json:"ContextWindow,omitempty" xml:"ContextWindow,omitempty"`
	// example:
	//
	// true
	Default  *bool     `json:"Default,omitempty" xml:"Default,omitempty"`
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// example:
	//
	// qwen3.7-max
	ModelId        *string   `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ThinkingLevels []*string `json:"ThinkingLevels,omitempty" xml:"ThinkingLevels,omitempty" type:"Repeated"`
}

func (s GetChatModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatModelResponseBodyData) GetContextWindow() *int64 {
	return s.ContextWindow
}

func (s *GetChatModelResponseBodyData) GetDefault() *bool {
	return s.Default
}

func (s *GetChatModelResponseBodyData) GetFeatures() []*string {
	return s.Features
}

func (s *GetChatModelResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *GetChatModelResponseBodyData) GetThinkingLevels() []*string {
	return s.ThinkingLevels
}

func (s *GetChatModelResponseBodyData) SetContextWindow(v int64) *GetChatModelResponseBodyData {
	s.ContextWindow = &v
	return s
}

func (s *GetChatModelResponseBodyData) SetDefault(v bool) *GetChatModelResponseBodyData {
	s.Default = &v
	return s
}

func (s *GetChatModelResponseBodyData) SetFeatures(v []*string) *GetChatModelResponseBodyData {
	s.Features = v
	return s
}

func (s *GetChatModelResponseBodyData) SetModelId(v string) *GetChatModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetChatModelResponseBodyData) SetThinkingLevels(v []*string) *GetChatModelResponseBodyData {
	s.ThinkingLevels = v
	return s
}

func (s *GetChatModelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
