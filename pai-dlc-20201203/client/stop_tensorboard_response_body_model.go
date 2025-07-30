// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTensorboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopTensorboardResponseBody
	GetRequestId() *string
	SetTensorboardId(v string) *StopTensorboardResponseBody
	GetTensorboardId() *string
}

type StopTensorboardResponseBody struct {
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

func (s StopTensorboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *StopTensorboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTensorboardResponseBody) GetTensorboardId() *string {
	return s.TensorboardId
}

func (s *StopTensorboardResponseBody) SetRequestId(v string) *StopTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTensorboardResponseBody) SetTensorboardId(v string) *StopTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *StopTensorboardResponseBody) Validate() error {
	return dara.Validate(s)
}
