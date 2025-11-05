// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSchemaPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RevokeSchemaPermissionResponseBody
	GetData() *bool
	SetRequestId(v string) *RevokeSchemaPermissionResponseBody
	GetRequestId() *string
}

type RevokeSchemaPermissionResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RevokeSchemaPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeSchemaPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeSchemaPermissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *RevokeSchemaPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeSchemaPermissionResponseBody) SetData(v bool) *RevokeSchemaPermissionResponseBody {
	s.Data = &v
	return s
}

func (s *RevokeSchemaPermissionResponseBody) SetRequestId(v string) *RevokeSchemaPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeSchemaPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
