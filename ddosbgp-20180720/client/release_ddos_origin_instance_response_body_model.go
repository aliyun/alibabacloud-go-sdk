// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDdosOriginInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseDdosOriginInstanceResponseBody
	GetRequestId() *string
}

type ReleaseDdosOriginInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4B379C2-9319-4C6B-B579-FE36831B09F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseDdosOriginInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDdosOriginInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseDdosOriginInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseDdosOriginInstanceResponseBody) SetRequestId(v string) *ReleaseDdosOriginInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseDdosOriginInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
