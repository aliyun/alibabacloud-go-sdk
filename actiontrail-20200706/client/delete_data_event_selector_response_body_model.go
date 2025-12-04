// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataEventSelectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataEventSelectorResponseBody
	GetRequestId() *string
}

type DeleteDataEventSelectorResponseBody struct {
	// example:
	//
	// 1D9DD159-DFFF-4882-ACEC-B4A727E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataEventSelectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataEventSelectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataEventSelectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataEventSelectorResponseBody) SetRequestId(v string) *DeleteDataEventSelectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataEventSelectorResponseBody) Validate() error {
	return dara.Validate(s)
}
