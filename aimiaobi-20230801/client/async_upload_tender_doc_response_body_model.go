// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadTenderDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AsyncUploadTenderDocResponseBody
	GetCode() *string
	SetData(v *AsyncUploadTenderDocResponseBodyData) *AsyncUploadTenderDocResponseBody
	GetData() *AsyncUploadTenderDocResponseBodyData
	SetHttpStatusCode(v int32) *AsyncUploadTenderDocResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AsyncUploadTenderDocResponseBody
	GetMessage() *string
	SetRequestId(v string) *AsyncUploadTenderDocResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AsyncUploadTenderDocResponseBody
	GetSuccess() *bool
}

type AsyncUploadTenderDocResponseBody struct {
	// example:
	//
	// successful
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AsyncUploadTenderDocResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AsyncUploadTenderDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadTenderDocResponseBody) GoString() string {
	return s.String()
}

func (s *AsyncUploadTenderDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *AsyncUploadTenderDocResponseBody) GetData() *AsyncUploadTenderDocResponseBodyData {
	return s.Data
}

func (s *AsyncUploadTenderDocResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AsyncUploadTenderDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AsyncUploadTenderDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsyncUploadTenderDocResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AsyncUploadTenderDocResponseBody) SetCode(v string) *AsyncUploadTenderDocResponseBody {
	s.Code = &v
	return s
}

func (s *AsyncUploadTenderDocResponseBody) SetData(v *AsyncUploadTenderDocResponseBodyData) *AsyncUploadTenderDocResponseBody {
	s.Data = v
	return s
}

func (s *AsyncUploadTenderDocResponseBody) SetHttpStatusCode(v int32) *AsyncUploadTenderDocResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AsyncUploadTenderDocResponseBody) SetMessage(v string) *AsyncUploadTenderDocResponseBody {
	s.Message = &v
	return s
}

func (s *AsyncUploadTenderDocResponseBody) SetRequestId(v string) *AsyncUploadTenderDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsyncUploadTenderDocResponseBody) SetSuccess(v bool) *AsyncUploadTenderDocResponseBody {
	s.Success = &v
	return s
}

func (s *AsyncUploadTenderDocResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AsyncUploadTenderDocResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AsyncUploadTenderDocResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadTenderDocResponseBodyData) GoString() string {
	return s.String()
}

func (s *AsyncUploadTenderDocResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncUploadTenderDocResponseBodyData) SetTaskId(v string) *AsyncUploadTenderDocResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *AsyncUploadTenderDocResponseBodyData) Validate() error {
	return dara.Validate(s)
}
