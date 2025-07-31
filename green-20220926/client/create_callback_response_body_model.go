// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateCallbackResponseBody
	GetData() *int64
	SetRequestId(v string) *CreateCallbackResponseBody
	GetRequestId() *string
}

type CreateCallbackResponseBody struct {
	// example:
	//
	// True
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCallbackResponseBody) SetData(v int64) *CreateCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCallbackResponseBody) SetRequestId(v string) *CreateCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
