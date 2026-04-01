// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLCollectorRetentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySQLCollectorRetentionResponseBody
	GetRequestId() *string
}

type ModifySQLCollectorRetentionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 21383BB3-3845-4628-B422-B4FB5C83DEBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySQLCollectorRetentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLCollectorRetentionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorRetentionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySQLCollectorRetentionResponseBody) SetRequestId(v string) *ModifySQLCollectorRetentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySQLCollectorRetentionResponseBody) Validate() error {
	return dara.Validate(s)
}
