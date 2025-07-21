// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelContent(v string) *AttachLabelRequest
	GetLabelContent() *string
	SetLabelId(v string) *AttachLabelRequest
	GetLabelId() *string
	SetSerialNo(v string) *AttachLabelRequest
	GetSerialNo() *string
}

type AttachLabelRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AttachLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachLabelRequest) GoString() string {
	return s.String()
}

func (s *AttachLabelRequest) GetLabelContent() *string {
	return s.LabelContent
}

func (s *AttachLabelRequest) GetLabelId() *string {
	return s.LabelId
}

func (s *AttachLabelRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *AttachLabelRequest) SetLabelContent(v string) *AttachLabelRequest {
	s.LabelContent = &v
	return s
}

func (s *AttachLabelRequest) SetLabelId(v string) *AttachLabelRequest {
	s.LabelId = &v
	return s
}

func (s *AttachLabelRequest) SetSerialNo(v string) *AttachLabelRequest {
	s.SerialNo = &v
	return s
}

func (s *AttachLabelRequest) Validate() error {
	return dara.Validate(s)
}
