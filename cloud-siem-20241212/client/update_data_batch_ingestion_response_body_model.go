// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataBatchIngestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataBatchIngestionResponseBody
	GetRequestId() *string
}

type UpdateDataBatchIngestionResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataBatchIngestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataBatchIngestionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataBatchIngestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataBatchIngestionResponseBody) SetRequestId(v string) *UpdateDataBatchIngestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataBatchIngestionResponseBody) Validate() error {
	return dara.Validate(s)
}
