// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTempStorageLeaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyTempStorageLeaseResponseBody
	GetCode() *string
	SetData(v *ApplyTempStorageLeaseResponseBodyData) *ApplyTempStorageLeaseResponseBody
	GetData() *ApplyTempStorageLeaseResponseBodyData
	SetMessage(v string) *ApplyTempStorageLeaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyTempStorageLeaseResponseBody
	GetRequestId() *string
	SetStatus(v string) *ApplyTempStorageLeaseResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ApplyTempStorageLeaseResponseBody
	GetSuccess() *bool
}

type ApplyTempStorageLeaseResponseBody struct {
	// The status code.
	//
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data fields.
	Data *ApplyTempStorageLeaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyTempStorageLeaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyTempStorageLeaseResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyTempStorageLeaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyTempStorageLeaseResponseBody) GetData() *ApplyTempStorageLeaseResponseBodyData {
	return s.Data
}

func (s *ApplyTempStorageLeaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyTempStorageLeaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyTempStorageLeaseResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ApplyTempStorageLeaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyTempStorageLeaseResponseBody) SetCode(v string) *ApplyTempStorageLeaseResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyTempStorageLeaseResponseBody) SetData(v *ApplyTempStorageLeaseResponseBodyData) *ApplyTempStorageLeaseResponseBody {
	s.Data = v
	return s
}

func (s *ApplyTempStorageLeaseResponseBody) SetMessage(v string) *ApplyTempStorageLeaseResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyTempStorageLeaseResponseBody) SetRequestId(v string) *ApplyTempStorageLeaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyTempStorageLeaseResponseBody) SetStatus(v string) *ApplyTempStorageLeaseResponseBody {
	s.Status = &v
	return s
}

func (s *ApplyTempStorageLeaseResponseBody) SetSuccess(v bool) *ApplyTempStorageLeaseResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyTempStorageLeaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyTempStorageLeaseResponseBodyData struct {
	// HTTP parameters used for uploading the file.
	Param *ApplyTempStorageLeaseResponseBodyDataParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// Unique lease ID. This parameter is required when retrieving the uploaded file within the application later.
	//
	// example:
	//
	// 1e6a159107384782be5e45ac4759b247.1719325231035
	TempStorageLeaseId *string `json:"TempStorageLeaseId,omitempty" xml:"TempStorageLeaseId,omitempty"`
}

func (s ApplyTempStorageLeaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ApplyTempStorageLeaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyTempStorageLeaseResponseBodyData) GetParam() *ApplyTempStorageLeaseResponseBodyDataParam {
	return s.Param
}

func (s *ApplyTempStorageLeaseResponseBodyData) GetTempStorageLeaseId() *string {
	return s.TempStorageLeaseId
}

func (s *ApplyTempStorageLeaseResponseBodyData) SetParam(v *ApplyTempStorageLeaseResponseBodyDataParam) *ApplyTempStorageLeaseResponseBodyData {
	s.Param = v
	return s
}

func (s *ApplyTempStorageLeaseResponseBodyData) SetTempStorageLeaseId(v string) *ApplyTempStorageLeaseResponseBodyData {
	s.TempStorageLeaseId = &v
	return s
}

func (s *ApplyTempStorageLeaseResponseBodyData) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyTempStorageLeaseResponseBodyDataParam struct {
	// K-V fields to be included in the Header; both Key and Value are strings.
	//
	// example:
	//
	// Content-Type: application/json
	Headers interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// HTTP method for the call. Valid values: PUT POST
	//
	// example:
	//
	// PUT
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// Authorized URL for the file upload.
	//
	// example:
	//
	// https://bailian-datahub-data-origin-prod.oss-cn-hangzhou.aliyuncs.com/1005426495169178/10024405/68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847.pdf?Expires=1716699536&OSSAccessKeyId=TestID&Signature=HfwPUZo4pR6DatSDym0zFKVh9Wg%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ApplyTempStorageLeaseResponseBodyDataParam) String() string {
	return dara.Prettify(s)
}

func (s ApplyTempStorageLeaseResponseBodyDataParam) GoString() string {
	return s.String()
}

func (s *ApplyTempStorageLeaseResponseBodyDataParam) GetHeaders() interface{} {
	return s.Headers
}

func (s *ApplyTempStorageLeaseResponseBodyDataParam) GetMethod() *string {
	return s.Method
}

func (s *ApplyTempStorageLeaseResponseBodyDataParam) GetUrl() *string {
	return s.Url
}

func (s *ApplyTempStorageLeaseResponseBodyDataParam) SetHeaders(v interface{}) *ApplyTempStorageLeaseResponseBodyDataParam {
	s.Headers = v
	return s
}

func (s *ApplyTempStorageLeaseResponseBodyDataParam) SetMethod(v string) *ApplyTempStorageLeaseResponseBodyDataParam {
	s.Method = &v
	return s
}

func (s *ApplyTempStorageLeaseResponseBodyDataParam) SetUrl(v string) *ApplyTempStorageLeaseResponseBodyDataParam {
	s.Url = &v
	return s
}

func (s *ApplyTempStorageLeaseResponseBodyDataParam) Validate() error {
	return dara.Validate(s)
}
