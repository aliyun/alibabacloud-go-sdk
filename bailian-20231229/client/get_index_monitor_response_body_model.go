// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIndexMonitorResponseBody
	GetCode() *string
	SetData(v interface{}) *GetIndexMonitorResponseBody
	GetData() interface{}
	SetMessage(v string) *GetIndexMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIndexMonitorResponseBody
	GetRequestId() *string
	SetStatus(v int32) *GetIndexMonitorResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *GetIndexMonitorResponseBody
	GetSuccess() *bool
}

type GetIndexMonitorResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//     "code": "Success",
	//
	//     "status_code": 200,
	//
	//     "data": {
	//
	//         "storageMonitorData": Object{...},
	//
	//         "qpsMonitorData": Object{...}
	//
	//     },
	//
	//     "success": true,
	//
	//     "message": "success",
	//
	//     "request_id": "65d34b79-b97e-478e-a0a3-xxx",
	//
	//     "status": "SUCCESS"
	//
	// }
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIndexMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIndexMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIndexMonitorResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetIndexMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIndexMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIndexMonitorResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetIndexMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIndexMonitorResponseBody) SetCode(v string) *GetIndexMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *GetIndexMonitorResponseBody) SetData(v interface{}) *GetIndexMonitorResponseBody {
	s.Data = v
	return s
}

func (s *GetIndexMonitorResponseBody) SetMessage(v string) *GetIndexMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *GetIndexMonitorResponseBody) SetRequestId(v string) *GetIndexMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexMonitorResponseBody) SetStatus(v int32) *GetIndexMonitorResponseBody {
	s.Status = &v
	return s
}

func (s *GetIndexMonitorResponseBody) SetSuccess(v bool) *GetIndexMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *GetIndexMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
