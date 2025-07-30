// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ModifyDBClusterConfigResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *ModifyDBClusterConfigResponseBodyData) *ModifyDBClusterConfigResponseBody
	GetData() *ModifyDBClusterConfigResponseBodyData
	SetDynamicCode(v string) *ModifyDBClusterConfigResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ModifyDBClusterConfigResponseBody
	GetDynamicMessage() *string
	SetRequestId(v string) *ModifyDBClusterConfigResponseBody
	GetRequestId() *string
}

type ModifyDBClusterConfigResponseBody struct {
	// example:
	//
	// failed
	AccessDeniedDetail *string                                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *ModifyDBClusterConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic code. This parameter is not returned.
	//
	// example:
	//
	// 0
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message. This parameter is not returned.
	//
	// example:
	//
	// An error occurred while processing your request.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// BC854513-E85E-54F3-9842-B9CCD3308CDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ModifyDBClusterConfigResponseBody) GetData() *ModifyDBClusterConfigResponseBodyData {
	return s.Data
}

func (s *ModifyDBClusterConfigResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ModifyDBClusterConfigResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyDBClusterConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterConfigResponseBody) SetAccessDeniedDetail(v string) *ModifyDBClusterConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) SetData(v *ModifyDBClusterConfigResponseBodyData) *ModifyDBClusterConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) SetDynamicCode(v string) *ModifyDBClusterConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) SetDynamicMessage(v string) *ModifyDBClusterConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) SetRequestId(v string) *ModifyDBClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyDBClusterConfigResponseBodyData struct {
	// example:
	//
	// selectdb-cn-wny3li00g02-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// example:
	//
	// 6585
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// example:
	//
	// selectdb-cn-wny3li00g02
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// example:
	//
	// 107878719
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBClusterConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponseBodyData) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *ModifyDBClusterConfigResponseBodyData) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyDBClusterConfigResponseBodyData) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *ModifyDBClusterConfigResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyDBClusterConfigResponseBodyData) SetDbClusterId(v string) *ModifyDBClusterConfigResponseBodyData {
	s.DbClusterId = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBodyData) SetDbInstanceId(v string) *ModifyDBClusterConfigResponseBodyData {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBodyData) SetDbInstanceName(v string) *ModifyDBClusterConfigResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBodyData) SetTaskId(v int32) *ModifyDBClusterConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
