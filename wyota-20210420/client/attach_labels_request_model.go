// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelIds(v string) *AttachLabelsRequest
	GetLabelIds() *string
	SetSerialNo(v string) *AttachLabelsRequest
	GetSerialNo() *string
	SetSerialNoList(v string) *AttachLabelsRequest
	GetSerialNoList() *string
}

type AttachLabelsRequest struct {
	LabelIds     *string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SerialNoList *string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty"`
}

func (s AttachLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachLabelsRequest) GoString() string {
	return s.String()
}

func (s *AttachLabelsRequest) GetLabelIds() *string {
	return s.LabelIds
}

func (s *AttachLabelsRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *AttachLabelsRequest) GetSerialNoList() *string {
	return s.SerialNoList
}

func (s *AttachLabelsRequest) SetLabelIds(v string) *AttachLabelsRequest {
	s.LabelIds = &v
	return s
}

func (s *AttachLabelsRequest) SetSerialNo(v string) *AttachLabelsRequest {
	s.SerialNo = &v
	return s
}

func (s *AttachLabelsRequest) SetSerialNoList(v string) *AttachLabelsRequest {
	s.SerialNoList = &v
	return s
}

func (s *AttachLabelsRequest) Validate() error {
	return dara.Validate(s)
}
