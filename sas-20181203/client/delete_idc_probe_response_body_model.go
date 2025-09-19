// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdcProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIdcProbeResponseBody
	GetRequestId() *string
}

type DeleteIdcProbeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5DFD6277-CC36-57F7-ACE6-F5952XXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIdcProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdcProbeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdcProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIdcProbeResponseBody) SetRequestId(v string) *DeleteIdcProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdcProbeResponseBody) Validate() error {
	return dara.Validate(s)
}
