// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOutputFileToResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveOutputFileToResourceResponseBody
	GetCode() *string
	SetMessage(v string) *SaveOutputFileToResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveOutputFileToResourceResponseBody
	GetRequestId() *string
	SetResults(v []*SaveOutputFileToResourceResponseBodyResults) *SaveOutputFileToResourceResponseBody
	GetResults() []*SaveOutputFileToResourceResponseBodyResults
}

type SaveOutputFileToResourceResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Results   []*SaveOutputFileToResourceResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s SaveOutputFileToResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveOutputFileToResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveOutputFileToResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveOutputFileToResourceResponseBody) GetResults() []*SaveOutputFileToResourceResponseBodyResults {
	return s.Results
}

func (s *SaveOutputFileToResourceResponseBody) SetCode(v string) *SaveOutputFileToResourceResponseBody {
	s.Code = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBody) SetMessage(v string) *SaveOutputFileToResourceResponseBody {
	s.Message = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBody) SetRequestId(v string) *SaveOutputFileToResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBody) SetResults(v []*SaveOutputFileToResourceResponseBodyResults) *SaveOutputFileToResourceResponseBody {
	s.Results = v
	return s
}

func (s *SaveOutputFileToResourceResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveOutputFileToResourceResponseBodyResults struct {
	// 失败时返回业务错误码（i18n key）
	//
	// example:
	//
	// string_value
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 失败时返回错误描述（已按请求 locale 国际化）
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 产出明细 ID
	//
	// example:
	//
	// exampleItemId
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 成功时返回新建的资源 sourceId
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 操作是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveOutputFileToResourceResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceResponseBodyResults) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetItemId() *string {
	return s.ItemId
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetSourceId() *string {
	return s.SourceId
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetSuccess() *bool {
	return s.Success
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetErrorCode(v string) *SaveOutputFileToResourceResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetErrorMessage(v string) *SaveOutputFileToResourceResponseBodyResults {
	s.ErrorMessage = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetItemId(v string) *SaveOutputFileToResourceResponseBodyResults {
	s.ItemId = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetSourceId(v string) *SaveOutputFileToResourceResponseBodyResults {
	s.SourceId = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetSuccess(v bool) *SaveOutputFileToResourceResponseBodyResults {
	s.Success = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
