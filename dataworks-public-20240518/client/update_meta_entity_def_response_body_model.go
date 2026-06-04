// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaEntityDefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMetaEntityDefResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMetaEntityDefResponseBody
	GetSuccess() *bool
}

type UpdateMetaEntityDefResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E08F38AB-3BA0-5047-8E9E-9AA4839263EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMetaEntityDefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityDefResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityDefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaEntityDefResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMetaEntityDefResponseBody) SetRequestId(v string) *UpdateMetaEntityDefResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaEntityDefResponseBody) SetSuccess(v bool) *UpdateMetaEntityDefResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMetaEntityDefResponseBody) Validate() error {
	return dara.Validate(s)
}
