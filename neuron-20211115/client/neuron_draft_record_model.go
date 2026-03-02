// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronDraftRecord interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *NeuronDraftRecord
	GetBizId() *string
	SetContent(v string) *NeuronDraftRecord
	GetContent() *string
	SetGmtCreate(v string) *NeuronDraftRecord
	GetGmtCreate() *string
	SetGmtModified(v string) *NeuronDraftRecord
	GetGmtModified() *string
	SetId(v int64) *NeuronDraftRecord
	GetId() *int64
	SetName(v string) *NeuronDraftRecord
	GetName() *string
	SetType(v string) *NeuronDraftRecord
	GetType() *string
}

type NeuronDraftRecord struct {
	// example:
	//
	// 1
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {}
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// app版本
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// APP_VERSION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s NeuronDraftRecord) String() string {
	return dara.Prettify(s)
}

func (s NeuronDraftRecord) GoString() string {
	return s.String()
}

func (s *NeuronDraftRecord) GetBizId() *string {
	return s.BizId
}

func (s *NeuronDraftRecord) GetContent() *string {
	return s.Content
}

func (s *NeuronDraftRecord) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *NeuronDraftRecord) GetGmtModified() *string {
	return s.GmtModified
}

func (s *NeuronDraftRecord) GetId() *int64 {
	return s.Id
}

func (s *NeuronDraftRecord) GetName() *string {
	return s.Name
}

func (s *NeuronDraftRecord) GetType() *string {
	return s.Type
}

func (s *NeuronDraftRecord) SetBizId(v string) *NeuronDraftRecord {
	s.BizId = &v
	return s
}

func (s *NeuronDraftRecord) SetContent(v string) *NeuronDraftRecord {
	s.Content = &v
	return s
}

func (s *NeuronDraftRecord) SetGmtCreate(v string) *NeuronDraftRecord {
	s.GmtCreate = &v
	return s
}

func (s *NeuronDraftRecord) SetGmtModified(v string) *NeuronDraftRecord {
	s.GmtModified = &v
	return s
}

func (s *NeuronDraftRecord) SetId(v int64) *NeuronDraftRecord {
	s.Id = &v
	return s
}

func (s *NeuronDraftRecord) SetName(v string) *NeuronDraftRecord {
	s.Name = &v
	return s
}

func (s *NeuronDraftRecord) SetType(v string) *NeuronDraftRecord {
	s.Type = &v
	return s
}

func (s *NeuronDraftRecord) Validate() error {
	return dara.Validate(s)
}
