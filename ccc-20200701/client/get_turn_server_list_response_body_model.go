// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTurnServerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTurnServerListResponseBody
	GetCode() *string
	SetData(v string) *GetTurnServerListResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetTurnServerListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTurnServerListResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetTurnServerListResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetTurnServerListResponseBody
	GetRequestId() *string
}

type GetTurnServerListResponseBody struct {
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A list of front-end access point servers.
	//
	// example:
	//
	// [
	//
	// 	{
	//
	// 		"region":"hangzhou",
	//
	// 		"name":"杭州",
	//
	// 		"domain":"turn-hz-ecs.ccc.aliyuncs.com",
	//
	// 		"cidr":"172.31.XX.XX/28"
	//
	// 	},
	//
	// 	{
	//
	// 		"region":"qingdao",
	//
	// 		"name":"青岛",
	//
	// 		"domain":"turn-qd-ecs.ccc.aliyuncs.com",
	//
	// 		"cidr":"172.31.XX.XX/28"
	//
	// 	},
	//
	// 	{
	//
	// 		"region":"shanghai",
	//
	// 		"name":"上海",
	//
	// 		"domain":"turn-sh-ecs.ccc.aliyuncs.com",
	//
	// 		"cidr":"172.31.XX.XX/28"
	//
	// 	},
	//
	// 	{
	//
	// 		"region":"chengdu",
	//
	// 		"name":"成都",
	//
	// 		"domain":"turn-cd-ecs.ccc.aliyuncs.com",
	//
	// 		"cidr":"172.31.XX.XX/28"
	//
	// 	},
	//
	// 	{
	//
	// 		"region":"beijing",
	//
	// 		"name":"北京",
	//
	// 		"domain":"turn-bj-ecs.ccc.aliyuncs.com",
	//
	// 		"cidr":"172.31.XX.XX/28"
	//
	// 	},
	//
	// 	{
	//
	// 		"region":"huanan",
	//
	// 		"name":"深圳",
	//
	// 		"domain":"turn-sz-ecs.ccc.aliyuncs.com",
	//
	// 		"cidr":"172.31.XX.XX/28"
	//
	// 	},
	//
	// 	{
	//
	// 		"region":"zhangbei",
	//
	// 		"name":"张北",
	//
	// 		"domain":"turn-zb-ecs.ccc.aliyuncs.com",
	//
	// 		"cidr":"172.31.XX.XX/28"
	//
	// 	}
	//
	// ]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTurnServerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTurnServerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTurnServerListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTurnServerListResponseBody) GetData() *string {
	return s.Data
}

func (s *GetTurnServerListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTurnServerListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTurnServerListResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetTurnServerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTurnServerListResponseBody) SetCode(v string) *GetTurnServerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetData(v string) *GetTurnServerListResponseBody {
	s.Data = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetHttpStatusCode(v int32) *GetTurnServerListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetMessage(v string) *GetTurnServerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetParams(v []*string) *GetTurnServerListResponseBody {
	s.Params = v
	return s
}

func (s *GetTurnServerListResponseBody) SetRequestId(v string) *GetTurnServerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTurnServerListResponseBody) Validate() error {
	return dara.Validate(s)
}
