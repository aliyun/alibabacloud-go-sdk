// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SynchronizeAppResponseBody
	GetRequestId() *string
}

type SynchronizeAppResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SynchronizeAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeAppResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SynchronizeAppResponseBody) SetRequestId(v string) *SynchronizeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *SynchronizeAppResponseBody) Validate() error {
	return dara.Validate(s)
}
