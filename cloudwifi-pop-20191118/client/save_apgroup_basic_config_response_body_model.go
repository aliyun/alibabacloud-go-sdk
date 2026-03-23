// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupBasicConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SaveApgroupBasicConfigResponseBodyData) *SaveApgroupBasicConfigResponseBody
	GetData() *SaveApgroupBasicConfigResponseBodyData
	SetErrorCode(v int32) *SaveApgroupBasicConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SaveApgroupBasicConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SaveApgroupBasicConfigResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *SaveApgroupBasicConfigResponseBody
	GetRequestId() *string
}

type SaveApgroupBasicConfigResponseBody struct {
	Data         *SaveApgroupBasicConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *int32                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                                   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveApgroupBasicConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupBasicConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApgroupBasicConfigResponseBody) GetData() *SaveApgroupBasicConfigResponseBodyData {
	return s.Data
}

func (s *SaveApgroupBasicConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SaveApgroupBasicConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveApgroupBasicConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SaveApgroupBasicConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveApgroupBasicConfigResponseBody) SetData(v *SaveApgroupBasicConfigResponseBodyData) *SaveApgroupBasicConfigResponseBody {
	s.Data = v
	return s
}

func (s *SaveApgroupBasicConfigResponseBody) SetErrorCode(v int32) *SaveApgroupBasicConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveApgroupBasicConfigResponseBody) SetErrorMessage(v string) *SaveApgroupBasicConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveApgroupBasicConfigResponseBody) SetIsSuccess(v bool) *SaveApgroupBasicConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SaveApgroupBasicConfigResponseBody) SetRequestId(v string) *SaveApgroupBasicConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveApgroupBasicConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveApgroupBasicConfigResponseBodyData struct {
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SaveApgroupBasicConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupBasicConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveApgroupBasicConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *SaveApgroupBasicConfigResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SaveApgroupBasicConfigResponseBodyData) SetId(v int64) *SaveApgroupBasicConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *SaveApgroupBasicConfigResponseBodyData) SetTaskId(v string) *SaveApgroupBasicConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SaveApgroupBasicConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
