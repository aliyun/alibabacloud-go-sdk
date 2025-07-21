// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelContent(v string) *DetachLabelRequest
	GetLabelContent() *string
	SetLabelId(v string) *DetachLabelRequest
	GetLabelId() *string
	SetSerialNo(v string) *DetachLabelRequest
	GetSerialNo() *string
}

type DetachLabelRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s DetachLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachLabelRequest) GoString() string {
	return s.String()
}

func (s *DetachLabelRequest) GetLabelContent() *string {
	return s.LabelContent
}

func (s *DetachLabelRequest) GetLabelId() *string {
	return s.LabelId
}

func (s *DetachLabelRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *DetachLabelRequest) SetLabelContent(v string) *DetachLabelRequest {
	s.LabelContent = &v
	return s
}

func (s *DetachLabelRequest) SetLabelId(v string) *DetachLabelRequest {
	s.LabelId = &v
	return s
}

func (s *DetachLabelRequest) SetSerialNo(v string) *DetachLabelRequest {
	s.SerialNo = &v
	return s
}

func (s *DetachLabelRequest) Validate() error {
	return dara.Validate(s)
}
