// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTensorboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTensorboardResponseBody
	GetRequestId() *string
	SetTensorboardId(v string) *UpdateTensorboardResponseBody
	GetTensorboardId() *string
}

type UpdateTensorboardResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the TensorBoard instance.
	//
	// example:
	//
	// tensorboard-20210114104214-xxxxxxxx
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s UpdateTensorboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTensorboardResponseBody) GetTensorboardId() *string {
	return s.TensorboardId
}

func (s *UpdateTensorboardResponseBody) SetRequestId(v string) *UpdateTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTensorboardResponseBody) SetTensorboardId(v string) *UpdateTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *UpdateTensorboardResponseBody) Validate() error {
	return dara.Validate(s)
}
