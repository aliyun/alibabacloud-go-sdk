// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasAgentSSEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v string) *GetDasAgentSSEResponseBody
	GetAnswer() *string
	SetEvent(v string) *GetDasAgentSSEResponseBody
	GetEvent() *string
	SetId(v string) *GetDasAgentSSEResponseBody
	GetId() *string
	SetMetadata(v *GetDasAgentSSEResponseBodyMetadata) *GetDasAgentSSEResponseBody
	GetMetadata() *GetDasAgentSSEResponseBodyMetadata
}

type GetDasAgentSSEResponseBody struct {
	// example:
	//
	// rm-xxxx
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// summary
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// xxx-xxx-xxx
	Id       *string                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Metadata *GetDasAgentSSEResponseBodyMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
}

func (s GetDasAgentSSEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDasAgentSSEResponseBody) GoString() string {
	return s.String()
}

func (s *GetDasAgentSSEResponseBody) GetAnswer() *string {
	return s.Answer
}

func (s *GetDasAgentSSEResponseBody) GetEvent() *string {
	return s.Event
}

func (s *GetDasAgentSSEResponseBody) GetId() *string {
	return s.Id
}

func (s *GetDasAgentSSEResponseBody) GetMetadata() *GetDasAgentSSEResponseBodyMetadata {
	return s.Metadata
}

func (s *GetDasAgentSSEResponseBody) SetAnswer(v string) *GetDasAgentSSEResponseBody {
	s.Answer = &v
	return s
}

func (s *GetDasAgentSSEResponseBody) SetEvent(v string) *GetDasAgentSSEResponseBody {
	s.Event = &v
	return s
}

func (s *GetDasAgentSSEResponseBody) SetId(v string) *GetDasAgentSSEResponseBody {
	s.Id = &v
	return s
}

func (s *GetDasAgentSSEResponseBody) SetMetadata(v *GetDasAgentSSEResponseBodyMetadata) *GetDasAgentSSEResponseBody {
	s.Metadata = v
	return s
}

func (s *GetDasAgentSSEResponseBody) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDasAgentSSEResponseBodyMetadata struct {
	// example:
	//
	// 1403
	CharCount *int64 `json:"CharCount,omitempty" xml:"CharCount,omitempty"`
	Code      *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 04930480-9404-50CB-8252-Axxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SqlFilter
	ToolName   *string   `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolParams []*string `json:"ToolParams,omitempty" xml:"ToolParams,omitempty" type:"Repeated"`
}

func (s GetDasAgentSSEResponseBodyMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetDasAgentSSEResponseBodyMetadata) GoString() string {
	return s.String()
}

func (s *GetDasAgentSSEResponseBodyMetadata) GetCharCount() *int64 {
	return s.CharCount
}

func (s *GetDasAgentSSEResponseBodyMetadata) GetCode() *int32 {
	return s.Code
}

func (s *GetDasAgentSSEResponseBodyMetadata) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDasAgentSSEResponseBodyMetadata) GetToolName() *string {
	return s.ToolName
}

func (s *GetDasAgentSSEResponseBodyMetadata) GetToolParams() []*string {
	return s.ToolParams
}

func (s *GetDasAgentSSEResponseBodyMetadata) SetCharCount(v int64) *GetDasAgentSSEResponseBodyMetadata {
	s.CharCount = &v
	return s
}

func (s *GetDasAgentSSEResponseBodyMetadata) SetCode(v int32) *GetDasAgentSSEResponseBodyMetadata {
	s.Code = &v
	return s
}

func (s *GetDasAgentSSEResponseBodyMetadata) SetRequestId(v string) *GetDasAgentSSEResponseBodyMetadata {
	s.RequestId = &v
	return s
}

func (s *GetDasAgentSSEResponseBodyMetadata) SetToolName(v string) *GetDasAgentSSEResponseBodyMetadata {
	s.ToolName = &v
	return s
}

func (s *GetDasAgentSSEResponseBodyMetadata) SetToolParams(v []*string) *GetDasAgentSSEResponseBodyMetadata {
	s.ToolParams = v
	return s
}

func (s *GetDasAgentSSEResponseBodyMetadata) Validate() error {
	return dara.Validate(s)
}
