// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalSearchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ListMultimodalSearchTaskResponseBody
	GetData() interface{}
	SetErrCode(v string) *ListMultimodalSearchTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMultimodalSearchTaskResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ListMultimodalSearchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMultimodalSearchTaskResponseBody
	GetSuccess() *bool
}

type ListMultimodalSearchTaskResponseBody struct {
	// example:
	//
	// 模型相关信息
	//
	//
	// {
	//
	//   "Data": {
	//
	//     "PageSize": 1,
	//
	//     "PageNumber": 1,
	//
	//     "Total": 12,
	//
	//     "Items": [
	//
	//       {
	//
	//         "Status": "Running",
	//
	//         "ProcessBatch": 0,
	//
	//         "TotalBatch": 0,
	//
	//         "EndTime": "",
	//
	//         "CreateTime": "2025-12-25T20:38:35",
	//
	//         "Id": "dd84065b-eb46-4c7b-abb0-0d633ec811d7",
	//
	//         "SuccessBatch": 0,
	//
	//         "Query" : "油漆桶",
	//
	//         "DatasetIds": ["ds-1", "ds-2"]
	//
	//       }
	//
	//     ]
	//
	//   },
	//
	//   "RequestId": "849E84DB-*****-9D734",
	//
	//   "Success": true
	//
	// }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMultimodalSearchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalSearchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultimodalSearchTaskResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ListMultimodalSearchTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMultimodalSearchTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMultimodalSearchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultimodalSearchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMultimodalSearchTaskResponseBody) SetData(v interface{}) *ListMultimodalSearchTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListMultimodalSearchTaskResponseBody) SetErrCode(v string) *ListMultimodalSearchTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMultimodalSearchTaskResponseBody) SetErrMessage(v string) *ListMultimodalSearchTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMultimodalSearchTaskResponseBody) SetRequestId(v string) *ListMultimodalSearchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultimodalSearchTaskResponseBody) SetSuccess(v bool) *ListMultimodalSearchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ListMultimodalSearchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
