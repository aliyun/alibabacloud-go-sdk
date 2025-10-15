// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetParseResultResponseBody
	GetCost() *int64
	SetData(v *GetParseResultResponseBodyData) *GetParseResultResponseBody
	GetData() *GetParseResultResponseBodyData
	SetDataType(v string) *GetParseResultResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetParseResultResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetParseResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetParseResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetParseResultResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetParseResultResponseBody
	GetTime() *string
}

type GetParseResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetParseResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0abb793617204049360065953ec6dd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetParseResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParseResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetParseResultResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetParseResultResponseBody) GetData() *GetParseResultResponseBodyData {
	return s.Data
}

func (s *GetParseResultResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetParseResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetParseResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetParseResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParseResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetParseResultResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetParseResultResponseBody) SetCost(v int64) *GetParseResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetParseResultResponseBody) SetData(v *GetParseResultResponseBodyData) *GetParseResultResponseBody {
	s.Data = v
	return s
}

func (s *GetParseResultResponseBody) SetDataType(v string) *GetParseResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetParseResultResponseBody) SetErrCode(v string) *GetParseResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetParseResultResponseBody) SetMessage(v string) *GetParseResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetParseResultResponseBody) SetRequestId(v string) *GetParseResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParseResultResponseBody) SetSuccess(v bool) *GetParseResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetParseResultResponseBody) SetTime(v string) *GetParseResultResponseBody {
	s.Time = &v
	return s
}

func (s *GetParseResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetParseResultResponseBodyData struct {
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// null
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
	// example:
	//
	// b0a202e2-5031-4589-a6d7-39185f0d8d01
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {
	//
	//           "Status": "Success",
	//
	//           "Data": {},
	//
	//           "Message": null,
	//
	//           "TaskId": "docmind-20240601-123abc"
	//
	//         }
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	ResultUrl *string                `json:"resultUrl,omitempty" xml:"resultUrl,omitempty"`
	// example:
	//
	// WaitRefresh
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetParseResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetParseResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetParseResultResponseBodyData) GetFileType() *string {
	return s.FileType
}

func (s *GetParseResultResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *GetParseResultResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParseResultResponseBodyData) GetResult() map[string]interface{} {
	return s.Result
}

func (s *GetParseResultResponseBodyData) GetResultUrl() *string {
	return s.ResultUrl
}

func (s *GetParseResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetParseResultResponseBodyData) SetFileType(v string) *GetParseResultResponseBodyData {
	s.FileType = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetProviderType(v string) *GetParseResultResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetRequestId(v string) *GetParseResultResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetResult(v map[string]interface{}) *GetParseResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetParseResultResponseBodyData) SetResultUrl(v string) *GetParseResultResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetStatus(v string) *GetParseResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetParseResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
