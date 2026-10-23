// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResourcePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAppResourcePermissionsResponseBody
	GetCode() *string
	SetDeletedCount(v int64) *DeleteAppResourcePermissionsResponseBody
	GetDeletedCount() *int64
	SetMessage(v string) *DeleteAppResourcePermissionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAppResourcePermissionsResponseBody
	GetRequestId() *string
}

type DeleteAppResourcePermissionsResponseBody struct {
	// The business status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The number of deleted permission records.
	//
	// example:
	//
	// 2
	DeletedCount *int64 `json:"deletedCount,omitempty" xml:"deletedCount,omitempty"`
	// The error message. This field is empty if the request is successful.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAppResourcePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResourcePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResourcePermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAppResourcePermissionsResponseBody) GetDeletedCount() *int64 {
	return s.DeletedCount
}

func (s *DeleteAppResourcePermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAppResourcePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppResourcePermissionsResponseBody) SetCode(v string) *DeleteAppResourcePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponseBody) SetDeletedCount(v int64) *DeleteAppResourcePermissionsResponseBody {
	s.DeletedCount = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponseBody) SetMessage(v string) *DeleteAppResourcePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponseBody) SetRequestId(v string) *DeleteAppResourcePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResourcePermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
