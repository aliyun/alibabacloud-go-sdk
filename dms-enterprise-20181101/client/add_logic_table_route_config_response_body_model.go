// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLogicTableRouteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddLogicTableRouteConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddLogicTableRouteConfigResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddLogicTableRouteConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddLogicTableRouteConfigResponseBody
	GetSuccess() *bool
}

type AddLogicTableRouteConfigResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// can not find table, tableId : 11133
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B43AD641-49C2-5299-9E06-1B37EC1B****
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

func (s AddLogicTableRouteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLogicTableRouteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLogicTableRouteConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddLogicTableRouteConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddLogicTableRouteConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLogicTableRouteConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddLogicTableRouteConfigResponseBody) SetErrorCode(v string) *AddLogicTableRouteConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddLogicTableRouteConfigResponseBody) SetErrorMessage(v string) *AddLogicTableRouteConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddLogicTableRouteConfigResponseBody) SetRequestId(v string) *AddLogicTableRouteConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLogicTableRouteConfigResponseBody) SetSuccess(v bool) *AddLogicTableRouteConfigResponseBody {
	s.Success = &v
	return s
}

func (s *AddLogicTableRouteConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
