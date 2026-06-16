// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAsyncTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAsyncTaskResultResponseBody
	GetCode() *string
	SetData(v *QueryAsyncTaskResultResponseBodyData) *QueryAsyncTaskResultResponseBody
	GetData() *QueryAsyncTaskResultResponseBodyData
	SetMessage(v string) *QueryAsyncTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAsyncTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAsyncTaskResultResponseBody
	GetSuccess() *bool
}

type QueryAsyncTaskResultResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asynchronous task result.
	Data *QueryAsyncTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAsyncTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAsyncTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAsyncTaskResultResponseBody) GetData() *QueryAsyncTaskResultResponseBodyData {
	return s.Data
}

func (s *QueryAsyncTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAsyncTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAsyncTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAsyncTaskResultResponseBody) SetCode(v string) *QueryAsyncTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAsyncTaskResultResponseBody) SetData(v *QueryAsyncTaskResultResponseBodyData) *QueryAsyncTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *QueryAsyncTaskResultResponseBody) SetMessage(v string) *QueryAsyncTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAsyncTaskResultResponseBody) SetRequestId(v string) *QueryAsyncTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAsyncTaskResultResponseBody) SetSuccess(v bool) *QueryAsyncTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAsyncTaskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAsyncTaskResultResponseBodyData struct {
	// The task result.
	//
	// example:
	//
	// {\\"url\\":\\"https://aidge-bailian-oss.oss-cn-beijing.aliyuncs.com/gaolinfeng/pdf_trans/translated_en.pdf?security-token=CAISzwJ1q6Ft5B2yfSjIr5ntKv7urOdn9YTeaVbb1lQRfcxi2Kz%2BgDz2IHhMeHFgAeAbs%2Fw%2Fm29W6v4SlqZdVplOWU3Da%2BB364xK7Q754wRDcULuv9I%2Bk5SANTW5KXyShb3%2FAYjQSNfaZY3eCTTtnTNyxr3XbCirW0ffX7SClZ9gaKZ8PGD6F00kYu1bPQx%2FssQXGGLMPPK2SH7Qj3HXEVBjt3gX6wo9y9zmk53FsUWA1QKmlr9F%2BdWhGPX%2BMZkwZqUYesyuwel7epDG1CNt8BVQ%2FM909vccoWuf7onNXgQJs0rZbbaMoscSJQ51aaV%2FFaUBt%2FXmi%2Fxzt6nJkID626jAvGbZzsW0rumBtyikcIvBXr5RHT3rIrVAU%2BuEf19557bo3dbfkNdWOrtHZDY5Qn9nURKxAbSEg2uBaWTIIJPmvc97r9wbhjuH87JeC0jQHt3xuRqAATp5EbLOeo%2BZktMLbi%2FUsZgcvdIEIv3tPBCYfyJnh%2Bj6U8IyaLKQYHQBtjsfyMngFfS09jFdjBcUebzvYJs21gyU5u%2FZ9SeReh%2FacuIMVoggWy3o9Y%2BnBA2QPCcKqVM7XlwriM%2FJOAyhvj%2Bjtj7BUoyD%2BrSgF5brq5ykjN7t2U7oIAA%3D&OSSAccessKeyId=STS.NZXaDZA8FBF5kpj2TDqCN7iUb&Expires=1780315869&Signature=CdbMkhcED4Ovhw438ZVe5nzU1mk%3D\\"}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The task status.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The downstream task ID.
	//
	// example:
	//
	// task-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage information.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s QueryAsyncTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAsyncTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResultResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *QueryAsyncTaskResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryAsyncTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAsyncTaskResultResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *QueryAsyncTaskResultResponseBodyData) SetResult(v string) *QueryAsyncTaskResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *QueryAsyncTaskResultResponseBodyData) SetStatus(v string) *QueryAsyncTaskResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryAsyncTaskResultResponseBodyData) SetTaskId(v string) *QueryAsyncTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryAsyncTaskResultResponseBodyData) SetUsageMap(v map[string]*int64) *QueryAsyncTaskResultResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *QueryAsyncTaskResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
