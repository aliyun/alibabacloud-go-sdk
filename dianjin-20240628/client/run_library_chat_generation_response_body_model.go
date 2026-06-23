// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLibraryChatGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *RunLibraryChatGenerationResponseBody
	GetCost() *int64
	SetData(v interface{}) *RunLibraryChatGenerationResponseBody
	GetData() interface{}
	SetDataType(v string) *RunLibraryChatGenerationResponseBody
	GetDataType() *string
	SetErrCode(v string) *RunLibraryChatGenerationResponseBody
	GetErrCode() *string
	SetMessage(v string) *RunLibraryChatGenerationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunLibraryChatGenerationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunLibraryChatGenerationResponseBody
	GetSuccess() *bool
	SetTime(v string) *RunLibraryChatGenerationResponseBody
	GetTime() *string
}

type RunLibraryChatGenerationResponseBody struct {
	// The time consumed.
	//
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// The response data (non-streaming).
	//
	// example:
	//
	// {
	//
	//     "finish":true,
	//
	//     "text":"是的，记名预付卡的有效期不得低于3年。",
	//
	//     "message": "是的，记名预付卡的有效期不得低于3年。",
	//
	//     "queryResult":{
	//
	//       "errCode": "0",	//接口协议层面的错误码，正常为0，未授权，参数错误，服务器异常时会发生变化，可以忽略不处理
	//
	//       "message": "ok",	//接口协议层面的消息
	//
	//       "data":{
	//
	//         "success": true,	//有回答true 无回答false
	//
	//         "answer": null,	//大模型生成的回答结果
	//
	//         "embeddingElapsedMs": 127,
	//
	//         "vectorSearchElapsedMs": 2745,
	//
	//         "llmElapsedMs": 7911,
	//
	//         "totalElapsedMs": 10820,
	//
	//         "chunks": [	//召回的分块信息，一般为top5或top10，可联系我们调整效果
	//
	//           {
	//
	//             "chunkId": "470182177892469799",	//分块信息的编号
	//
	//             "docId": "22666332",	//分块关联的文档编号
	//
	//             "chunkText": "Profits plummeted in the first quarter, can\\"t you bear it? In fact, previous rounds of price cuts have indeed had a certain impact on Tesla\\"s financial data. Tesla has just released its financial report for the first quarter of this year. The data shows that in Q1 2023, Tesla achieved revenue of 23.33 billion US dollars, an increase of 24% over the previous year; Tesla delivered more than 422,000 electric vehicles worldwide in the first quarter, an increase of 36% over the previous year",	//新闻原始内容
	//
	//             "chunkMeta": {	// demo数据中的其他metadata
	//
	//               "language": "en",
	//
	//               "unique_id": "news_22666332_13",
	//
	//               "content_type": "news",
	//
	//               "stock_id_list": []
	//
	//             }
	//
	//           }],
	//
	//         "documents": [{
	//
	//           "docId": "1686637056086872065",	//文档编号
	//
	//           "gmtCreate": "2023-08-02 15:16:25",	//文档的创建时间
	//
	//           "libraryId": "a1b2c3",	//文档关联的知识库编号
	//
	//           "title": "2023年工银信用卡微信、京东绑卡消费累计积分活动",	//文档标题
	//
	//           "url": null	//文档连接，如有
	//
	//         }]	//块文本关联的文档
	//
	//       },
	//
	//       "success": true	//接口协议层面的成功/失败状态 true就是errCode为0
	//
	//     }
	//
	//   }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// The data type.
	//
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RunLibraryChatGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *RunLibraryChatGenerationResponseBody) GetData() interface{} {
	return s.Data
}

func (s *RunLibraryChatGenerationResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *RunLibraryChatGenerationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RunLibraryChatGenerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunLibraryChatGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunLibraryChatGenerationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunLibraryChatGenerationResponseBody) GetTime() *string {
	return s.Time
}

func (s *RunLibraryChatGenerationResponseBody) SetCost(v int64) *RunLibraryChatGenerationResponseBody {
	s.Cost = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetData(v interface{}) *RunLibraryChatGenerationResponseBody {
	s.Data = v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetDataType(v string) *RunLibraryChatGenerationResponseBody {
	s.DataType = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetErrCode(v string) *RunLibraryChatGenerationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetMessage(v string) *RunLibraryChatGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetRequestId(v string) *RunLibraryChatGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetSuccess(v bool) *RunLibraryChatGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetTime(v string) *RunLibraryChatGenerationResponseBody {
	s.Time = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) Validate() error {
	return dara.Validate(s)
}
