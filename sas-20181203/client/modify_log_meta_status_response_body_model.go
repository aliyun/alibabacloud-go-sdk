// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogMetaStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLogMetaStatusResponseBody
	GetRequestId() *string
}

type ModifyLogMetaStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 861445A7-B6D6-5825-B015-CD46ED90613A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogMetaStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogMetaStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogMetaStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLogMetaStatusResponseBody) SetRequestId(v string) *ModifyLogMetaStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLogMetaStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
