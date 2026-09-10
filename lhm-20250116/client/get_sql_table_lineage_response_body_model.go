// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlTableLineageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSqlTableLineageResponseBodyData) *GetSqlTableLineageResponseBody
	GetData() *GetSqlTableLineageResponseBodyData
	SetErrCode(v string) *GetSqlTableLineageResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetSqlTableLineageResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetSqlTableLineageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSqlTableLineageResponseBody
	GetSuccess() *bool
}

type GetSqlTableLineageResponseBody struct {
	// The data body returned by the operation. For the field structure, see the child parameters.
	Data *GetSqlTableLineageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues with the current call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed. Use errCode and errMessage to troubleshoot the issue.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSqlTableLineageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlTableLineageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlTableLineageResponseBody) GetData() *GetSqlTableLineageResponseBodyData {
	return s.Data
}

func (s *GetSqlTableLineageResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetSqlTableLineageResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetSqlTableLineageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlTableLineageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSqlTableLineageResponseBody) SetData(v *GetSqlTableLineageResponseBodyData) *GetSqlTableLineageResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlTableLineageResponseBody) SetErrCode(v string) *GetSqlTableLineageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSqlTableLineageResponseBody) SetErrMessage(v string) *GetSqlTableLineageResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetSqlTableLineageResponseBody) SetRequestId(v string) *GetSqlTableLineageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlTableLineageResponseBody) SetSuccess(v bool) *GetSqlTableLineageResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlTableLineageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSqlTableLineageResponseBodyData struct {
	// The list of downstream tables.
	DownstreamTables []*string `json:"downstreamTables,omitempty" xml:"downstreamTables,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// connection timeout
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed. Use errCode and errMessage to troubleshoot the issue.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The list of upstream tables.
	UpstreamTables []*string `json:"upstreamTables,omitempty" xml:"upstreamTables,omitempty" type:"Repeated"`
}

func (s GetSqlTableLineageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSqlTableLineageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlTableLineageResponseBodyData) GetDownstreamTables() []*string {
	return s.DownstreamTables
}

func (s *GetSqlTableLineageResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetSqlTableLineageResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GetSqlTableLineageResponseBodyData) GetUpstreamTables() []*string {
	return s.UpstreamTables
}

func (s *GetSqlTableLineageResponseBodyData) SetDownstreamTables(v []*string) *GetSqlTableLineageResponseBodyData {
	s.DownstreamTables = v
	return s
}

func (s *GetSqlTableLineageResponseBodyData) SetErrorMsg(v string) *GetSqlTableLineageResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetSqlTableLineageResponseBodyData) SetSuccess(v bool) *GetSqlTableLineageResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetSqlTableLineageResponseBodyData) SetUpstreamTables(v []*string) *GetSqlTableLineageResponseBodyData {
	s.UpstreamTables = v
	return s
}

func (s *GetSqlTableLineageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
