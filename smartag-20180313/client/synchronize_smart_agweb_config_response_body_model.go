// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeSmartAGWebConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SynchronizeSmartAGWebConfigResponseBody
	GetRequestId() *string
}

type SynchronizeSmartAGWebConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CBBE5EBF-69C1-4395-B36B-26B7605F87EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SynchronizeSmartAGWebConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeSmartAGWebConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeSmartAGWebConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SynchronizeSmartAGWebConfigResponseBody) SetRequestId(v string) *SynchronizeSmartAGWebConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
