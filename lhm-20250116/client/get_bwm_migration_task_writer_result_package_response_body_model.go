// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationTaskWriterResultPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetBwmMigrationTaskWriterResultPackageResponseBody
	GetData() *string
	SetErrCode(v string) *GetBwmMigrationTaskWriterResultPackageResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetBwmMigrationTaskWriterResultPackageResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetBwmMigrationTaskWriterResultPackageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBwmMigrationTaskWriterResultPackageResponseBody
	GetSuccess() *bool
}

type GetBwmMigrationTaskWriterResultPackageResponseBody struct {
	// The response data.
	//
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// - true: The call is successful.
	//
	// - false: The call failed. Use errCode and errMessage to troubleshoot the issue.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetBwmMigrationTaskWriterResultPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationTaskWriterResultPackageResponseBody) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) GetData() *string {
	return s.Data
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) SetData(v string) *GetBwmMigrationTaskWriterResultPackageResponseBody {
	s.Data = &v
	return s
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) SetErrCode(v string) *GetBwmMigrationTaskWriterResultPackageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) SetErrMessage(v string) *GetBwmMigrationTaskWriterResultPackageResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) SetRequestId(v string) *GetBwmMigrationTaskWriterResultPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) SetSuccess(v bool) *GetBwmMigrationTaskWriterResultPackageResponseBody {
	s.Success = &v
	return s
}

func (s *GetBwmMigrationTaskWriterResultPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
