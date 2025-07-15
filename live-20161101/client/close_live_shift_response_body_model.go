// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseLiveShiftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseLiveShiftResponseBody
	GetRequestId() *string
}

type CloseLiveShiftResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseLiveShiftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseLiveShiftResponseBody) GoString() string {
	return s.String()
}

func (s *CloseLiveShiftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseLiveShiftResponseBody) SetRequestId(v string) *CloseLiveShiftResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseLiveShiftResponseBody) Validate() error {
	return dara.Validate(s)
}
