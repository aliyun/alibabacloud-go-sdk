// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseColdDataVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseColdDataVolumeResponseBody
	GetRequestId() *string
}

type ReleaseColdDataVolumeResponseBody struct {
	// example:
	//
	// EA330983-C895-57C0-AE82-5A63106EBB10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseColdDataVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseColdDataVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseColdDataVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseColdDataVolumeResponseBody) SetRequestId(v string) *ReleaseColdDataVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseColdDataVolumeResponseBody) Validate() error {
	return dara.Validate(s)
}
