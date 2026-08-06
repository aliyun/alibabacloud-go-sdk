// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrompts(v []*Prompt) *ListPromptsResponseBody
	GetPrompts() []*Prompt
	SetRequestId(v string) *ListPromptsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPromptsResponseBody
	GetTotalCount() *int32
}

type ListPromptsResponseBody struct {
	// The list of prompts.
	Prompts []*Prompt `json:"Prompts,omitempty" xml:"Prompts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D5BFFEE3-6025-443F-8A03-02D619B5C4B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned under the current request conditions. This parameter is optional and may not be returned by default.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPromptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPromptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromptsResponseBody) GetPrompts() []*Prompt {
	return s.Prompts
}

func (s *ListPromptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPromptsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPromptsResponseBody) SetPrompts(v []*Prompt) *ListPromptsResponseBody {
	s.Prompts = v
	return s
}

func (s *ListPromptsResponseBody) SetRequestId(v string) *ListPromptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPromptsResponseBody) SetTotalCount(v int32) *ListPromptsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPromptsResponseBody) Validate() error {
	if s.Prompts != nil {
		for _, item := range s.Prompts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
