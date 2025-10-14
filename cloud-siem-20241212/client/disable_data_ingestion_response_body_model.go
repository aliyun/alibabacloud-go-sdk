// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDataIngestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDataIngestionResponseBody
	GetRequestId() *string
}

type DisableDataIngestionResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDataIngestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDataIngestionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDataIngestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDataIngestionResponseBody) SetRequestId(v string) *DisableDataIngestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDataIngestionResponseBody) Validate() error {
	return dara.Validate(s)
}
