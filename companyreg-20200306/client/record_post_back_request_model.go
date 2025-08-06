// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordPostBackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *RecordPostBackRequest
	GetBizId() *string
	SetBizType(v string) *RecordPostBackRequest
	GetBizType() *string
	SetContactor(v string) *RecordPostBackRequest
	GetContactor() *string
	SetContent(v string) *RecordPostBackRequest
	GetContent() *string
	SetEntityKey(v string) *RecordPostBackRequest
	GetEntityKey() *string
}

type RecordPostBackRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// P111111111111
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.zhangsan
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	Contactor *string `json:"contactor,omitempty" xml:"contactor,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp_produce
	EntityKey *string `json:"entityKey,omitempty" xml:"entityKey,omitempty"`
}

func (s RecordPostBackRequest) String() string {
	return dara.Prettify(s)
}

func (s RecordPostBackRequest) GoString() string {
	return s.String()
}

func (s *RecordPostBackRequest) GetBizId() *string {
	return s.BizId
}

func (s *RecordPostBackRequest) GetBizType() *string {
	return s.BizType
}

func (s *RecordPostBackRequest) GetContactor() *string {
	return s.Contactor
}

func (s *RecordPostBackRequest) GetContent() *string {
	return s.Content
}

func (s *RecordPostBackRequest) GetEntityKey() *string {
	return s.EntityKey
}

func (s *RecordPostBackRequest) SetBizId(v string) *RecordPostBackRequest {
	s.BizId = &v
	return s
}

func (s *RecordPostBackRequest) SetBizType(v string) *RecordPostBackRequest {
	s.BizType = &v
	return s
}

func (s *RecordPostBackRequest) SetContactor(v string) *RecordPostBackRequest {
	s.Contactor = &v
	return s
}

func (s *RecordPostBackRequest) SetContent(v string) *RecordPostBackRequest {
	s.Content = &v
	return s
}

func (s *RecordPostBackRequest) SetEntityKey(v string) *RecordPostBackRequest {
	s.EntityKey = &v
	return s
}

func (s *RecordPostBackRequest) Validate() error {
	return dara.Validate(s)
}
