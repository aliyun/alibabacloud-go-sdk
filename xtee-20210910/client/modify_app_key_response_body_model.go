// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppKeyResponseBody
	GetRequestId() *string
	SetData(v bool) *ModifyAppKeyResponseBody
	GetData() *bool
}

type ModifyAppKeyResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned data object.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ModifyAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppKeyResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyAppKeyResponseBody) SetRequestId(v string) *ModifyAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppKeyResponseBody) SetData(v bool) *ModifyAppKeyResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
