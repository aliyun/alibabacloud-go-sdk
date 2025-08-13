// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCreatedByResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenCreatedByResponseBody
	GetRequestId() *string
}

type OpenCreatedByResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 74F03511-FDFA-54AF-96A4-71575B41E74D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenCreatedByResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenCreatedByResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCreatedByResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenCreatedByResponseBody) SetRequestId(v string) *OpenCreatedByResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCreatedByResponseBody) Validate() error {
	return dara.Validate(s)
}
