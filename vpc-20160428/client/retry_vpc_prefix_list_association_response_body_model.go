// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryVpcPrefixListAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RetryVpcPrefixListAssociationResponseBody
	GetRequestId() *string
}

type RetryVpcPrefixListAssociationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetryVpcPrefixListAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryVpcPrefixListAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *RetryVpcPrefixListAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryVpcPrefixListAssociationResponseBody) SetRequestId(v string) *RetryVpcPrefixListAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryVpcPrefixListAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
