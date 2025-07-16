// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetTicketResponseBody
	GetCreateTime() *string
	SetCreator(v *GetTicketResponseBodyCreator) *GetTicketResponseBody
	GetCreator() *GetTicketResponseBodyCreator
	SetCustomFields(v string) *GetTicketResponseBody
	GetCustomFields() *string
	SetOpenConversationId(v string) *GetTicketResponseBody
	GetOpenConversationId() *string
	SetOpenTicketId(v string) *GetTicketResponseBody
	GetOpenTicketId() *string
	SetProcessor(v *GetTicketResponseBodyProcessor) *GetTicketResponseBody
	GetProcessor() *GetTicketResponseBodyProcessor
	SetRequestId(v string) *GetTicketResponseBody
	GetRequestId() *string
	SetScene(v string) *GetTicketResponseBody
	GetScene() *string
	SetSceneContext(v string) *GetTicketResponseBody
	GetSceneContext() *string
	SetStage(v string) *GetTicketResponseBody
	GetStage() *string
	SetTakers(v []*GetTicketResponseBodyTakers) *GetTicketResponseBody
	GetTakers() []*GetTicketResponseBodyTakers
	SetTemplate(v *GetTicketResponseBodyTemplate) *GetTicketResponseBody
	GetTemplate() *GetTicketResponseBodyTemplate
	SetTitle(v string) *GetTicketResponseBody
	GetTitle() *string
	SetUpdateTime(v string) *GetTicketResponseBody
	GetUpdateTime() *string
	SetVendorRequestId(v string) *GetTicketResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetTicketResponseBody
	GetVendorType() *string
}

type GetTicketResponseBody struct {
	// example:
	//
	// 2021-07-09 14:45:01
	CreateTime *string                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator    *GetTicketResponseBodyCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// example:
	//
	// [{ "customerVisible": true,"editable": false,"identifier": "input1","name": "正文","placeholder": "请输入","required": false,"type": "TEXT_AREA","value": "123"},{"customerVisible": true,"editable": false,"identifier": "input2","name": "单选","options": [{"color": "#000000","value": "选项1"},{"color": "#000000","value": "选项2"},{"color": "#000000","value": "选项3"}],"required": false,"type": "RADIO"},{"customerVisible": true,"editable": false,"identifier": "input3","name": "多选","options": [{"color": "#000000","value": "选项1"},{"color": "#000000","value": "选项2"},{"color": "#000000","value": "选项3"}],"required": false,"type": "CHECKBOX"},{"customerVisible": true,"editable": false,"identifier": "input5","name": "数字","required": false,"type": "INPUT_NUMBER"},{"customerVisible": true,"editable": false,"identifier": "input6","maxFileNum": 20,"name": "上传","required": false,"type": "FILE"},{"customerVisible": true,"editable": false,"identifier": "input7","maxFileNum": 20,"name": "图片","required": false,"type": "IMAGE"},{"customerVisible": true,"editable": false,"format": "DATE_TIME","identifier": "input8","name": "日期","required": false,"type": "DATE_TIME_PICKER"}]
	CustomFields *string `json:"customFields,omitempty" xml:"customFields,omitempty"`
	// example:
	//
	// cidZBSNxxxxXUCrAA==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// a8iSxxxxtgiE
	OpenTicketId *string                         `json:"openTicketId,omitempty" xml:"openTicketId,omitempty"`
	Processor    *GetTicketResponseBodyProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// SG
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// {"groupId":1893227,"groupIsNormal":true,"groupName":"API和SPI测试群","groupSetId":29003,"groupSetName":"默认服务群组","scene":"SG"}
	SceneContext *string `json:"sceneContext,omitempty" xml:"sceneContext,omitempty"`
	// example:
	//
	// FINISHED
	Stage    *string                        `json:"stage,omitempty" xml:"stage,omitempty"`
	Takers   []*GetTicketResponseBodyTakers `json:"takers,omitempty" xml:"takers,omitempty" type:"Repeated"`
	Template *GetTicketResponseBodyTemplate `json:"template,omitempty" xml:"template,omitempty" type:"Struct"`
	// example:
	//
	// 贤文api测试，处理人王鸿程和李航宇
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2021-07-09 19:26:09
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTicketResponseBody) GetCreator() *GetTicketResponseBodyCreator {
	return s.Creator
}

func (s *GetTicketResponseBody) GetCustomFields() *string {
	return s.CustomFields
}

func (s *GetTicketResponseBody) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *GetTicketResponseBody) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *GetTicketResponseBody) GetProcessor() *GetTicketResponseBodyProcessor {
	return s.Processor
}

func (s *GetTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTicketResponseBody) GetScene() *string {
	return s.Scene
}

func (s *GetTicketResponseBody) GetSceneContext() *string {
	return s.SceneContext
}

func (s *GetTicketResponseBody) GetStage() *string {
	return s.Stage
}

func (s *GetTicketResponseBody) GetTakers() []*GetTicketResponseBodyTakers {
	return s.Takers
}

func (s *GetTicketResponseBody) GetTemplate() *GetTicketResponseBodyTemplate {
	return s.Template
}

func (s *GetTicketResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetTicketResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetTicketResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetTicketResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetTicketResponseBody) SetCreateTime(v string) *GetTicketResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTicketResponseBody) SetCreator(v *GetTicketResponseBodyCreator) *GetTicketResponseBody {
	s.Creator = v
	return s
}

func (s *GetTicketResponseBody) SetCustomFields(v string) *GetTicketResponseBody {
	s.CustomFields = &v
	return s
}

func (s *GetTicketResponseBody) SetOpenConversationId(v string) *GetTicketResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetTicketResponseBody) SetOpenTicketId(v string) *GetTicketResponseBody {
	s.OpenTicketId = &v
	return s
}

func (s *GetTicketResponseBody) SetProcessor(v *GetTicketResponseBodyProcessor) *GetTicketResponseBody {
	s.Processor = v
	return s
}

func (s *GetTicketResponseBody) SetRequestId(v string) *GetTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketResponseBody) SetScene(v string) *GetTicketResponseBody {
	s.Scene = &v
	return s
}

func (s *GetTicketResponseBody) SetSceneContext(v string) *GetTicketResponseBody {
	s.SceneContext = &v
	return s
}

func (s *GetTicketResponseBody) SetStage(v string) *GetTicketResponseBody {
	s.Stage = &v
	return s
}

func (s *GetTicketResponseBody) SetTakers(v []*GetTicketResponseBodyTakers) *GetTicketResponseBody {
	s.Takers = v
	return s
}

func (s *GetTicketResponseBody) SetTemplate(v *GetTicketResponseBodyTemplate) *GetTicketResponseBody {
	s.Template = v
	return s
}

func (s *GetTicketResponseBody) SetTitle(v string) *GetTicketResponseBody {
	s.Title = &v
	return s
}

func (s *GetTicketResponseBody) SetUpdateTime(v string) *GetTicketResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetTicketResponseBody) SetVendorRequestId(v string) *GetTicketResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetTicketResponseBody) SetVendorType(v string) *GetTicketResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTicketResponseBodyCreator struct {
	// example:
	//
	// 贤文
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 012345
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetTicketResponseBodyCreator) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyCreator) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyCreator) GetNickName() *string {
	return s.NickName
}

func (s *GetTicketResponseBodyCreator) GetUnionId() *string {
	return s.UnionId
}

func (s *GetTicketResponseBodyCreator) SetNickName(v string) *GetTicketResponseBodyCreator {
	s.NickName = &v
	return s
}

func (s *GetTicketResponseBodyCreator) SetUnionId(v string) *GetTicketResponseBodyCreator {
	s.UnionId = &v
	return s
}

func (s *GetTicketResponseBodyCreator) Validate() error {
	return dara.Validate(s)
}

type GetTicketResponseBodyProcessor struct {
	// example:
	//
	// 贤文
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 012345
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetTicketResponseBodyProcessor) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyProcessor) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyProcessor) GetNickName() *string {
	return s.NickName
}

func (s *GetTicketResponseBodyProcessor) GetUnionId() *string {
	return s.UnionId
}

func (s *GetTicketResponseBodyProcessor) SetNickName(v string) *GetTicketResponseBodyProcessor {
	s.NickName = &v
	return s
}

func (s *GetTicketResponseBodyProcessor) SetUnionId(v string) *GetTicketResponseBodyProcessor {
	s.UnionId = &v
	return s
}

func (s *GetTicketResponseBodyProcessor) Validate() error {
	return dara.Validate(s)
}

type GetTicketResponseBodyTakers struct {
	// example:
	//
	// 贤文
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 012345
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetTicketResponseBodyTakers) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyTakers) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyTakers) GetNickName() *string {
	return s.NickName
}

func (s *GetTicketResponseBodyTakers) GetUnionId() *string {
	return s.UnionId
}

func (s *GetTicketResponseBodyTakers) SetNickName(v string) *GetTicketResponseBodyTakers {
	s.NickName = &v
	return s
}

func (s *GetTicketResponseBodyTakers) SetUnionId(v string) *GetTicketResponseBodyTakers {
	s.UnionId = &v
	return s
}

func (s *GetTicketResponseBodyTakers) Validate() error {
	return dara.Validate(s)
}

type GetTicketResponseBodyTemplate struct {
	// OpenTemplateBizId
	//
	// example:
	//
	// OpenTemplateBizId
	OpenTemplateBizId *string `json:"OpenTemplateBizId,omitempty" xml:"OpenTemplateBizId,omitempty"`
	// OpenTemplateBizId
	//
	// example:
	//
	// OpenTemplateBizId
	OpenTemplateId *string `json:"OpenTemplateId,omitempty" xml:"OpenTemplateId,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetTicketResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyTemplate) GetOpenTemplateBizId() *string {
	return s.OpenTemplateBizId
}

func (s *GetTicketResponseBodyTemplate) GetOpenTemplateId() *string {
	return s.OpenTemplateId
}

func (s *GetTicketResponseBodyTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetTicketResponseBodyTemplate) SetOpenTemplateBizId(v string) *GetTicketResponseBodyTemplate {
	s.OpenTemplateBizId = &v
	return s
}

func (s *GetTicketResponseBodyTemplate) SetOpenTemplateId(v string) *GetTicketResponseBodyTemplate {
	s.OpenTemplateId = &v
	return s
}

func (s *GetTicketResponseBodyTemplate) SetTemplateName(v string) *GetTicketResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetTicketResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
