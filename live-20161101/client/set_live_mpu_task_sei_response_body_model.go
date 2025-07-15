// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveMpuTaskSeiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveMpuTaskSeiResponseBody
	GetRequestId() *string
}

type SetLiveMpuTaskSeiResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BF95F2A-3B24-4CDE-9346-7F6FA86697A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveMpuTaskSeiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveMpuTaskSeiResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveMpuTaskSeiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveMpuTaskSeiResponseBody) SetRequestId(v string) *SetLiveMpuTaskSeiResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveMpuTaskSeiResponseBody) Validate() error {
	return dara.Validate(s)
}
