// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterExportMemberBalanceOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ExportTaskDTO) *ModelRouterExportMemberBalanceOrdersResponseBody
	GetData() *ExportTaskDTO
	SetErrCode(v string) *ModelRouterExportMemberBalanceOrdersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterExportMemberBalanceOrdersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterExportMemberBalanceOrdersResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterExportMemberBalanceOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterExportMemberBalanceOrdersResponseBody
	GetSuccess() *bool
}

type ModelRouterExportMemberBalanceOrdersResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *ExportTaskDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterExportMemberBalanceOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterExportMemberBalanceOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) GetData() *ExportTaskDTO {
	return s.Data
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) SetData(v *ExportTaskDTO) *ModelRouterExportMemberBalanceOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) SetErrCode(v string) *ModelRouterExportMemberBalanceOrdersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) SetErrMessage(v string) *ModelRouterExportMemberBalanceOrdersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) SetHttpStatusCode(v int32) *ModelRouterExportMemberBalanceOrdersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) SetRequestId(v string) *ModelRouterExportMemberBalanceOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) SetSuccess(v bool) *ModelRouterExportMemberBalanceOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterExportMemberBalanceOrdersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
