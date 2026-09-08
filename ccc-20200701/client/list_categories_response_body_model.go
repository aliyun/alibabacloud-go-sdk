// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCategoriesResponseBody
	GetCode() *string
	SetData(v string) *ListCategoriesResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ListCategoriesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCategoriesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCategoriesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCategoriesResponseBody
	GetRequestId() *string
}

type ListCategoriesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data. The category node information, in the format of a JSON string.
	//
	// example:
	//
	// [{\\"categoryId\\":\\"43c2671b-8939-4223-****-6bd187905cc8\\",\\"childCategoryList\\":[{\\"categoryId\\":\\"120816ad-4392-4edf-****-6d053d5cfa5a\\",\\"childCategoryList\\":[],\\"deleted\\":0,\\"editor\\":\\"283277706217028904\\",\\"editorName\\":\\"283277706217028904\\",\\"instanceId\\":\\"cccV2-kmz\\",\\"itemCount\\":0,\\"level\\":2,\\"name\\":\\"客户反馈\\",\\"parentCategoryId\\":\\"43c2671b-8939-****-86d0-6bd187905cc8\\",\\"type\\":\\"Ticket\\"}],\\"deleted\\":0,\\"editor\\":\\"283277706217028904\\",\\"editorName\\":\\"283277706217028904\\",\\"instanceId\\":\\"cccV2-kmz\\",\\"itemCount\\":10,\\"level\\":1,\\"name\\":\\"测试一01类目\\",\\"type\\":\\"Ticket\\"},{\\"categoryId\\":\\"4948fcd0-2972-****-81c6-1a00927e1802\\",\\"childCategoryList\\":[],\\"deleted\\":0,\\"editor\\":\\"283277706217028904\\",\\"editorName\\":\\"283277706217028904\\",\\"instanceId\\":\\"cccV2-kmz\\",\\"itemCount\\":0,\\"level\\":1,\\"name\\":\\"生产环境验证\\",\\"type\\":\\"Ticket\\"},{\\"categoryId\\":\\"c426bd7f-9661-47c3-****-2508f1a32f66\\",\\"childCategoryList\\":[],\\"deleted\\":0,\\"editor\\":\\"269801834095770945\\",\\"editorName\\":\\"269801834095770945\\",\\"instanceId\\":\\"cccV2-kmz\\",\\"itemCount\\":4,\\"level\\":1,\\"name\\":\\"测试环境测试\\",\\"type\\":\\"Ticket\\"}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of incorrect parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DE803553-8AA9-4B9D-9E4E-A82BC69EDCEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCategoriesResponseBody) GetData() *string {
	return s.Data
}

func (s *ListCategoriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCategoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCategoriesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCategoriesResponseBody) SetCode(v string) *ListCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCategoriesResponseBody) SetData(v string) *ListCategoriesResponseBody {
	s.Data = &v
	return s
}

func (s *ListCategoriesResponseBody) SetHttpStatusCode(v int32) *ListCategoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCategoriesResponseBody) SetMessage(v string) *ListCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoriesResponseBody) SetParams(v []*string) *ListCategoriesResponseBody {
	s.Params = v
	return s
}

func (s *ListCategoriesResponseBody) SetRequestId(v string) *ListCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoriesResponseBody) Validate() error {
	return dara.Validate(s)
}
