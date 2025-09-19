// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageFixCycleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyImageFixCycleConfigResponseBody
	GetData() *bool
	SetRequestId(v string) *ModifyImageFixCycleConfigResponseBody
	GetRequestId() *string
}

type ModifyImageFixCycleConfigResponseBody struct {
	// Indicates whether the configurations of the scheduled image fix are modified.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7A4DF7E-57A4-5BBF-8290-223754AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageFixCycleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageFixCycleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageFixCycleConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyImageFixCycleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImageFixCycleConfigResponseBody) SetData(v bool) *ModifyImageFixCycleConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyImageFixCycleConfigResponseBody) SetRequestId(v string) *ModifyImageFixCycleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImageFixCycleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
