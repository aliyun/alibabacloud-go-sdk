// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgAccountLoginPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAgAccountLoginPermissionResponseBody
	GetCode() *string
	SetHasLoginPermission(v bool) *QueryAgAccountLoginPermissionResponseBody
	GetHasLoginPermission() *bool
	SetMessage(v string) *QueryAgAccountLoginPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAgAccountLoginPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAgAccountLoginPermissionResponseBody
	GetSuccess() *bool
}

type QueryAgAccountLoginPermissionResponseBody struct {
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HasLoginPermission *bool   `json:"HasLoginPermission,omitempty" xml:"HasLoginPermission,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAgAccountLoginPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAgAccountLoginPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAgAccountLoginPermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAgAccountLoginPermissionResponseBody) GetHasLoginPermission() *bool {
	return s.HasLoginPermission
}

func (s *QueryAgAccountLoginPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAgAccountLoginPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAgAccountLoginPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAgAccountLoginPermissionResponseBody) SetCode(v string) *QueryAgAccountLoginPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAgAccountLoginPermissionResponseBody) SetHasLoginPermission(v bool) *QueryAgAccountLoginPermissionResponseBody {
	s.HasLoginPermission = &v
	return s
}

func (s *QueryAgAccountLoginPermissionResponseBody) SetMessage(v string) *QueryAgAccountLoginPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAgAccountLoginPermissionResponseBody) SetRequestId(v string) *QueryAgAccountLoginPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAgAccountLoginPermissionResponseBody) SetSuccess(v bool) *QueryAgAccountLoginPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAgAccountLoginPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
