// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeDatabasePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RevokeDatabasePermissionResponseBody
	GetData() *bool
	SetRequestId(v string) *RevokeDatabasePermissionResponseBody
	GetRequestId() *string
}

type RevokeDatabasePermissionResponseBody struct {
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

func (s RevokeDatabasePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeDatabasePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeDatabasePermissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *RevokeDatabasePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeDatabasePermissionResponseBody) SetData(v bool) *RevokeDatabasePermissionResponseBody {
	s.Data = &v
	return s
}

func (s *RevokeDatabasePermissionResponseBody) SetRequestId(v string) *RevokeDatabasePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeDatabasePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
