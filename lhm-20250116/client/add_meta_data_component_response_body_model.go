// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMetaDataComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *AddMetaDataComponentResponseBody
	GetData() *int64
	SetErrCode(v string) *AddMetaDataComponentResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AddMetaDataComponentResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *AddMetaDataComponentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddMetaDataComponentResponseBody
	GetSuccess() *bool
}

type AddMetaDataComponentResponseBody struct {
	// The ID of the newly created data source. Subsequent operations reference this data source by this ID or the data source name.
	//
	// example:
	//
	// 100
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
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
	// The request ID, which is used to locate and troubleshoot issues with this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to troubleshoot the issue.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddMetaDataComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMetaDataComponentResponseBody) GoString() string {
	return s.String()
}

func (s *AddMetaDataComponentResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddMetaDataComponentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AddMetaDataComponentResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AddMetaDataComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMetaDataComponentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMetaDataComponentResponseBody) SetData(v int64) *AddMetaDataComponentResponseBody {
	s.Data = &v
	return s
}

func (s *AddMetaDataComponentResponseBody) SetErrCode(v string) *AddMetaDataComponentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddMetaDataComponentResponseBody) SetErrMessage(v string) *AddMetaDataComponentResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddMetaDataComponentResponseBody) SetRequestId(v string) *AddMetaDataComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMetaDataComponentResponseBody) SetSuccess(v bool) *AddMetaDataComponentResponseBody {
	s.Success = &v
	return s
}

func (s *AddMetaDataComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
