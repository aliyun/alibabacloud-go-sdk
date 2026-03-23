// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApgroupSsidConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteApgroupSsidConfigResponseBodyData) *DeleteApgroupSsidConfigResponseBody
	GetData() *DeleteApgroupSsidConfigResponseBodyData
	SetErrorCode(v int32) *DeleteApgroupSsidConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *DeleteApgroupSsidConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *DeleteApgroupSsidConfigResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteApgroupSsidConfigResponseBody
	GetRequestId() *string
}

type DeleteApgroupSsidConfigResponseBody struct {
	Data         *DeleteApgroupSsidConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *int32                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                                    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApgroupSsidConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupSsidConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApgroupSsidConfigResponseBody) GetData() *DeleteApgroupSsidConfigResponseBodyData {
	return s.Data
}

func (s *DeleteApgroupSsidConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DeleteApgroupSsidConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteApgroupSsidConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteApgroupSsidConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApgroupSsidConfigResponseBody) SetData(v *DeleteApgroupSsidConfigResponseBodyData) *DeleteApgroupSsidConfigResponseBody {
	s.Data = v
	return s
}

func (s *DeleteApgroupSsidConfigResponseBody) SetErrorCode(v int32) *DeleteApgroupSsidConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApgroupSsidConfigResponseBody) SetErrorMessage(v string) *DeleteApgroupSsidConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteApgroupSsidConfigResponseBody) SetIsSuccess(v bool) *DeleteApgroupSsidConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteApgroupSsidConfigResponseBody) SetRequestId(v string) *DeleteApgroupSsidConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApgroupSsidConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteApgroupSsidConfigResponseBodyData struct {
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteApgroupSsidConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteApgroupSsidConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteApgroupSsidConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteApgroupSsidConfigResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteApgroupSsidConfigResponseBodyData) SetId(v int64) *DeleteApgroupSsidConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteApgroupSsidConfigResponseBodyData) SetTaskId(v string) *DeleteApgroupSsidConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteApgroupSsidConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
