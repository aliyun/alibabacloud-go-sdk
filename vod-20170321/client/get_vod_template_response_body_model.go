// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetVodTemplateResponseBody
	GetRequestId() *string
	SetVodTemplateInfo(v *GetVodTemplateResponseBodyVodTemplateInfo) *GetVodTemplateResponseBody
	GetVodTemplateInfo() *GetVodTemplateResponseBodyVodTemplateInfo
}

type GetVodTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DE7A1F49-41C1-47*****DF-4CD0C02087DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the snapshot template.
	VodTemplateInfo *GetVodTemplateResponseBodyVodTemplateInfo `json:"VodTemplateInfo,omitempty" xml:"VodTemplateInfo,omitempty" type:"Struct"`
}

func (s GetVodTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetVodTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVodTemplateResponseBody) GetVodTemplateInfo() *GetVodTemplateResponseBodyVodTemplateInfo {
	return s.VodTemplateInfo
}

func (s *GetVodTemplateResponseBody) SetRequestId(v string) *GetVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVodTemplateResponseBody) SetVodTemplateInfo(v *GetVodTemplateResponseBodyVodTemplateInfo) *GetVodTemplateResponseBody {
	s.VodTemplateInfo = v
	return s
}

func (s *GetVodTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVodTemplateResponseBodyVodTemplateInfo struct {
	// The time when the template was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-30T08:05:59:57Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the template is the default one. Valid values:
	//
	// 	- **Default**: The template is the default one.
	//
	// 	- **NotDefault**: The template is not the default one.
	//
	// example:
	//
	// NotDefault
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The time when the template was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-30T09:05:59:57Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The detailed configurations of the template. The value is a JSON-formatted string. For more information about the data structure, see the "SnapshotTemplateConfig" section of the [Media processing parameters](https://help.aliyun.com/document_detail/98618.html) topic.
	//
	// example:
	//
	// {\\"SnapshotConfig\\":{\\"Count\\":10,\\"SpecifiedOffsetTime\\":0,\\"Interval\\":1},\\"SnapshotType\\":\\"NormalSnapshot\\"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The type of the template. Valid values:
	//
	// 	- **Snapshot**
	//
	// 	- **DynamicImage**
	//
	// example:
	//
	// Snapshot
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 7c49f2f4c09*****69fcd446690
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s GetVodTemplateResponseBodyVodTemplateInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVodTemplateResponseBodyVodTemplateInfo) GoString() string {
	return s.String()
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) GetName() *string {
	return s.Name
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) GetVodTemplateId() *string {
	return s.VodTemplateId
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetCreationTime(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.CreationTime = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetIsDefault(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.IsDefault = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetModifyTime(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetName(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.Name = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetTemplateConfig(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.TemplateConfig = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetTemplateType(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.TemplateType = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetVodTemplateId(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.VodTemplateId = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) Validate() error {
	return dara.Validate(s)
}
