// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyAppInfoResponseBody
	GetData() *bool
	SetRequestId(v string) *ModifyAppInfoResponseBody
	GetRequestId() *string
}

type ModifyAppInfoResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppInfoResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppInfoResponseBody) SetData(v bool) *ModifyAppInfoResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyAppInfoResponseBody) SetRequestId(v string) *ModifyAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
