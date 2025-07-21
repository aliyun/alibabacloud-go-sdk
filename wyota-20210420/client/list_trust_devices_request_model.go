// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelContent(v string) *ListTrustDevicesRequest
	GetLabelContent() *string
	SetLabelId(v string) *ListTrustDevicesRequest
	GetLabelId() *string
	SetSerialNo(v string) *ListTrustDevicesRequest
	GetSerialNo() *string
	SetUserCustomId(v string) *ListTrustDevicesRequest
	GetUserCustomId() *string
}

type ListTrustDevicesRequest struct {
	LabelContent *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId      *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	UserCustomId *string `json:"UserCustomId,omitempty" xml:"UserCustomId,omitempty"`
}

func (s ListTrustDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrustDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListTrustDevicesRequest) GetLabelContent() *string {
	return s.LabelContent
}

func (s *ListTrustDevicesRequest) GetLabelId() *string {
	return s.LabelId
}

func (s *ListTrustDevicesRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListTrustDevicesRequest) GetUserCustomId() *string {
	return s.UserCustomId
}

func (s *ListTrustDevicesRequest) SetLabelContent(v string) *ListTrustDevicesRequest {
	s.LabelContent = &v
	return s
}

func (s *ListTrustDevicesRequest) SetLabelId(v string) *ListTrustDevicesRequest {
	s.LabelId = &v
	return s
}

func (s *ListTrustDevicesRequest) SetSerialNo(v string) *ListTrustDevicesRequest {
	s.SerialNo = &v
	return s
}

func (s *ListTrustDevicesRequest) SetUserCustomId(v string) *ListTrustDevicesRequest {
	s.UserCustomId = &v
	return s
}

func (s *ListTrustDevicesRequest) Validate() error {
	return dara.Validate(s)
}
