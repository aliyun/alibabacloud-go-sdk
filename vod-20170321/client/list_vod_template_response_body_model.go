// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVodTemplateResponseBody
	GetRequestId() *string
	SetVodTemplateInfoList(v []*ListVodTemplateResponseBodyVodTemplateInfoList) *ListVodTemplateResponseBody
	GetVodTemplateInfoList() []*ListVodTemplateResponseBodyVodTemplateInfoList
}

type ListVodTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2A56B75B-B7E6-48*****27-A9BEAA3E50A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of snapshot template data.
	VodTemplateInfoList []*ListVodTemplateResponseBodyVodTemplateInfoList `json:"VodTemplateInfoList,omitempty" xml:"VodTemplateInfoList,omitempty" type:"Repeated"`
}

func (s ListVodTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodTemplateResponseBody) GetVodTemplateInfoList() []*ListVodTemplateResponseBodyVodTemplateInfoList {
	return s.VodTemplateInfoList
}

func (s *ListVodTemplateResponseBody) SetRequestId(v string) *ListVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodTemplateResponseBody) SetVodTemplateInfoList(v []*ListVodTemplateResponseBodyVodTemplateInfoList) *ListVodTemplateResponseBody {
	s.VodTemplateInfoList = v
	return s
}

func (s *ListVodTemplateResponseBody) Validate() error {
	if s.VodTemplateInfoList != nil {
		for _, item := range s.VodTemplateInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVodTemplateResponseBodyVodTemplateInfoList struct {
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the template was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-11-30T08:05:59:57Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the template is the default template. Valid values:
	//
	// - **Default**: The template is the default template.
	//
	// - **NotDefault**: The template is not the default template.
	//
	// example:
	//
	// NotDefault
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The time when the template was last modified. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2018-11-30T09:05:59:97Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The template name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The detailed template configuration in JSON format. For more information about the data structure, see [SnapshotTemplateConfig](https://help.aliyun.com/document_detail/98618.html).
	//
	// example:
	//
	// {\\"SnapshotConfig\\":{\\"Count\\":10,\\"SpecifiedOffsetTime\\":0,\\"Interval\\":1},\\"SnapshotType\\":\\"NormalSnapshot\\"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The templatetype. Valid values:
	//
	// - **Snapshot**: snapshot.
	//
	// - **DynamicImage**: animated image.
	//
	// example:
	//
	// Snapshot
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 7c49f2f42b1c*****0969fcd446690
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s ListVodTemplateResponseBodyVodTemplateInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListVodTemplateResponseBodyVodTemplateInfoList) GoString() string {
	return s.String()
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) GetAppId() *string {
	return s.AppId
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) GetName() *string {
	return s.Name
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetAppId(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.AppId = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetCreationTime(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetIsDefault(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.IsDefault = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetModifyTime(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.ModifyTime = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetName(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.Name = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetTemplateConfig(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.TemplateConfig = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetTemplateType(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.TemplateType = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetVodTemplateId(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.VodTemplateId = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) Validate() error {
	return dara.Validate(s)
}
