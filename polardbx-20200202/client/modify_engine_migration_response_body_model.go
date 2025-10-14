// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEngineMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyEngineMigrationResponseBody
	GetCode() *string
	SetData(v *ModifyEngineMigrationResponseBodyData) *ModifyEngineMigrationResponseBody
	GetData() *ModifyEngineMigrationResponseBodyData
	SetHttpStatusCode(v string) *ModifyEngineMigrationResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *ModifyEngineMigrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyEngineMigrationResponseBody
	GetRequestId() *string
	SetSuccessResponse(v string) *ModifyEngineMigrationResponseBody
	GetSuccessResponse() *string
}

type ModifyEngineMigrationResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ModifyEngineMigrationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SuccessResponse *string `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s ModifyEngineMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEngineMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEngineMigrationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyEngineMigrationResponseBody) GetData() *ModifyEngineMigrationResponseBodyData {
	return s.Data
}

func (s *ModifyEngineMigrationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ModifyEngineMigrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyEngineMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEngineMigrationResponseBody) GetSuccessResponse() *string {
	return s.SuccessResponse
}

func (s *ModifyEngineMigrationResponseBody) SetCode(v string) *ModifyEngineMigrationResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyEngineMigrationResponseBody) SetData(v *ModifyEngineMigrationResponseBodyData) *ModifyEngineMigrationResponseBody {
	s.Data = v
	return s
}

func (s *ModifyEngineMigrationResponseBody) SetHttpStatusCode(v string) *ModifyEngineMigrationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyEngineMigrationResponseBody) SetMessage(v string) *ModifyEngineMigrationResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyEngineMigrationResponseBody) SetRequestId(v string) *ModifyEngineMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEngineMigrationResponseBody) SetSuccessResponse(v string) *ModifyEngineMigrationResponseBody {
	s.SuccessResponse = &v
	return s
}

func (s *ModifyEngineMigrationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyEngineMigrationResponseBodyData struct {
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 42292****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyEngineMigrationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyEngineMigrationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyEngineMigrationResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyEngineMigrationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyEngineMigrationResponseBodyData) SetDBInstanceName(v string) *ModifyEngineMigrationResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyEngineMigrationResponseBodyData) SetTaskId(v string) *ModifyEngineMigrationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyEngineMigrationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
