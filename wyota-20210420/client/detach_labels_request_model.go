// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelIds(v string) *DetachLabelsRequest
	GetLabelIds() *string
	SetSerialNo(v string) *DetachLabelsRequest
	GetSerialNo() *string
	SetSerialNoList(v string) *DetachLabelsRequest
	GetSerialNoList() *string
}

type DetachLabelsRequest struct {
	LabelIds     *string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SerialNoList *string `json:"SerialNoList,omitempty" xml:"SerialNoList,omitempty"`
}

func (s DetachLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachLabelsRequest) GoString() string {
	return s.String()
}

func (s *DetachLabelsRequest) GetLabelIds() *string {
	return s.LabelIds
}

func (s *DetachLabelsRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *DetachLabelsRequest) GetSerialNoList() *string {
	return s.SerialNoList
}

func (s *DetachLabelsRequest) SetLabelIds(v string) *DetachLabelsRequest {
	s.LabelIds = &v
	return s
}

func (s *DetachLabelsRequest) SetSerialNo(v string) *DetachLabelsRequest {
	s.SerialNo = &v
	return s
}

func (s *DetachLabelsRequest) SetSerialNoList(v string) *DetachLabelsRequest {
	s.SerialNoList = &v
	return s
}

func (s *DetachLabelsRequest) Validate() error {
	return dara.Validate(s)
}
