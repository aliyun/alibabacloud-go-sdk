// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronDraftRecordUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *NeuronDraftRecordUpdateCmd
	GetContent() *string
	SetId(v int64) *NeuronDraftRecordUpdateCmd
	GetId() *int64
	SetName(v string) *NeuronDraftRecordUpdateCmd
	GetName() *string
}

type NeuronDraftRecordUpdateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// {}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// app版本
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s NeuronDraftRecordUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronDraftRecordUpdateCmd) GoString() string {
	return s.String()
}

func (s *NeuronDraftRecordUpdateCmd) GetContent() *string {
	return s.Content
}

func (s *NeuronDraftRecordUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *NeuronDraftRecordUpdateCmd) GetName() *string {
	return s.Name
}

func (s *NeuronDraftRecordUpdateCmd) SetContent(v string) *NeuronDraftRecordUpdateCmd {
	s.Content = &v
	return s
}

func (s *NeuronDraftRecordUpdateCmd) SetId(v int64) *NeuronDraftRecordUpdateCmd {
	s.Id = &v
	return s
}

func (s *NeuronDraftRecordUpdateCmd) SetName(v string) *NeuronDraftRecordUpdateCmd {
	s.Name = &v
	return s
}

func (s *NeuronDraftRecordUpdateCmd) Validate() error {
	return dara.Validate(s)
}
