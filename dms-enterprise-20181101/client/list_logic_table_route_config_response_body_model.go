// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicTableRouteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListLogicTableRouteConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListLogicTableRouteConfigResponseBody
	GetErrorMessage() *string
	SetLogicTableRouteConfigList(v *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList) *ListLogicTableRouteConfigResponseBody
	GetLogicTableRouteConfigList() *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList
	SetRequestId(v string) *ListLogicTableRouteConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLogicTableRouteConfigResponseBody
	GetSuccess() *bool
}

type ListLogicTableRouteConfigResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// MissingTableId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// TableId is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The routing algorithms.
	LogicTableRouteConfigList *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList `json:"LogicTableRouteConfigList,omitempty" xml:"LogicTableRouteConfigList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7C6A0D7D-B034-59F6-854C-32425AC6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLogicTableRouteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTableRouteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogicTableRouteConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListLogicTableRouteConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListLogicTableRouteConfigResponseBody) GetLogicTableRouteConfigList() *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList {
	return s.LogicTableRouteConfigList
}

func (s *ListLogicTableRouteConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogicTableRouteConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLogicTableRouteConfigResponseBody) SetErrorCode(v string) *ListLogicTableRouteConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLogicTableRouteConfigResponseBody) SetErrorMessage(v string) *ListLogicTableRouteConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListLogicTableRouteConfigResponseBody) SetLogicTableRouteConfigList(v *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList) *ListLogicTableRouteConfigResponseBody {
	s.LogicTableRouteConfigList = v
	return s
}

func (s *ListLogicTableRouteConfigResponseBody) SetRequestId(v string) *ListLogicTableRouteConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogicTableRouteConfigResponseBody) SetSuccess(v bool) *ListLogicTableRouteConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListLogicTableRouteConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList struct {
	LogicTableRouteConfig []*ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig `json:"LogicTableRouteConfig,omitempty" xml:"LogicTableRouteConfig,omitempty" type:"Repeated"`
}

func (s ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList) GoString() string {
	return s.String()
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList) GetLogicTableRouteConfig() []*ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig {
	return s.LogicTableRouteConfig
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList) SetLogicTableRouteConfig(v []*ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList {
	s.LogicTableRouteConfig = v
	return s
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigList) Validate() error {
	return dara.Validate(s)
}

type ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig struct {
	// The routing algorithm expression.
	//
	// example:
	//
	// #id#%16\\t
	RouteExpr *string `json:"RouteExpr,omitempty" xml:"RouteExpr,omitempty"`
	// The unique key of the routing algorithm.
	//
	// example:
	//
	// 1
	RouteKey *string `json:"RouteKey,omitempty" xml:"RouteKey,omitempty"`
	// The ID of the logical table.
	//
	// example:
	//
	// 4****
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) GoString() string {
	return s.String()
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) GetRouteExpr() *string {
	return s.RouteExpr
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) GetRouteKey() *string {
	return s.RouteKey
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) GetTableId() *int64 {
	return s.TableId
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) SetRouteExpr(v string) *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig {
	s.RouteExpr = &v
	return s
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) SetRouteKey(v string) *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig {
	s.RouteKey = &v
	return s
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) SetTableId(v int64) *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig {
	s.TableId = &v
	return s
}

func (s *ListLogicTableRouteConfigResponseBodyLogicTableRouteConfigListLogicTableRouteConfig) Validate() error {
	return dara.Validate(s)
}
