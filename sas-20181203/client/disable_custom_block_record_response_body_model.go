// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomBlockRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCustomBlockRecordResponseBody
	GetRequestId() *string
}

type DisableCustomBlockRecordResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F749D54C-3CA0-5F68-835C-AD35A2BD29EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCustomBlockRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomBlockRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCustomBlockRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCustomBlockRecordResponseBody) SetRequestId(v string) *DisableCustomBlockRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCustomBlockRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
