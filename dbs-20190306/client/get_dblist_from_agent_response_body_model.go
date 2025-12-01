// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBListFromAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbList(v *GetDBListFromAgentResponseBodyDbList) *GetDBListFromAgentResponseBody
	GetDbList() *GetDBListFromAgentResponseBodyDbList
	SetErrCode(v string) *GetDBListFromAgentResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetDBListFromAgentResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *GetDBListFromAgentResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetDBListFromAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDBListFromAgentResponseBody
	GetSuccess() *bool
}

type GetDBListFromAgentResponseBody struct {
	// The details of the databases.
	DbList *GetDBListFromAgentResponseBodyDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This backupPlan can\\"t support this action
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EB4DFD5E-3618-498D-BE35-4DBEA0072122
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDBListFromAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDBListFromAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBListFromAgentResponseBody) GetDbList() *GetDBListFromAgentResponseBodyDbList {
	return s.DbList
}

func (s *GetDBListFromAgentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDBListFromAgentResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetDBListFromAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDBListFromAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDBListFromAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDBListFromAgentResponseBody) SetDbList(v *GetDBListFromAgentResponseBodyDbList) *GetDBListFromAgentResponseBody {
	s.DbList = v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetErrCode(v string) *GetDBListFromAgentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetErrMessage(v string) *GetDBListFromAgentResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetHttpStatusCode(v int32) *GetDBListFromAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetRequestId(v string) *GetDBListFromAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) SetSuccess(v bool) *GetDBListFromAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetDBListFromAgentResponseBody) Validate() error {
	if s.DbList != nil {
		if err := s.DbList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDBListFromAgentResponseBodyDbList struct {
	DbName []*string `json:"dbName,omitempty" xml:"dbName,omitempty" type:"Repeated"`
}

func (s GetDBListFromAgentResponseBodyDbList) String() string {
	return dara.Prettify(s)
}

func (s GetDBListFromAgentResponseBodyDbList) GoString() string {
	return s.String()
}

func (s *GetDBListFromAgentResponseBodyDbList) GetDbName() []*string {
	return s.DbName
}

func (s *GetDBListFromAgentResponseBodyDbList) SetDbName(v []*string) *GetDBListFromAgentResponseBodyDbList {
	s.DbName = v
	return s
}

func (s *GetDBListFromAgentResponseBodyDbList) Validate() error {
	return dara.Validate(s)
}
