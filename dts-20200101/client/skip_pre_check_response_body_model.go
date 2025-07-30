// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipPreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SkipPreCheckResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *SkipPreCheckResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *SkipPreCheckResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SkipPreCheckResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *SkipPreCheckResponseBody
	GetHttpStatusCode() *int32
	SetMigrationJobId(v string) *SkipPreCheckResponseBody
	GetMigrationJobId() *string
	SetRequestId(v string) *SkipPreCheckResponseBody
	GetRequestId() *string
	SetScheduleJobId(v string) *SkipPreCheckResponseBody
	GetScheduleJobId() *string
	SetSkipItems(v string) *SkipPreCheckResponseBody
	GetSkipItems() *string
	SetSkipNames(v string) *SkipPreCheckResponseBody
	GetSkipNames() *string
	SetSuccess(v bool) *SkipPreCheckResponseBody
	GetSuccess() *bool
}

type SkipPreCheckResponseBody struct {
	// The error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace %s in ErrMessage.
	//
	// > If JobId is invalid, JobId is returned for DynamicMessage, and the following message is returned for ErrMessage: The Value of Input Parameter %s is not valid.
	//
	// example:
	//
	// JobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status codes returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The precheck task ID.
	//
	// example:
	//
	// b4my3zg929a****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C498360-7892-433C-847A-BA71A850****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The precheck task ID.
	//
	// example:
	//
	// b4my3zg929a****
	ScheduleJobId *string `json:"ScheduleJobId,omitempty" xml:"ScheduleJobId,omitempty"`
	// The shortened name of the precheck item.
	//
	// example:
	//
	// CHECK_SAME_OBJ
	SkipItems *string `json:"SkipItems,omitempty" xml:"SkipItems,omitempty"`
	// The precheck item name.
	//
	// example:
	//
	// CHECK_SAME_OBJ_DETAIL
	SkipNames *string `json:"SkipNames,omitempty" xml:"SkipNames,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SkipPreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SkipPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SkipPreCheckResponseBody) GetCode() *string {
	return s.Code
}

func (s *SkipPreCheckResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SkipPreCheckResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SkipPreCheckResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SkipPreCheckResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SkipPreCheckResponseBody) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *SkipPreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SkipPreCheckResponseBody) GetScheduleJobId() *string {
	return s.ScheduleJobId
}

func (s *SkipPreCheckResponseBody) GetSkipItems() *string {
	return s.SkipItems
}

func (s *SkipPreCheckResponseBody) GetSkipNames() *string {
	return s.SkipNames
}

func (s *SkipPreCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SkipPreCheckResponseBody) SetCode(v string) *SkipPreCheckResponseBody {
	s.Code = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetDynamicMessage(v string) *SkipPreCheckResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetErrCode(v string) *SkipPreCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetErrMessage(v string) *SkipPreCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetHttpStatusCode(v int32) *SkipPreCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetMigrationJobId(v string) *SkipPreCheckResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetRequestId(v string) *SkipPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetScheduleJobId(v string) *SkipPreCheckResponseBody {
	s.ScheduleJobId = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetSkipItems(v string) *SkipPreCheckResponseBody {
	s.SkipItems = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetSkipNames(v string) *SkipPreCheckResponseBody {
	s.SkipNames = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetSuccess(v bool) *SkipPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *SkipPreCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
