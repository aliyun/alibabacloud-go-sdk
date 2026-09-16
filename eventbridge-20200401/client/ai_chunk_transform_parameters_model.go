// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiChunkTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetChunkingType(v string) *AiChunkTransformParameters
	GetChunkingType() *string
	SetInputField(v *AiTransformField) *AiChunkTransformParameters
	GetInputField() *AiTransformField
	SetMaxOverlapSize(v int32) *AiChunkTransformParameters
	GetMaxOverlapSize() *int32
	SetMaxSegmentSize(v int32) *AiChunkTransformParameters
	GetMaxSegmentSize() *int32
	SetStepName(v string) *AiChunkTransformParameters
	GetStepName() *string
	SetUnpack(v bool) *AiChunkTransformParameters
	GetUnpack() *bool
}

type AiChunkTransformParameters struct {
	// The chunking algorithm. Valid values: markdown_header and recursive_character.
	//
	// example:
	//
	// recursive_character
	ChunkingType *string `json:"ChunkingType,omitempty" xml:"ChunkingType,omitempty"`
	// The text field to chunk. This operator does not use InputField.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The chunk overlap length. Default value: 10.
	//
	// example:
	//
	// 10
	MaxOverlapSize *int32 `json:"MaxOverlapSize,omitempty" xml:"MaxOverlapSize,omitempty"`
	// The maximum chunk length. Default value: 1000.
	//
	// example:
	//
	// 1000
	MaxSegmentSize *int32 `json:"MaxSegmentSize,omitempty" xml:"MaxSegmentSize,omitempty"`
	// The field name in the CloudEvent to which the output is attached. Default value: transform0.
	//
	// example:
	//
	// chunks
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// Specifies whether to split the output into multiple events. Default value: true.
	//
	// example:
	//
	// true
	Unpack *bool `json:"Unpack,omitempty" xml:"Unpack,omitempty"`
}

func (s AiChunkTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiChunkTransformParameters) GoString() string {
	return s.String()
}

func (s *AiChunkTransformParameters) GetChunkingType() *string {
	return s.ChunkingType
}

func (s *AiChunkTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiChunkTransformParameters) GetMaxOverlapSize() *int32 {
	return s.MaxOverlapSize
}

func (s *AiChunkTransformParameters) GetMaxSegmentSize() *int32 {
	return s.MaxSegmentSize
}

func (s *AiChunkTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiChunkTransformParameters) GetUnpack() *bool {
	return s.Unpack
}

func (s *AiChunkTransformParameters) SetChunkingType(v string) *AiChunkTransformParameters {
	s.ChunkingType = &v
	return s
}

func (s *AiChunkTransformParameters) SetInputField(v *AiTransformField) *AiChunkTransformParameters {
	s.InputField = v
	return s
}

func (s *AiChunkTransformParameters) SetMaxOverlapSize(v int32) *AiChunkTransformParameters {
	s.MaxOverlapSize = &v
	return s
}

func (s *AiChunkTransformParameters) SetMaxSegmentSize(v int32) *AiChunkTransformParameters {
	s.MaxSegmentSize = &v
	return s
}

func (s *AiChunkTransformParameters) SetStepName(v string) *AiChunkTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiChunkTransformParameters) SetUnpack(v bool) *AiChunkTransformParameters {
	s.Unpack = &v
	return s
}

func (s *AiChunkTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
