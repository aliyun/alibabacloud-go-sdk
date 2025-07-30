// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTensorboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTensorboardResponseBody
	GetRequestId() *string
	SetTensorboardId(v string) *DeleteTensorboardResponseBody
	GetTensorboardId() *string
}

type DeleteTensorboardResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The TensorBoard ID.
	//
	// example:
	//
	// tensorboard-20210114104214-vf9lowjt3pso
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s DeleteTensorboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTensorboardResponseBody) GetTensorboardId() *string {
	return s.TensorboardId
}

func (s *DeleteTensorboardResponseBody) SetRequestId(v string) *DeleteTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTensorboardResponseBody) SetTensorboardId(v string) *DeleteTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *DeleteTensorboardResponseBody) Validate() error {
	return dara.Validate(s)
}
