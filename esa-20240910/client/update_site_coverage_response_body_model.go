// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteCoverageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSiteCoverageResponseBody
	GetRequestId() *string
}

type UpdateSiteCoverageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 65C66B7B-671A-8297-9187-2R5477247B76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteCoverageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteCoverageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteCoverageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSiteCoverageResponseBody) SetRequestId(v string) *UpdateSiteCoverageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteCoverageResponseBody) Validate() error {
	return dara.Validate(s)
}
