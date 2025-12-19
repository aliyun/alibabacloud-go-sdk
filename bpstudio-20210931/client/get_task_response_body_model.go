// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTaskResponseBody
	GetCode() *string
	SetData(v *GetTaskResponseBodyData) *GetTaskResponseBody
	GetData() *GetTaskResponseBodyData
	SetMessage(v string) *GetTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
}

type GetTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTaskResponseBody) GetData() *GetTaskResponseBodyData {
	return s.Data
}

func (s *GetTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) SetCode(v string) *GetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBody) SetData(v *GetTaskResponseBodyData) *GetTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResponseBody) SetMessage(v string) *GetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyData struct {
	// example:
	//
	// JXT**LZJ568665D6
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// SyntaxError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 2025-08-21 14:07:02
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1333
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {\\"app\\":\\"JJXT**LZJ568665D6\\"}
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 64QQBC02**FA3W3G
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTaskResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTaskResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *GetTaskResponseBodyData) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTaskResponseBodyData) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBodyData) SetAppId(v string) *GetTaskResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetTaskResponseBodyData) SetErrorMessage(v string) *GetTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskResponseBodyData) SetGmtCreate(v string) *GetTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetTaskResponseBodyData) SetId(v int32) *GetTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTaskResponseBodyData) SetInstanceIds(v string) *GetTaskResponseBodyData {
	s.InstanceIds = &v
	return s
}

func (s *GetTaskResponseBodyData) SetName(v string) *GetTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyData) SetProcessId(v string) *GetTaskResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GetTaskResponseBodyData) SetStatus(v string) *GetTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
