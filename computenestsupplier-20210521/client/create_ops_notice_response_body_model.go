// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpsNoticeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOpsNoticeResponseBody
	GetRequestId() *string
}

type CreateOpsNoticeResponseBody struct {
	// example:
	//
	// E73F09DC-6C13-5CB1-A10F-7A4E125ABD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOpsNoticeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpsNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpsNoticeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpsNoticeResponseBody) SetRequestId(v string) *CreateOpsNoticeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpsNoticeResponseBody) Validate() error {
	return dara.Validate(s)
}
