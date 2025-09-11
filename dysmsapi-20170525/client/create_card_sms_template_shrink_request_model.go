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
	// The mobile phone manufacturer. Valid values:
	//
	// 	- **HuaWei**: HUAWEI
	//
	// 	- **XiaoMi**: Xiaomi
	//
	// 	- **OPPO**: OPPO
	//
	// 	- **VIVO**: vivo
	//
	// 	- **MEIZU**: MEIZU
	//
	// > If this parameter is not specified, the system automatically specifies a supported mobile phone manufacturer.
	//
	// example:
	//
	// XiaoMi
	Factorys *string `json:"Factorys,omitempty" xml:"Factorys,omitempty"`
	// The description of the message template.
	//
	// example:
	//
	// Image and Text Template
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// The content of the card message template.
	//
	// >
	//
	// 	- For information about fields such as Template, ExtendInfo, TemplateContent, TmpCard, and Action, see [Parameters of card message templates](https://help.aliyun.com/document_detail/434929.html).
	//
	// 	- Message template content varies based on the template type. For more information, see [Sample message templates](https://help.aliyun.com/document_detail/435361.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//        "extendInfo":{
	//
	//               "scene":"HMOVM",
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
	//                                                  "merchantName":"test-template",
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
	//                                           "content":"this is a test msg.",
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
	//                                           "content":"Promotional information",
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
	//                                           "content":"Promotional information,",
	//
	//                                           "actionType":"OPEN_BROWSER",
	//
	//                                           "action":{
	//
	//                                                  "target":"https://www.aliyun.com",
	//
	//                                                  "merchantName":"Currently on the Alibaba Cloud official website."
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
	//        "cardSignName":"aliyun",
	//
	//        "cardType":5
	//
	// }
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The name of the card message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun Image and Text Template
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
