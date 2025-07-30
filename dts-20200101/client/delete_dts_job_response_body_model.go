// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DeleteDtsJobResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteDtsJobResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DeleteDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteDtsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DeleteDtsJobResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteDtsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDtsJobResponseBody
	GetSuccess() *bool
}

type DeleteDtsJobResponseBody struct {
	// The operation that you want to perform. Set the value to **DeleteDtsJob**.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The operation that you want to perform. Set the value to **DeleteDtsJob**.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**: Data is synchronized from the source database to the destination database.
	//
	// 	- **Reverse**: Data is synchronized from the destination database to the source database.
	//
	// >
	//
	// 	- Default value: **Forward**.
	//
	// 	- You can set this parameter to **Reverse*	- to delete the reverse synchronization task only if the topology is two-way synchronization.
	//
	// example:
	//
	// 01B6F25-21E7-4484-99D5-3EF2625C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the region in which the data migration or synchronization task resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteDtsJobResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteDtsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteDtsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDtsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDtsJobResponseBody) SetDynamicCode(v string) *DeleteDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetDynamicMessage(v string) *DeleteDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetErrCode(v string) *DeleteDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetErrMessage(v string) *DeleteDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetHttpStatusCode(v int32) *DeleteDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetRequestId(v string) *DeleteDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetSuccess(v bool) *DeleteDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
