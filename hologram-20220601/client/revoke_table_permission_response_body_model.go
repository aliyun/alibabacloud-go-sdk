// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTablePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RevokeTablePermissionResponseBody
	GetData() *bool
	SetRequestId(v string) *RevokeTablePermissionResponseBody
	GetRequestId() *string
}

type RevokeTablePermissionResponseBody struct {
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

func (s RevokeTablePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeTablePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeTablePermissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *RevokeTablePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeTablePermissionResponseBody) SetData(v bool) *RevokeTablePermissionResponseBody {
	s.Data = &v
	return s
}

func (s *RevokeTablePermissionResponseBody) SetRequestId(v string) *RevokeTablePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeTablePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
