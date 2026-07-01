// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCardSmsTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFactorys(v string) *CreateCardSmsTemplateShrinkRequest
	GetFactorys() *string
	SetMemo(v string) *CreateCardSmsTemplateShrinkRequest
	GetMemo() *string
	SetTemplateShrink(v string) *CreateCardSmsTemplateShrinkRequest
	GetTemplateShrink() *string
	SetTemplateName(v string) *CreateCardSmsTemplateShrinkRequest
	GetTemplateName() *string
}

type CreateCardSmsTemplateShrinkRequest struct {
	// The vendors to which the template will be submitted. Valid values:
	//
	// - **HuaWei**: Huawei
	//
	// - **XiaoMi**: Xiaomi
	//
	// - **OPPO**: OPPO
	//
	// - **VIVO**: VIVO
	//
	// - **MEIZU**: MEIZU
	//
	// - **HONOR**: HONOR
	//
	// > If you do not specify this parameter, the system automatically submits the template to all supported mobile phone vendors.
	//
	// example:
	//
	// ["HuaWei","XiaoMi"]
	Factorys *string `json:"Factorys,omitempty" xml:"Factorys,omitempty"`
	// A description of the template.
	//
	// example:
	//
	// 图文类模板
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// The content of the card SMS template.
	//
	// > - For more information about the `Template`, `ExtendInfo`, `TemplateContent`, `TmpCard`, and `Action` fields, see [Card SMS template parameters](https://help.aliyun.com/document_detail/434929.html).
	//
	// >
	//
	// > - The content structure varies based on the type of card SMS template. For more information, see [Card SMS template examples](https://help.aliyun.com/document_detail/435361.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//        "extendInfo":{
	//
	//               "scene":"HMOVM图文",
	//
	//               "purpose":"2",
	//
	//               "userExt":{
	//
	//                      "outId":"1234554321"
	//
	//               }
	//
	//        },
	//
	//        "templateContent":{
	//
	//               "pages":[
	//
	//                      {
	//
	// "tmpCards":[
	//
	//                                    {
	//
	//                                           "type":"IMAGE",
	//
	//                                           "srcType":1,
	//
	//                                           "src":"28755",
	//
	//                                           "actionType":"OPEN_APP",
	//
	//                                           "action":{
	//
	//                                                  "target":"https://s.tb.cn/c.KxzZ",
	//
	//                                                  "merchantName":"测试-图文模板",
	//
	//                                                  "packageName":[
	//
	//                                                         "com.taobao.taobao"],
	//
	//                                                  "floorUrl":"https://s.tb.cn/c.KxzZ"
	//
	//                                           },
	//
	//                                           "positionNumber":1
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"TEXT",
	//
	//                                           "content":"测试- BENZ AMG 2020 试驾邀请",
	//
	//                                           "isTextTitle":true,
	//
	//                                           "positionNumber":2
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"TEXT",
	//
	//                                           "content":"测试-梅赛德斯-奔驰，创新激情永不灭。作为汽车 XXXX",
	//
	//                                           "isTextTitle":false,
	//
	//                                           "positionNumber":3
	//
	//                                    },
	//
	//                                    {
	//
	//                                           "type":"BUTTON",
	//
	//                                           "content":"预约试驾",
	//
	//                                           "actionType":"OPEN_BROWSER",
	//
	//                                           "action":{
	//
	//                                                  "target":"https://www.mercedes-benz.com.cn",
	//
	//                                                  "merchantName":"测试-正在跳转梅赛德斯-奔驰"
	//
	// },
	//
	//                                           "positionNumber":4
	//
	//                                    }]
	//
	//                      }]
	//
	//        },
	//
	//        "cardSignName":"阿里云",
	//
	//        "cardType":5,
	//
	//        "companyName": "投放企业名称"
	//
	// }
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the card SMS template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云图文类模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateCardSmsTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCardSmsTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCardSmsTemplateShrinkRequest) GetFactorys() *string {
	return s.Factorys
}

func (s *CreateCardSmsTemplateShrinkRequest) GetMemo() *string {
	return s.Memo
}

func (s *CreateCardSmsTemplateShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *CreateCardSmsTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateCardSmsTemplateShrinkRequest) SetFactorys(v string) *CreateCardSmsTemplateShrinkRequest {
	s.Factorys = &v
	return s
}

func (s *CreateCardSmsTemplateShrinkRequest) SetMemo(v string) *CreateCardSmsTemplateShrinkRequest {
	s.Memo = &v
	return s
}

func (s *CreateCardSmsTemplateShrinkRequest) SetTemplateShrink(v string) *CreateCardSmsTemplateShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *CreateCardSmsTemplateShrinkRequest) SetTemplateName(v string) *CreateCardSmsTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateCardSmsTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
