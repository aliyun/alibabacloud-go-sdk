// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseEngineMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CloseEngineMigrationResponseBody
	GetCode() *string
	SetData(v *CloseEngineMigrationResponseBodyData) *CloseEngineMigrationResponseBody
	GetData() *CloseEngineMigrationResponseBodyData
	SetHttpStatusCode(v string) *CloseEngineMigrationResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CloseEngineMigrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloseEngineMigrationResponseBody
	GetRequestId() *string
	SetSuccessResponse(v string) *CloseEngineMigrationResponseBody
	GetSuccessResponse() *string
}

type CloseEngineMigrationResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloseEngineMigrationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// *****
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

func (s CloseEngineMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseEngineMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *CloseEngineMigrationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloseEngineMigrationResponseBody) GetData() *CloseEngineMigrationResponseBodyData {
	return s.Data
}

func (s *CloseEngineMigrationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CloseEngineMigrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloseEngineMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseEngineMigrationResponseBody) GetSuccessResponse() *string {
	return s.SuccessResponse
}

func (s *CloseEngineMigrationResponseBody) SetCode(v string) *CloseEngineMigrationResponseBody {
	s.Code = &v
	return s
}

func (s *CloseEngineMigrationResponseBody) SetData(v *CloseEngineMigrationResponseBodyData) *CloseEngineMigrationResponseBody {
	s.Data = v
	return s
}

func (s *CloseEngineMigrationResponseBody) SetHttpStatusCode(v string) *CloseEngineMigrationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CloseEngineMigrationResponseBody) SetMessage(v string) *CloseEngineMigrationResponseBody {
	s.Message = &v
	return s
}

func (s *CloseEngineMigrationResponseBody) SetRequestId(v string) *CloseEngineMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseEngineMigrationResponseBody) SetSuccessResponse(v string) *CloseEngineMigrationResponseBody {
	s.SuccessResponse = &v
	return s
}

func (s *CloseEngineMigrationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloseEngineMigrationResponseBodyData struct {
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 2209883
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloseEngineMigrationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloseEngineMigrationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloseEngineMigrationResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CloseEngineMigrationResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CloseEngineMigrationResponseBodyData) SetDBInstanceName(v string) *CloseEngineMigrationResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CloseEngineMigrationResponseBodyData) SetTaskId(v string) *CloseEngineMigrationResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CloseEngineMigrationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
