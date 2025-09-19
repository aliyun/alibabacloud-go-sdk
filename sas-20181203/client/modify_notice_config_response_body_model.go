// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNoticeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNoticeConfigResponseBody
	GetRequestId() *string
}

type ModifyNoticeConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5989D067-621F-51E2-A636-D94D1388****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNoticeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNoticeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNoticeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNoticeConfigResponseBody) SetRequestId(v string) *ModifyNoticeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNoticeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
