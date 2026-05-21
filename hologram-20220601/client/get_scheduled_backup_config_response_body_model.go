// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledBackupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetScheduledBackupConfigResponseBodyData) *GetScheduledBackupConfigResponseBody
	GetData() *GetScheduledBackupConfigResponseBodyData
	SetErrorCode(v string) *GetScheduledBackupConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetScheduledBackupConfigResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetScheduledBackupConfigResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GetScheduledBackupConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScheduledBackupConfigResponseBody
	GetSuccess() *bool
}

type GetScheduledBackupConfigResponseBody struct {
	Data *GetScheduledBackupConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 623EF0B6-A6C1-5764-82CC-0BC176654CB9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScheduledBackupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledBackupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledBackupConfigResponseBody) GetData() *GetScheduledBackupConfigResponseBodyData {
	return s.Data
}

func (s *GetScheduledBackupConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetScheduledBackupConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetScheduledBackupConfigResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetScheduledBackupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledBackupConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScheduledBackupConfigResponseBody) SetData(v *GetScheduledBackupConfigResponseBodyData) *GetScheduledBackupConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetScheduledBackupConfigResponseBody) SetErrorCode(v string) *GetScheduledBackupConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetScheduledBackupConfigResponseBody) SetErrorMessage(v string) *GetScheduledBackupConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetScheduledBackupConfigResponseBody) SetHttpStatusCode(v string) *GetScheduledBackupConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetScheduledBackupConfigResponseBody) SetRequestId(v string) *GetScheduledBackupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledBackupConfigResponseBody) SetSuccess(v bool) *GetScheduledBackupConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetScheduledBackupConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScheduledBackupConfigResponseBodyData struct {
	// example:
	//
	// true
	Enabled *int64 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// {"schedule":"0 10 	- 	- 1,2","instance_id":"hgprecn-cn-lbj3dfgog004","week":"1,2","hour":10,"data_keep_quantity":3,"type":"periodical","enabled":true}
	TaskParameter *string `json:"TaskParameter,omitempty" xml:"TaskParameter,omitempty"`
	// example:
	//
	// backup
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetScheduledBackupConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledBackupConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScheduledBackupConfigResponseBodyData) GetEnabled() *int64 {
	return s.Enabled
}

func (s *GetScheduledBackupConfigResponseBodyData) GetTaskParameter() *string {
	return s.TaskParameter
}

func (s *GetScheduledBackupConfigResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetScheduledBackupConfigResponseBodyData) SetEnabled(v int64) *GetScheduledBackupConfigResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetScheduledBackupConfigResponseBodyData) SetTaskParameter(v string) *GetScheduledBackupConfigResponseBodyData {
	s.TaskParameter = &v
	return s
}

func (s *GetScheduledBackupConfigResponseBodyData) SetTaskType(v string) *GetScheduledBackupConfigResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetScheduledBackupConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
