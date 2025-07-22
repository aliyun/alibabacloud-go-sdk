// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRecordStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppRecordStatusResponseBody
	GetRequestId() *string
}

type ModifyAppRecordStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D53303DB-AA68-5D09-90C2-A345072DCC5A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppRecordStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppRecordStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppRecordStatusResponseBody) SetRequestId(v string) *ModifyAppRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppRecordStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
