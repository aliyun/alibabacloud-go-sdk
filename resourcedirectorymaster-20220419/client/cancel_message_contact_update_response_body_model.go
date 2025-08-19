// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMessageContactUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelMessageContactUpdateResponseBody
	GetRequestId() *string
}

type CancelMessageContactUpdateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelMessageContactUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelMessageContactUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *CancelMessageContactUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelMessageContactUpdateResponseBody) SetRequestId(v string) *CancelMessageContactUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelMessageContactUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
