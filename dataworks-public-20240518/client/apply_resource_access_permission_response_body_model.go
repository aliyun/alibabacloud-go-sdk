// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyResourceAccessPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *ApplyResourceAccessPermissionResponseBody
	GetData() []*string
	SetRequestId(v string) *ApplyResourceAccessPermissionResponseBody
	GetRequestId() *string
}

type ApplyResourceAccessPermissionResponseBody struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyResourceAccessPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyResourceAccessPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyResourceAccessPermissionResponseBody) GetData() []*string {
	return s.Data
}

func (s *ApplyResourceAccessPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyResourceAccessPermissionResponseBody) SetData(v []*string) *ApplyResourceAccessPermissionResponseBody {
	s.Data = v
	return s
}

func (s *ApplyResourceAccessPermissionResponseBody) SetRequestId(v string) *ApplyResourceAccessPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyResourceAccessPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
