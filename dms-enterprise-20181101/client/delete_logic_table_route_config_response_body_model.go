// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicTableRouteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteLogicTableRouteConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteLogicTableRouteConfigResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteLogicTableRouteConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLogicTableRouteConfigResponseBody
	GetSuccess() *bool
}

type DeleteLogicTableRouteConfigResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// MissingRouteKey
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// RouteKey is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A916A004-A88C-5B39-ABDB-DE808E80****
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

func (s DeleteLogicTableRouteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicTableRouteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogicTableRouteConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteLogicTableRouteConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteLogicTableRouteConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogicTableRouteConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLogicTableRouteConfigResponseBody) SetErrorCode(v string) *DeleteLogicTableRouteConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteLogicTableRouteConfigResponseBody) SetErrorMessage(v string) *DeleteLogicTableRouteConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteLogicTableRouteConfigResponseBody) SetRequestId(v string) *DeleteLogicTableRouteConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogicTableRouteConfigResponseBody) SetSuccess(v bool) *DeleteLogicTableRouteConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLogicTableRouteConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
