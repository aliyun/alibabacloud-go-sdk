// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakeHouseSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateLakeHouseSpaceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateLakeHouseSpaceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateLakeHouseSpaceResponseBody
	GetRequestId() *string
	SetSpaceId(v int64) *CreateLakeHouseSpaceResponseBody
	GetSpaceId() *int64
	SetSuccess(v bool) *CreateLakeHouseSpaceResponseBody
	GetSuccess() *bool
}

type CreateLakeHouseSpaceResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// InvalidParameterValid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// Invalid parameters: space name,mode,prod db id,db type,config.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 24
	SpaceId *int64 `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLakeHouseSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLakeHouseSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLakeHouseSpaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateLakeHouseSpaceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateLakeHouseSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLakeHouseSpaceResponseBody) GetSpaceId() *int64 {
	return s.SpaceId
}

func (s *CreateLakeHouseSpaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLakeHouseSpaceResponseBody) SetErrorCode(v string) *CreateLakeHouseSpaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateLakeHouseSpaceResponseBody) SetErrorMessage(v string) *CreateLakeHouseSpaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLakeHouseSpaceResponseBody) SetRequestId(v string) *CreateLakeHouseSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLakeHouseSpaceResponseBody) SetSpaceId(v int64) *CreateLakeHouseSpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *CreateLakeHouseSpaceResponseBody) SetSuccess(v bool) *CreateLakeHouseSpaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLakeHouseSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
