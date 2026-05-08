// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAvatarProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateAvatarProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateAvatarProjectResponseBody
	GetSuccess() *bool
}

type OperateAvatarProjectResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 86A90C40-D1AB-50DA-A4B1-0D545F80F2FE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OperateAvatarProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateAvatarProjectResponseBody) GoString() string {
	return s.String()
}

func (s *OperateAvatarProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateAvatarProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateAvatarProjectResponseBody) SetRequestId(v string) *OperateAvatarProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateAvatarProjectResponseBody) SetSuccess(v bool) *OperateAvatarProjectResponseBody {
	s.Success = &v
	return s
}

func (s *OperateAvatarProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
