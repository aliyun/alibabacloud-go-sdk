// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomBlockRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCustomBlockRecordResponseBody
	GetRequestId() *string
}

type ModifyCustomBlockRecordResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EE4B1F-CEF8-5A75-86D3-D012CB3D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCustomBlockRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomBlockRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomBlockRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomBlockRecordResponseBody) SetRequestId(v string) *ModifyCustomBlockRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomBlockRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
