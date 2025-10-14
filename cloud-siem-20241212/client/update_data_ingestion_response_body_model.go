// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataIngestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataIngestionResponseBody
	GetRequestId() *string
}

type UpdateDataIngestionResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataIngestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataIngestionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataIngestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataIngestionResponseBody) SetRequestId(v string) *UpdateDataIngestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataIngestionResponseBody) Validate() error {
	return dara.Validate(s)
}
