// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakeHouseSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteLakeHouseSpaceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteLakeHouseSpaceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteLakeHouseSpaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLakeHouseSpaceResponseBody
	GetSuccess() *bool
}

type DeleteLakeHouseSpaceResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E76DD2E7-EBAC-5724-B163-19AAC233F8F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLakeHouseSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakeHouseSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLakeHouseSpaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteLakeHouseSpaceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteLakeHouseSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLakeHouseSpaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLakeHouseSpaceResponseBody) SetErrorCode(v string) *DeleteLakeHouseSpaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteLakeHouseSpaceResponseBody) SetErrorMessage(v string) *DeleteLakeHouseSpaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteLakeHouseSpaceResponseBody) SetRequestId(v string) *DeleteLakeHouseSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLakeHouseSpaceResponseBody) SetSuccess(v bool) *DeleteLakeHouseSpaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLakeHouseSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
