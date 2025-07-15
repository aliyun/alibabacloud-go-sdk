// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealTimeRecordCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RealTimeRecordCommandResponseBody
	GetRequestId() *string
}

type RealTimeRecordCommandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RealTimeRecordCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RealTimeRecordCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RealTimeRecordCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RealTimeRecordCommandResponseBody) SetRequestId(v string) *RealTimeRecordCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RealTimeRecordCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
