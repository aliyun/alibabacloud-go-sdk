// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridCloudServerResponseBody
	GetRequestId() *string
}

type ModifyHybridCloudServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 26DCD663-5EB8-5103-B270-E24A32***5F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridCloudServerResponseBody) SetRequestId(v string) *ModifyHybridCloudServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridCloudServerResponseBody) Validate() error {
	return dara.Validate(s)
}
