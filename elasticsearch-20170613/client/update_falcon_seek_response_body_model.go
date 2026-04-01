// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFalconSeekResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFalconSeekResponseBody
	GetRequestId() *string
}

type UpdateFalconSeekResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFalconSeekResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFalconSeekResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFalconSeekResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFalconSeekResponseBody) SetRequestId(v string) *UpdateFalconSeekResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFalconSeekResponseBody) Validate() error {
	return dara.Validate(s)
}
