// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteNameExclusiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSiteNameExclusiveResponseBody
	GetRequestId() *string
}

type UpdateSiteNameExclusiveResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteNameExclusiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteNameExclusiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteNameExclusiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSiteNameExclusiveResponseBody) SetRequestId(v string) *UpdateSiteNameExclusiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteNameExclusiveResponseBody) Validate() error {
	return dara.Validate(s)
}
