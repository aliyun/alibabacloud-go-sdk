// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubCrowdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSubCrowdResponseBody
	GetRequestId() *string
}

type DeleteSubCrowdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EE97D06A-2AA0-5AD9-B6CF-8A267924D691
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubCrowdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubCrowdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubCrowdResponseBody) SetRequestId(v string) *DeleteSubCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubCrowdResponseBody) Validate() error {
	return dara.Validate(s)
}
