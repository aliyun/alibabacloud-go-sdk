// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleBatchMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSampleBatchMetaResponseBody
	GetRequestId() *string
}

type DeleteSampleBatchMetaResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSampleBatchMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleBatchMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSampleBatchMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSampleBatchMetaResponseBody) SetRequestId(v string) *DeleteSampleBatchMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSampleBatchMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
