// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataCheckConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *AddDataCheckConfigResponseBody
	GetData() *int64
	SetErrCode(v string) *AddDataCheckConfigResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AddDataCheckConfigResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *AddDataCheckConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDataCheckConfigResponseBody
	GetSuccess() *bool
}

type AddDataCheckConfigResponseBody struct {
	// The ID of the table-level configuration saved in this operation. You can use this ID to query or delete the configuration later.
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
	// Indicates whether the call is successful. Valid values:
	//
	// - true: successful.
	//
	// - false: failed. Check errCode and errMessage for the cause.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDataCheckConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataCheckConfigResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddDataCheckConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AddDataCheckConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AddDataCheckConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataCheckConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataCheckConfigResponseBody) SetData(v int64) *AddDataCheckConfigResponseBody {
	s.Data = &v
	return s
}

func (s *AddDataCheckConfigResponseBody) SetErrCode(v string) *AddDataCheckConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddDataCheckConfigResponseBody) SetErrMessage(v string) *AddDataCheckConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AddDataCheckConfigResponseBody) SetRequestId(v string) *AddDataCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataCheckConfigResponseBody) SetSuccess(v bool) *AddDataCheckConfigResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataCheckConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
