// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordRunCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOperationRecordRunCodeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetOperationRecordRunCodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOperationRecordRunCodeResponseBody
	GetMessage() *string
	SetOperationLogCodeResponse(v *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) *GetOperationRecordRunCodeResponseBody
	GetOperationLogCodeResponse() *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse
	SetRequestId(v string) *GetOperationRecordRunCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOperationRecordRunCodeResponseBody
	GetSuccess() *bool
}

type GetOperationRecordRunCodeResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The operation log code response.
	OperationLogCodeResponse *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse `json:"OperationLogCodeResponse,omitempty" xml:"OperationLogCodeResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOperationRecordRunCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordRunCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationRecordRunCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOperationRecordRunCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOperationRecordRunCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOperationRecordRunCodeResponseBody) GetOperationLogCodeResponse() *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse {
	return s.OperationLogCodeResponse
}

func (s *GetOperationRecordRunCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOperationRecordRunCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOperationRecordRunCodeResponseBody) SetCode(v string) *GetOperationRecordRunCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBody) SetHttpStatusCode(v int32) *GetOperationRecordRunCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBody) SetMessage(v string) *GetOperationRecordRunCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBody) SetOperationLogCodeResponse(v *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) *GetOperationRecordRunCodeResponseBody {
	s.OperationLogCodeResponse = v
	return s
}

func (s *GetOperationRecordRunCodeResponseBody) SetRequestId(v string) *GetOperationRecordRunCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBody) SetSuccess(v bool) *GetOperationRecordRunCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBody) Validate() error {
	if s.OperationLogCodeResponse != nil {
		if err := s.OperationLogCodeResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse struct {
	// The code content.
	//
	// example:
	//
	// SELECT 	- FROM test_table WHERE id = 1;
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 123
	OperatorId *int32 `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// The operator name.
	//
	// example:
	//
	// SQL查询节点
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The number of SQL statements.
	//
	// example:
	//
	// 1
	SqlNum *int32 `json:"SqlNum,omitempty" xml:"SqlNum,omitempty"`
}

func (s GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) GoString() string {
	return s.String()
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) GetCode() *string {
	return s.Code
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) GetOperatorId() *int32 {
	return s.OperatorId
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) GetSqlNum() *int32 {
	return s.SqlNum
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) SetCode(v string) *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse {
	s.Code = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) SetOperatorId(v int32) *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse {
	s.OperatorId = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) SetOperatorName(v string) *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse {
	s.OperatorName = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) SetSqlNum(v int32) *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse {
	s.SqlNum = &v
	return s
}

func (s *GetOperationRecordRunCodeResponseBodyOperationLogCodeResponse) Validate() error {
	return dara.Validate(s)
}
