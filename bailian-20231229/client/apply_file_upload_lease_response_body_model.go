// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFileUploadLeaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyFileUploadLeaseResponseBody
	GetCode() *string
	SetData(v *ApplyFileUploadLeaseResponseBodyData) *ApplyFileUploadLeaseResponseBody
	GetData() *ApplyFileUploadLeaseResponseBodyData
	SetMessage(v string) *ApplyFileUploadLeaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyFileUploadLeaseResponseBody
	GetRequestId() *string
	SetStatus(v string) *ApplyFileUploadLeaseResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ApplyFileUploadLeaseResponseBody
	GetSuccess() *bool
}

type ApplyFileUploadLeaseResponseBody struct {
	// The status code.
	//
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data fields.
	Data *ApplyFileUploadLeaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyFileUploadLeaseResponseBody) GetData() *ApplyFileUploadLeaseResponseBodyData {
	return s.Data
}

func (s *ApplyFileUploadLeaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyFileUploadLeaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyFileUploadLeaseResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ApplyFileUploadLeaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyFileUploadLeaseResponseBody) SetCode(v string) *ApplyFileUploadLeaseResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetData(v *ApplyFileUploadLeaseResponseBodyData) *ApplyFileUploadLeaseResponseBody {
	s.Data = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetMessage(v string) *ApplyFileUploadLeaseResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetRequestId(v string) *ApplyFileUploadLeaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetStatus(v string) *ApplyFileUploadLeaseResponseBody {
	s.Status = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetSuccess(v bool) *ApplyFileUploadLeaseResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyFileUploadLeaseResponseBodyData struct {
	// The unique ID of the lease.
	//
	// example:
	//
	// 1e6a159107384782be5e45ac4759b247.1719325231035
	FileUploadLeaseId *string `json:"FileUploadLeaseId,omitempty" xml:"FileUploadLeaseId,omitempty"`
	// The HTTP request parameters used to upload the document.
	Param *ApplyFileUploadLeaseResponseBodyDataParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// The upload method of the document. Valid values:
	//
	// 	- OSS.PreSignedURL
	//
	// 	- HTTP
	//
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBodyData) GetFileUploadLeaseId() *string {
	return s.FileUploadLeaseId
}

func (s *ApplyFileUploadLeaseResponseBodyData) GetParam() *ApplyFileUploadLeaseResponseBodyDataParam {
	return s.Param
}

func (s *ApplyFileUploadLeaseResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetFileUploadLeaseId(v string) *ApplyFileUploadLeaseResponseBodyData {
	s.FileUploadLeaseId = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetParam(v *ApplyFileUploadLeaseResponseBodyDataParam) *ApplyFileUploadLeaseResponseBodyData {
	s.Param = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetType(v string) *ApplyFileUploadLeaseResponseBodyData {
	s.Type = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyFileUploadLeaseResponseBodyDataParam struct {
	// The key-value pair to be placed in the Header. Both the key and the value are strings.
	//
	// example:
	//
	// "X-bailian-extra": "MTAwNTQyNjQ5NTE2OTE3OA==",
	//
	//         "Content-Type": "application/pdf"
	Headers interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The HTTP call method. Valid values:
	//
	// 	- PUT
	//
	// 	- POST
	//
	// example:
	//
	// PUT
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The upload URL of the document.
	//
	// example:
	//
	// https://bailian-datahub-data-origin-prod.oss-cn-hangzhou.aliyuncs.com/1005426495169178/10024405/68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847.pdf?Expires=1716699536&OSSAccessKeyId=TestID&Signature=HfwPUZo4pR6DatSDym0zFKVh9Wg%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBodyDataParam) String() string {
	return dara.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBodyDataParam) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) GetHeaders() interface{} {
	return s.Headers
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) GetMethod() *string {
	return s.Method
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) GetUrl() *string {
	return s.Url
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetHeaders(v interface{}) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Headers = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetMethod(v string) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Method = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetUrl(v string) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Url = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) Validate() error {
	return dara.Validate(s)
}
