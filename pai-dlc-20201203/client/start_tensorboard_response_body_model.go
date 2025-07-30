// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTensorboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartTensorboardResponseBody
	GetRequestId() *string
	SetTensorboardId(v string) *StartTensorboardResponseBody
	GetTensorboardId() *string
}

type StartTensorboardResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The TensorBoard instance ID.
	//
	// example:
	//
	// tensorboard-20210114104214-vf9lowjt3pso
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s StartTensorboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *StartTensorboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTensorboardResponseBody) GetTensorboardId() *string {
	return s.TensorboardId
}

func (s *StartTensorboardResponseBody) SetRequestId(v string) *StartTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTensorboardResponseBody) SetTensorboardId(v string) *StartTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *StartTensorboardResponseBody) Validate() error {
	return dara.Validate(s)
}
