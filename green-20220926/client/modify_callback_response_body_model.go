// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyCallbackResponseBody
	GetData() *bool
	SetRequestId(v string) *ModifyCallbackResponseBody
	GetRequestId() *string
}

type ModifyCallbackResponseBody struct {
	// Return result.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Backend-assigned ID, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCallbackResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCallbackResponseBody) SetData(v bool) *ModifyCallbackResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyCallbackResponseBody) SetRequestId(v string) *ModifyCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
