// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStreamingDataSourceResponseBody
	GetRequestId() *string
}

type ModifyStreamingDataSourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStreamingDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStreamingDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStreamingDataSourceResponseBody) SetRequestId(v string) *ModifyStreamingDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStreamingDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
