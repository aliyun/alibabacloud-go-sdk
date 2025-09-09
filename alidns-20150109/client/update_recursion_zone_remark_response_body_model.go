// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecursionZoneRemarkResponseBody
	GetRequestId() *string
}

type UpdateRecursionZoneRemarkResponseBody struct {
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecursionZoneRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecursionZoneRemarkResponseBody) SetRequestId(v string) *UpdateRecursionZoneRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecursionZoneRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
