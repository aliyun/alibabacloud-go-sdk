// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlTransTableMetaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *GetSqlTransTableMetaInfoResponseBody
	GetData() []*string
	SetErrCode(v string) *GetSqlTransTableMetaInfoResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetSqlTransTableMetaInfoResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetSqlTransTableMetaInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSqlTransTableMetaInfoResponseBody
	GetSuccess() *bool
}

type GetSqlTransTableMetaInfoResponseBody struct {
	// The list of table names parsed from the script.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. This value is an empty string if the call succeeds.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. This value is an empty string if the call succeeds.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID that uniquely identifies the call. Provide this value when troubleshooting issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSqlTransTableMetaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlTransTableMetaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlTransTableMetaInfoResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetSqlTransTableMetaInfoResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetSqlTransTableMetaInfoResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetSqlTransTableMetaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlTransTableMetaInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSqlTransTableMetaInfoResponseBody) SetData(v []*string) *GetSqlTransTableMetaInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlTransTableMetaInfoResponseBody) SetErrCode(v string) *GetSqlTransTableMetaInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSqlTransTableMetaInfoResponseBody) SetErrMessage(v string) *GetSqlTransTableMetaInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetSqlTransTableMetaInfoResponseBody) SetRequestId(v string) *GetSqlTransTableMetaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlTransTableMetaInfoResponseBody) SetSuccess(v bool) *GetSqlTransTableMetaInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlTransTableMetaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
