// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyServiceInfoResponseBody
	GetData() *bool
	SetRequestId(v string) *ModifyServiceInfoResponseBody
	GetRequestId() *string
}

type ModifyServiceInfoResponseBody struct {
	// Returned data.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyServiceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyServiceInfoResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyServiceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyServiceInfoResponseBody) SetData(v bool) *ModifyServiceInfoResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyServiceInfoResponseBody) SetRequestId(v string) *ModifyServiceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyServiceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
